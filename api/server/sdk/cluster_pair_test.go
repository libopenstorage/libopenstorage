/*
Package sdk is the gRPC implementation of the SDK gRPC server
Copyright 2018 Portworx

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
package sdk

import (
	"context"
	"testing"

	"github.com/libopenstorage/openstorage/api"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestClusterPairServer_CreateSuccess(t *testing.T) {
	// Create server and client connection
	s := newTestServer(t)
	defer s.Stop()

	// Create response
	remoteClusterID := "id"
	remoteClusterName := "name"
	remoteClusterIP := "127.0.0.1"
	var remoteClusterPort uint32
	remoteClusterPort = uint32(12345)

	req := api.ClusterPairCreateRequest{
		RemoteClusterIp:    remoteClusterIP,
		RemoteClusterPort:  remoteClusterPort,
		RemoteClusterToken: "<Auth-Token>",
		SetDefault:         false,
	}
	resp := &api.ClusterPairCreateResponse{
		RemoteClusterId:   remoteClusterID,
		RemoteClusterName: remoteClusterName,
	}

	s.MockCluster().
		EXPECT().
		CreatePair(&api.ClusterPairCreateRequest{
			RemoteClusterIp:    remoteClusterIP,
			RemoteClusterPort:  remoteClusterPort,
			RemoteClusterToken: "<Auth-Token>",
			SetDefault:         false}).
		Return(resp, nil)

	// Setup client
	c := api.NewOpenStorageClusterPairClient(s.Conn())
	// create the pair
	r, err := c.Create(context.Background(), &req)
	assert.NoError(t, err)
	assert.NotNil(t, r.GetRemoteClusterId())
	assert.Equal(t, remoteClusterID, r.GetRemoteClusterId())
	assert.Equal(t, remoteClusterName, r.GetRemoteClusterName())
}
func TestClusterPairServer_CreateFailure(t *testing.T) {
	// Create server and client connection
	s := newTestServer(t)
	defer s.Stop()

	// Create response
	remoteClusterIP := "127.0.0.1"
	var remoteClusterPort uint32
	remoteClusterPort = uint32(12345)

	noip := &api.ClusterPairCreateRequest{
		RemoteClusterPort:  remoteClusterPort,
		RemoteClusterToken: "<Auth-Token>",
		SetDefault:         false,
	}

	noport := &api.ClusterPairCreateRequest{
		RemoteClusterIp:    remoteClusterIP,
		RemoteClusterToken: "<Auth-Token>",
		SetDefault:         false,
	}

	notoken := &api.ClusterPairCreateRequest{
		RemoteClusterIp:   remoteClusterIP,
		RemoteClusterPort: remoteClusterPort,
		SetDefault:        false,
	}

	// Setup client
	c := api.NewOpenStorageClusterPairClient(s.Conn())
	// create the pair
	//noip
	r, err := c.Create(context.Background(), noip)
	assert.Error(t, err)
	assert.Nil(t, r)
	serverError, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, serverError.Code(), codes.InvalidArgument)
	assert.Contains(t, serverError.Message(), "Must supply Remote cluster IP")

	r, err = c.Create(context.Background(), noport)
	assert.Error(t, err)
	assert.Nil(t, r)
	serverError, ok = status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, serverError.Code(), codes.InvalidArgument)
	assert.Contains(t, serverError.Message(), "Must supply Remote cluster Port")

	r, err = c.Create(context.Background(), notoken)
	assert.Error(t, err)
	assert.Nil(t, r)
	serverError, ok = status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, serverError.Code(), codes.InvalidArgument)
	assert.Contains(t, serverError.Message(), "Must supply Authentication Token")

}
