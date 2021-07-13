/*
Package csi is CSI driver interface for OSD
Copyright 2017 Portworx

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package csi

import (
	"fmt"
	"os"
	"strings"

	csi "github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/libopenstorage/openstorage/api"
	"github.com/libopenstorage/openstorage/pkg/grpcutil"
	"github.com/portworx/kvdb"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ephemeralDenyList = []string{
		api.SpecPriorityAlias,
		api.SpecPriority,
		api.SpecSticky,
		api.SpecScale,
	}
)

func (s *OsdCsiServer) NodeGetInfo(
	ctx context.Context,
	req *csi.NodeGetInfoRequest,
) (*csi.NodeGetInfoResponse, error) {

	clus, err := s.cluster.Enumerate()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to Enumerate cluster: %s", err)
	}

	result := &csi.NodeGetInfoResponse{
		NodeId: clus.NodeId,
	}

	return result, nil
}

// NodePublishVolume is a CSI API call which mounts the volume on the specified
// target path on the node.
//
// TODO: Support READ ONLY Mounts
//
func (s *OsdCsiServer) NodePublishVolume(
	ctx context.Context,
	req *csi.NodePublishVolumeRequest,
) (*csi.NodePublishVolumeResponse, error) {
	volumeId := req.GetVolumeId()
	targetPath := req.GetTargetPath()

	logrus.Infof("csi.NodePublishVolume request received. VolumeID: %s, TargetPath: %s", volumeId, targetPath)

	// Check arguments
	if len(volumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume id must be provided")
	}
	if len(targetPath) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Target path must be provided")
	}
	if req.GetVolumeCapability() == nil || req.GetVolumeCapability().GetAccessMode() == nil ||
		req.GetVolumeCapability().GetAccessMode().Mode == csi.VolumeCapability_AccessMode_UNKNOWN {
		return nil, status.Error(codes.InvalidArgument, "Volume access mode must be provided")
	}

	// Ensure target location is created correctly
	isBlockAccessType := false
	if req.GetVolumeCapability().GetBlock() != nil {
		isBlockAccessType = true
	}
	if err := ensureMountPathCreated(targetPath, isBlockAccessType); err != nil {
		return nil, status.Errorf(
			codes.Aborted,
			"Failed to use target location %s: %s",
			targetPath,
			err.Error())
	}

	// Get grpc connection
	conn, err := s.getConn()
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Unable to connect to SDK server: %v", err)
	}

	// Get secret if any was passed
	ctx = s.setupContextWithToken(ctx, req.GetSecrets())
	ctx, cancel := grpcutil.WithDefaultTimeout(ctx)
	defer cancel()

	// Check if block device
	driverType := s.driver.Type()
	if driverType != api.DriverType_DRIVER_TYPE_BLOCK &&
		req.GetVolumeCapability().GetBlock() != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Trying to attach as block a non block device")
	}

	// Gather volume attributes
	spec, locator, _, err := s.specHandler.SpecFromOpts(req.GetVolumeContext())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Invalid volume attributes: %#v",
			req.GetVolumeContext())
	}

	// Get volume encryption info from req.Secrets
	driverOpts := s.addEncryptionInfoToLabels(make(map[string]string), req.GetSecrets())

	// Parse storage class 'mountOptions' flags from CSI req
	// flags from 'mountOptions' will be used as the only source of truth for Pure volumes upon mounting
	if req.GetVolumeCapability() != nil && req.GetVolumeCapability().GetMount() != nil {
		mountFlags := strings.Join(req.GetVolumeCapability().GetMount().GetMountFlags(), ",")
		if mountFlags != "" {
			driverOpts[api.SpecCSIMountOptions] = mountFlags
		}
	}

	// can use either spec.Ephemeral or VolumeContext label
	if req.GetVolumeContext()["csi.storage.k8s.io/ephemeral"] == "true" || spec.Ephemeral {
		if !s.allowInlineVolumes {
			return nil, status.Error(codes.InvalidArgument, "CSI ephemeral inline volumes are disabled on this cluster")
		}

		if err := validateEphemeralVolumeAttributes(req.GetVolumeContext()); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		spec.Ephemeral = true
		volumes := api.NewOpenStorageVolumeClient(conn)
		resp, err := volumes.Create(ctx, &api.SdkVolumeCreateRequest{
			Name:   volumeId,
			Spec:   spec,
			Labels: locator.GetVolumeLabels(),
		})
		if err != nil {
			return nil, err
		}
		volumeId = resp.VolumeId
	}

	// prepare for mount/attaching
	opts := &api.SdkVolumeAttachOptions{
		SecretName: spec.GetPassphrase(),
	}
	mounts := api.NewOpenStorageMountAttachClient(conn)
	if driverType == api.DriverType_DRIVER_TYPE_BLOCK {
		// attach is assumed to be idempotent
		// attach is assumed to return the same DevicePath on each call
		if _, err = mounts.Attach(ctx, &api.SdkVolumeAttachRequest{
			VolumeId:      volumeId,
			Options:       opts,
			DriverOptions: driverOpts,
		}); err != nil {
			if spec.Ephemeral {
				logrus.Errorf("Failed to attach ephemeral volume %s: %v", volumeId, err.Error())
				s.cleanupEphemeral(ctx, conn, volumeId, false)
			}
			return nil, err
		}
	}

	// for volumes with mount access type just mount volume onto the path
	if _, err := mounts.Mount(ctx, &api.SdkVolumeMountRequest{
		VolumeId:      volumeId,
		MountPath:     targetPath,
		Options:       opts,
		DriverOptions: driverOpts,
	}); err != nil {
		if spec.Ephemeral {
			logrus.Errorf("Failed to mount ephemeral volume %s: %v", volumeId, err.Error())
			s.cleanupEphemeral(ctx, conn, volumeId, true)
		}
		return nil, err
	}

	logrus.Infof("CSI Volume %s mounted on %s",
		volumeId,
		req.GetTargetPath())

	return &csi.NodePublishVolumeResponse{}, nil
}

// NodeUnpublishVolume is a CSI API call which unmounts the volume.
func (s *OsdCsiServer) NodeUnpublishVolume(
	ctx context.Context,
	req *csi.NodeUnpublishVolumeRequest,
) (*csi.NodeUnpublishVolumeResponse, error) {
	volumeId := req.GetVolumeId()
	targetPath := req.GetTargetPath()

	logrus.Infof("csi.NodeUnpublishVolume request received. VolumeID: %s, TargetPath: %s", volumeId, targetPath)

	// Check arguments
	if len(volumeId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume id must be provided")
	}
	if len(targetPath) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Target path must be provided")
	}

	// Get volume information
	vols, err := s.driver.Inspect([]string{req.GetVolumeId()})
	if err != nil || len(vols) < 1 {
		if err == kvdb.ErrNotFound {
			logrus.Infof("Volume %s was deleted or cannot be found: %s", req.GetVolumeId(), err.Error())
			return &csi.NodeUnpublishVolumeResponse{}, nil
		} else if err != nil {
			return nil, status.Errorf(codes.NotFound, "Volume id %s not found: %s",
				req.GetVolumeId(),
				err.Error())
		} else {
			logrus.Infof("Volume %s was deleted or cannot be found", req.GetVolumeId())
			return &csi.NodeUnpublishVolumeResponse{}, nil
		}
	}

	if err = s.driver.Unmount(req.GetVolumeId(), req.GetTargetPath(), nil); err != nil {
		logrus.Infof("unable to unmount volume %s onto %s: %s",
			req.GetVolumeId(),
			req.GetTargetPath(),
			err.Error(),
		)
	}

	if s.driver.Type() == api.DriverType_DRIVER_TYPE_BLOCK {
		if err = s.driver.Detach(volumeId, nil); err != nil {
			return nil, status.Errorf(
				codes.Internal,
				"Unable to detach volume: %s",
				err.Error())
		}
	}

	// Attempt to remove volume path
	// Kubernetes handles this after NodeUnpublishVolume finishes, but this allows for cross-CO compatibility
	if err := os.Remove(req.GetTargetPath()); err != nil && !os.IsNotExist(err) {
		logrus.Warnf("Failed to delete mount path %s: %s", targetPath, err.Error())
	}

	// Return error to Kubelet if mount path still exists to force a retry
	if _, err := os.Stat(targetPath); !os.IsNotExist(err) {
		return nil, status.Errorf(
			codes.Internal,
			"Mount path still exists: %s",
			targetPath)
	}

	logrus.Infof("CSI Volume %s unmounted from path %s", volumeId, targetPath)

	return &csi.NodeUnpublishVolumeResponse{}, nil
}

// NodeGetCapabilities is a CSI API function which seems to be setup for
// future patches
func (s *OsdCsiServer) NodeGetCapabilities(
	ctx context.Context,
	req *csi.NodeGetCapabilitiesRequest,
) (*csi.NodeGetCapabilitiesResponse, error) {

	logrus.Infof("csi.NodeGetCapabilities request received")

	caps := []csi.NodeServiceCapability_RPC_Type{
		// Getting volume stats for volume health monitoring
		csi.NodeServiceCapability_RPC_GET_VOLUME_STATS,
		// Indicates that the Node service can report volume conditions.
		csi.NodeServiceCapability_RPC_VOLUME_CONDITION,
	}

	var serviceCapabilities []*csi.NodeServiceCapability
	for _, cap := range caps {
		serviceCapabilities = append(serviceCapabilities, &csi.NodeServiceCapability{
			Type: &csi.NodeServiceCapability_Rpc{
				Rpc: &csi.NodeServiceCapability_RPC{
					Type: cap,
				},
			},
		})
	}

	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: serviceCapabilities,
	}, nil
}

func getVolumeCondition(vol *api.Volume) *csi.VolumeCondition {
	condition := &csi.VolumeCondition{}
	if vol.Status != api.VolumeStatus_VOLUME_STATUS_UP {
		condition.Abnormal = true
	}

	switch vol.Status {
	case api.VolumeStatus_VOLUME_STATUS_UP:
		condition.Message = "Volume status is up"

	case api.VolumeStatus_VOLUME_STATUS_NOT_PRESENT:
		condition.Message = "Volume status is not present"

	case api.VolumeStatus_VOLUME_STATUS_DOWN:
		condition.Message = "Volume status is down"

	case api.VolumeStatus_VOLUME_STATUS_DEGRADED:
		condition.Message = "Volume status is degraded"

	default:
		condition.Message = "Volume status is unknown"
	}

	return condition
}

// NodeGetVolumeStats get volume stats for a given node.
// This function skips auth and directly hits the driver as it is read-only
// and only exposed via the CSI unix domain socket. If a secrets field is added
// in csi.NodeGetVolumeStatsRequest, we can update this to hit the SDK and use auth.
func (s *OsdCsiServer) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	logrus.Infof("NodeGetVolumeStats request received. VolumeID: %s, VolumePath: %s", req.GetVolumeId(), req.GetVolumePath())

	// Check arguments
	id := req.GetVolumeId()
	if len(id) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume id must be provided")
	}
	path := req.GetVolumePath()
	if len(path) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume path must be provided")
	}

	// Driver inspect as NodeGetVolumeStatsRequest does not support secrets
	vol, err := s.driverGetVolume(req.GetVolumeId())
	if err != nil {
		return nil, err
	}

	var attachPathMatch bool
	for _, attachPath := range vol.AttachPath {
		if attachPath == path {
			attachPathMatch = true
		}
	}
	if !attachPathMatch {
		return nil, status.Errorf(codes.NotFound, "Volume %s not mounted on path %s", id, path)
	}

	// Define volume usage
	total := int64(vol.Spec.Size)
	used := int64(vol.Usage)
	usage := &csi.VolumeUsage{
		Available: total - used,
		Total:     total,
		Used:      used,
		Unit:      csi.VolumeUsage_BYTES,
	}

	// Define volume condition
	return &csi.NodeGetVolumeStatsResponse{
		Usage: []*csi.VolumeUsage{
			usage,
		},
		VolumeCondition: getVolumeCondition(vol),
	}, nil
}

// cleanupEphemeral detaches and deletes an ephemeral volume if either attach or mount fails
func (s *OsdCsiServer) cleanupEphemeral(ctx context.Context, conn *grpc.ClientConn, volumeId string, detach bool) {
	if detach {
		mounts := api.NewOpenStorageMountAttachClient(conn)
		if _, err := mounts.Detach(ctx, &api.SdkVolumeDetachRequest{
			VolumeId: volumeId,
		}); err != nil {
			logrus.Errorf("Failed to detach ephemeral volume %s during cleanup: %v", volumeId, err.Error())
			return
		}
	}
	volumes := api.NewOpenStorageVolumeClient(conn)
	if _, err := volumes.Delete(ctx, &api.SdkVolumeDeleteRequest{
		VolumeId: volumeId,
	}); err != nil {
		logrus.Errorf("Failed to delete ephemeral volume %s during cleanup: %v", volumeId, err.Error())
	}
}

func ensureMountPathCreated(targetPath string, isBlock bool) error {
	// Check if targetpath exists
	fileInfo, err := os.Lstat(targetPath)
	if err != nil && os.IsNotExist(err) {
		// Create if does not exist
		// 1. Block - create targetPath file
		// 2. Mount - create targetpath directory
		if isBlock {
			if err = makeFile(targetPath); err != nil {
				return err
			}
		} else {
			if err = makeDir(targetPath); err != nil {
				return err
			}
		}

		return nil
	} else if err != nil {
		return fmt.Errorf(
			"unknown error while verifying target location %s: %s",
			targetPath,
			err.Error())
	}

	// Check for directory or file.
	// 1. Block - should be file
	// 2. Mount - should be directory
	if isBlock {
		if fileInfo.IsDir() {
			return fmt.Errorf("Target location %s is not a file", targetPath)
		}
	} else {
		if !fileInfo.IsDir() {
			return fmt.Errorf("Target location %s is not a directory", targetPath)
		}
	}

	return nil
}

func validateEphemeralVolumeAttributes(volumeAttributes map[string]string) error {
	for attr := range volumeAttributes {
		for _, deny := range ephemeralDenyList {
			if attr == deny {
				return fmt.Errorf("invalid ephemeral volume attribute provided. "+
					"Volume attributes %v are not allowed for ephemeral volumes", ephemeralDenyList)
			}
		}
	}

	return nil
}

func makeFile(pathname string) error {
	f, err := os.OpenFile(pathname, os.O_CREATE, os.FileMode(0644))
	defer func() {
		err := f.Close()
		if err != nil {
			logrus.Warnf("failed to close file: %s", err.Error())
		}
	}()
	if err != nil {
		if !os.IsExist(err) {
			return fmt.Errorf("failed to create block file: %s", err.Error())
		}
	}

	return nil
}

func makeDir(targetPath string) error {
	err := os.MkdirAll(targetPath, 0750)
	if err != nil {
		return fmt.Errorf(
			"failed to create target path %s: %s",
			targetPath,
			err.Error())
	}

	return nil
}
