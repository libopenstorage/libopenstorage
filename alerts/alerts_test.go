package alerts

import (
	"testing"

	"github.com/libopenstorage/openstorage/api"
	"github.com/portworx/kvdb"
	"github.com/portworx/kvdb/mem"
)

// helper function go get a new kvdb instance
func newInMemKvdb() (kvdb.Kvdb, error) {
	// create in memory kvdb
	if kv, err := kvdb.New(mem.Name, "", []string{}, nil, nil); err != nil {
		return nil, err
	} else {
		return kv, nil
	}
}

func TestManager_Raise(t *testing.T) {
	kv, err := newInMemKvdb()
	if err != nil {
		t.Fatal(err)
	}

	manager := NewManager(kv)
	var alert *api.Alert

	alert = new(api.Alert)
	alert.AlertType = 10
	alert.Resource = api.ResourceType_RESOURCE_TYPE_VOLUME
	alert.ResourceId = "inca"
	if err := manager.Raise(alert); err != nil {
		t.Fatal(err)
	}

	alert = new(api.Alert)
	alert.AlertType = 12
	alert.Resource = api.ResourceType_RESOURCE_TYPE_CLUSTER
	alert.ResourceId = "aztec"
	if err := manager.Raise(alert); err != nil {
		t.Fatal(err)
	}

	alert = new(api.Alert)
	alert.AlertType = 10
	alert.Resource = api.ResourceType_RESOURCE_TYPE_DRIVE
	alert.ResourceId = "maya"
	if err := manager.Raise(alert); err != nil {
		t.Fatal(err)
	}

	alert = new(api.Alert)
	alert.AlertType = 10
	alert.Resource = api.ResourceType_RESOURCE_TYPE_DRIVE
	alert.ResourceId = "inca"
	if err := manager.Raise(alert); err != nil {
		t.Fatal(err)
	}

	alert = new(api.Alert)
	alert.AlertType = 14
	alert.Resource = api.ResourceType_RESOURCE_TYPE_DRIVE
	alert.ResourceId = "aztec"
	if err := manager.Raise(alert); err != nil {
		t.Fatal(err)
	}

	alert = new(api.Alert)
	alert.AlertType = 12
	alert.Resource = api.ResourceType_RESOURCE_TYPE_DRIVE
	alert.ResourceId = "maya"
	if err := manager.Raise(alert); err != nil {
		t.Fatal(err)
	}

	// prepare a test configuration table
	configs := []struct {
		name          string
		filters       []Filter
		expectedCount int
	}{
		{
			name:          "by none",
			expectedCount: 6,
		},
		{
			name: "by 1 resource type",
			filters: []Filter{
				NewResourceTypeFilter(api.ResourceType_RESOURCE_TYPE_VOLUME),
			},
			expectedCount: 1,
		},
		{
			name: "by 2 resource types",
			filters: []Filter{
				NewResourceTypeFilter(api.ResourceType_RESOURCE_TYPE_VOLUME),
				NewResourceTypeFilter(api.ResourceType_RESOURCE_TYPE_CLUSTER),
			},
			expectedCount: 2,
		},
		{
			name: "by 2 resource types",
			filters: []Filter{
				NewResourceTypeFilter(api.ResourceType_RESOURCE_TYPE_DRIVE),
				NewResourceTypeFilter(api.ResourceType_RESOURCE_TYPE_CLUSTER),
			},
			expectedCount: 5,
		},
		{
			name: "by 1 resource id",
			filters: []Filter{
				NewResourceIDFilter("inca", api.ResourceType_RESOURCE_TYPE_VOLUME),
			},
			expectedCount: 1,
		},
		{
			name: "by 2 resource ids",
			filters: []Filter{
				NewResourceIDFilter("inca", api.ResourceType_RESOURCE_TYPE_VOLUME),
				NewResourceIDFilter("maya", api.ResourceType_RESOURCE_TYPE_DRIVE),
			},
			expectedCount: 3,
		},
		{
			name: "by 2 different filter types",
			filters: []Filter{
				NewResourceTypeFilter(api.ResourceType_RESOURCE_TYPE_VOLUME),
				NewResourceIDFilter("maya", api.ResourceType_RESOURCE_TYPE_DRIVE),
			},
			expectedCount: 3,
		},
	}

	for _, config := range configs {
		myAlerts, err := manager.Enumerate(config.filters...)
		if err != nil {
			t.Fatal(err)
		}

		if len(myAlerts) != config.expectedCount {
			t.Fatal("test:", config.name, ", alert count: expected:", config.expectedCount, ", found:", len(myAlerts))
		}
	}
}
