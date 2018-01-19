package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/portworx/kvdb"
	"github.com/sdeoras/openstorage/osdconfig"
	"github.com/sdeoras/openstorage/osdconfig/proto"
	"golang.org/x/net/context"
)

type MyKVObj struct {
	kv kvdb.Kvdb
}

func (m *MyKVObj) Handler() kvdb.Kvdb {
	return m.kv
}

func TestKV(t *testing.T) {
	config := new(proto.Config)
	config.Description = "this is description text"
	config.Global = new(proto.GlobalConfig)
	config.Global.AlertingUrl = "this is alerting url"

	options := make(map[string]string)
	options["KvUseInterface"] = ""
	kv, err := kvdb.New("pwx/test", "", nil, options, nil)
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})
	go func(c chan struct{}) {
		client := osdconfig.NewKVConnection(&MyKVObj{kv})

		ack, err := client.Set(context.Background(), config)
		if err != nil {
			t.Fatal(err)
		}

		t.Log("Bytes written:", ack.N)

		c <- struct{}{}
	}(done)
	<-done

	go func(c chan struct{}) {
		file, err := os.Open(ConfigFile)
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		client := osdconfig.NewKVConnection(&MyKVObj{kv})
		config, err := client.Get(context.Background(), &proto.Empty{})
		if err != nil {
			t.Fatal(err)
		}

		jb, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(string(jb))
		c <- struct{}{}
	}(done)
	<-done

}
