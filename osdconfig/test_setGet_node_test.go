package osdconfig

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestSetGetNode(t *testing.T) {
	// create in memory kvdb
	kv, err := newInMemKvdb()
	if err != nil {
		t.Fatal(err)
	}

	// get new config manager using handle to kvdb
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	manager, err := NewManager(ctx, kv)
	if err != nil {
		t.Fatal(err)
	}
	defer manager.Close()

	// prepare expected cluster config
	expectedConf := new(NodeConfig)
	expectedConf.NodeId = "myNodeID"
	expectedConf.Storage = new(StorageConfig)
	expectedConf.Storage.Devices = []string{"dev1", "dev2"}

	// set the expected cluster config value
	if err := manager.SetNodeConf(expectedConf); err != nil {
		t.Fatal(err)
	}

	// get the cluster config value
	receivedConf, err := manager.GetNodeConf(expectedConf.NodeId)
	if err != nil {
		t.Fatal(err)
	}

	// compare expected and received
	if !reflect.DeepEqual(expectedConf, receivedConf) {
		t.Fatal("expected and received values are not deep equal")
	}
}
