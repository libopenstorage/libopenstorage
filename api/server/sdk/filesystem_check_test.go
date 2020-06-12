package sdk

import (
	"context"
	"testing"

	"github.com/libopenstorage/openstorage/api"
	"github.com/stretchr/testify/assert"
)

func TestSdkFilesystemCheckCheckHealth(t *testing.T) {
	// Create server and client connection
	s := newTestServer(t)
	defer s.Stop()

	testVolumeId := "myvol"
	testStatus := api.FilesystemCheck_FS_CHECK_STARTED
	testMode := "check_health"
	testMessage := "Test Message"
	req := &api.SdkFilesystemCheckStartRequest{
		VolumeId: testVolumeId,
		Mode:     testMode,
	}

	testMockResp := &api.SdkFilesystemCheckStartResponse{
		Status:  testStatus,
		Message: testMessage,
	}

	// Create response
	s.MockDriver().
		EXPECT().
		FilesystemCheckStart(&api.SdkFilesystemCheckStartRequest{
			VolumeId: testVolumeId,
			Mode:     testMode,
		}).
		Return(testMockResp, nil).
		Times(1)

	// Setup client
	c := api.NewOpenStorageFilesystemCheckClient(s.Conn())

	// Get info
	resp, err := c.Start(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, resp.Status, testMockResp.Status)
	assert.Equal(t, resp.Message, testMockResp.Message)
}

func TestSdkFilesystemCheckCheckHealthGetStatus(t *testing.T) {
	// Create server and client connection
	s := newTestServer(t)
	defer s.Stop()

	testVolumeId := "myvol"
	testStatus := api.FilesystemCheck_FS_CHECK_INPROGRESS
	testMessage := "Test Message : FSCheck in progress"
	testMode := "check_health"
	req := &api.SdkFilesystemCheckGetStatusRequest{
		VolumeId: testVolumeId,
	}

	testMockResp := &api.SdkFilesystemCheckGetStatusResponse{
		Status:  testStatus,
		Mode:    testMode,
		Message: testMessage,
	}

	// Create response
	s.MockDriver().
		EXPECT().
		FilesystemCheckGetStatus(&api.SdkFilesystemCheckGetStatusRequest{
			VolumeId: testVolumeId,
		}).
		Return(testMockResp, nil).
		Times(1)

	// Setup client
	c := api.NewOpenStorageFilesystemCheckClient(s.Conn())

	// Get info
	resp, err := c.GetStatus(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, resp.Status, testMockResp.Status)
	assert.Equal(t, resp.Mode, testMockResp.Mode)
	assert.Equal(t, resp.Message, testMockResp.Message)
}

func TestSdkFilesystemCheckFixAll(t *testing.T) {
	// Create server and client connection
	s := newTestServer(t)
	defer s.Stop()

	testVolumeId := "myvol"
	testStatus := api.FilesystemCheck_FS_CHECK_STARTED
	testMode := "fix_all"
	testMessage := "Test Message"
	req := &api.SdkFilesystemCheckStartRequest{
		VolumeId: testVolumeId,
		Mode:     testMode,
	}

	testMockResp := &api.SdkFilesystemCheckStartResponse{
		Status:  testStatus,
		Message: testMessage,
	}

	// Create response
	s.MockDriver().
		EXPECT().
		FilesystemCheckStart(&api.SdkFilesystemCheckStartRequest{
			VolumeId: testVolumeId,
			Mode:     testMode,
		}).
		Return(testMockResp, nil).
		Times(1)

	// Setup client
	c := api.NewOpenStorageFilesystemCheckClient(s.Conn())

	// Get info
	resp, err := c.Start(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, resp.Status, testMockResp.Status)
	assert.Equal(t, resp.Message, testMockResp.Message)
}

func TestSdkFilesystemCheckFixAllGetStatus(t *testing.T) {
	// Create server and client connection
	s := newTestServer(t)
	defer s.Stop()

	testVolumeId := "myvol"
	testStatus := api.FilesystemCheck_FS_CHECK_INPROGRESS
	testMessage := "Test Message : FSCheck in progress"
	testMode := "fix_all"
	req := &api.SdkFilesystemCheckGetStatusRequest{
		VolumeId: testVolumeId,
	}

	testMockResp := &api.SdkFilesystemCheckGetStatusResponse{
		Status:  testStatus,
		Mode:    testMode,
		Message: testMessage,
	}

	// Create response
	s.MockDriver().
		EXPECT().
		FilesystemCheckGetStatus(&api.SdkFilesystemCheckGetStatusRequest{
			VolumeId: testVolumeId,
		}).
		Return(testMockResp, nil).
		Times(1)

	// Setup client
	c := api.NewOpenStorageFilesystemCheckClient(s.Conn())

	// Get info
	resp, err := c.GetStatus(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, resp.Status, testMockResp.Status)
	assert.Equal(t, resp.Mode, testMockResp.Mode)
	assert.Equal(t, resp.Message, testMockResp.Message)
}

func TestSdkFilesystemCheckFixSafe(t *testing.T) {
	// Create server and client connection
	s := newTestServer(t)
	defer s.Stop()

	testVolumeId := "myvol"
	testStatus := api.FilesystemCheck_FS_CHECK_STARTED
	testMode := "fix_safe"
	testMessage := "Test Message"
	req := &api.SdkFilesystemCheckStartRequest{
		VolumeId: testVolumeId,
		Mode:     testMode,
	}

	testMockResp := &api.SdkFilesystemCheckStartResponse{
		Status:  testStatus,
		Message: testMessage,
	}

	// Create response
	s.MockDriver().
		EXPECT().
		FilesystemCheckStart(&api.SdkFilesystemCheckStartRequest{
			VolumeId: testVolumeId,
			Mode:     testMode,
		}).
		Return(testMockResp, nil).
		Times(1)

	// Setup client
	c := api.NewOpenStorageFilesystemCheckClient(s.Conn())

	// Get info
	resp, err := c.Start(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, resp.Status, testMockResp.Status)
	assert.Equal(t, resp.Message, testMockResp.Message)
}

func TestSdkFilesystemCheckFixSafeGetStatus(t *testing.T) {
	// Create server and client connection
	s := newTestServer(t)
	defer s.Stop()

	testVolumeId := "myvol"
	testStatus := api.FilesystemCheck_FS_CHECK_INPROGRESS
	testMessage := "Test Message : FSCheck in progress"
	testMode := "fix_safe"
	req := &api.SdkFilesystemCheckGetStatusRequest{
		VolumeId: testVolumeId,
	}

	testMockResp := &api.SdkFilesystemCheckGetStatusResponse{
		Status:  testStatus,
		Mode:    testMode,
		Message: testMessage,
	}

	// Create response
	s.MockDriver().
		EXPECT().
		FilesystemCheckGetStatus(&api.SdkFilesystemCheckGetStatusRequest{
			VolumeId: testVolumeId,
		}).
		Return(testMockResp, nil).
		Times(1)

	// Setup client
	c := api.NewOpenStorageFilesystemCheckClient(s.Conn())

	// Get info
	resp, err := c.GetStatus(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, resp.Status, testMockResp.Status)
	assert.Equal(t, resp.Mode, testMockResp.Mode)
	assert.Equal(t, resp.Message, testMockResp.Message)
}

func TestSdkFilesystemCheckStop(t *testing.T) {

	// Create server and client connection
	s := newTestServer(t)
	defer s.Stop()

	testVolumeId := "myvol"
	req := &api.SdkFilesystemCheckStopRequest{
		VolumeId: testVolumeId,
	}

	testMockResp := &api.SdkFilesystemCheckStopResponse{}

	// Create response
	s.MockDriver().
		EXPECT().
		FilesystemCheckStop(&api.SdkFilesystemCheckStopRequest{
			VolumeId: testVolumeId,
		}).
		Return(testMockResp, nil).
		Times(1)
	// Setup client
	c := api.NewOpenStorageFilesystemCheckClient(s.Conn())

	// Get info
	_, err := c.Stop(context.Background(), req)
	assert.NoError(t, err)
}
