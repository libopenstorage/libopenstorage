package alert

import (
	"github.com/libopenstorage/openstorage/api"
	"github.com/portworx/kvdb"
	"github.com/portworx/kvdb/mem"
	"github.com/stretchr/testify/require"
	"go.pedge.io/proto/time"
	"strconv"
	"testing"
	"time"
)

var (
	kva             AlertClient
	nextId          int64
	isWatcherCalled int
	watcherAction   api.AlertActionType
	watcherAlert    api.Alert
	watcherPrefix   string
	watcherKey      string
)

const (
	kvdbDomain = "openstorage"
	clusterName = "1"
)

func TestSetup(t *testing.T) {
	kv := kvdb.Instance()
	if kv == nil {
		kv, err := kvdb.New(mem.Name, kvdbDomain+"/"+clusterName, []string{}, nil)
		if err != nil {
			t.Fatalf("Failed to set default KV store : (%v): %v", mem.Name, err)
		}
		err = kvdb.SetInstance(kv)
		if err != nil {
			t.Fatalf("Failed to set default KV store: (%v): %v", mem.Name, err)
		}
	}

	var err error
	kva, err = New("alert_kvdb", mem.Name, kvdbDomain, []string{}, clusterName)
	if err != nil {
		t.Fatalf("Failed to create new Kvapi.Alert object")
	}
}

func TestRaiseAndErase(t *testing.T) {
	// Raise api.Alert Id : 1

	raiseAlert, err := kva.Raise(api.Alert{Resource: api.ResourceType_RESOURCE_TYPE_VOLUME, Severity: api.SeverityType_SEVERITY_TYPE_NOTIFY, Message: "Test Message"})
	require.NoError(t, err, "Failed in raising an alert")

	kv := kva.GetKvdbInstance()
	var alert api.Alert

	_, err = kv.GetVal(getResourceKey(api.ResourceType_RESOURCE_TYPE_VOLUME)+strconv.FormatInt(raiseAlert.Id, 10), &alert)
	require.NoError(t, err, "Failed to retrieve alert from kvdb")
	require.NotNil(t, alert, "api.Alert object null in kvdb")
	require.Equal(t, raiseAlert.Id, alert.Id, "api.Alert Id mismatch")
	require.Equal(t, api.ResourceType_RESOURCE_TYPE_VOLUME, alert.Resource, "api.Alert Resource mismatch")
	require.Equal(t, api.SeverityType_SEVERITY_TYPE_NOTIFY, alert.Severity, "api.Alert Severity mismatch")

	// Raise api.Alert with no Resource
	_, err = kva.Raise(api.Alert{Severity: api.SeverityType_SEVERITY_TYPE_NOTIFY})
	require.Error(t, err, "An error was expected")
	require.Equal(t, ErrResourceNotFound, err, "Error mismatch")

	// Erase api.Alert Id : 1
	err = kva.Erase(api.ResourceType_RESOURCE_TYPE_VOLUME, raiseAlert.Id)
	require.NoError(t, err, "Failed to erase an alert")

	_, err = kv.GetVal(getResourceKey(api.ResourceType_RESOURCE_TYPE_VOLUME)+"1", &alert)
	require.Error(t, err, "api.Alert not erased from kvdb")
}

func TestRetrieve(t *testing.T) {
	var alert api.Alert

	// Raise a ResourceType_RESOURCE_TYPE_NODE specific api.Alert
	raiseAlert, err := kva.Raise(api.Alert{Resource: api.ResourceType_RESOURCE_TYPE_NODE, Severity: api.SeverityType_SEVERITY_TYPE_ALARM})

	alert, err = kva.Retrieve(api.ResourceType_RESOURCE_TYPE_NODE, raiseAlert.Id)
	require.NoError(t, err, "Failed to retrieve alert")
	require.NotNil(t, alert, "api.Alert object null")
	require.Equal(t, raiseAlert.Id, alert.Id, "api.Alert Id mismatch")
	require.Equal(t, api.ResourceType_RESOURCE_TYPE_NODE, alert.Resource, "api.Alert resource mismatch")
	require.Equal(t, api.SeverityType_SEVERITY_TYPE_ALARM, alert.Severity, "api.Alert severity mismatch")

	// Retrieve non existing alert
	alert, err = kva.Retrieve(api.ResourceType_RESOURCE_TYPE_VOLUME, 5)
	require.Error(t, err, "Expected an error")

	// Cleanup
	err = kva.Erase(api.ResourceType_RESOURCE_TYPE_NODE, raiseAlert.Id)
}

func TestClear(t *testing.T) {
	// Raise an alert
	var alert api.Alert
	kv := kva.GetKvdbInstance()
	raiseAlert, err := kva.Raise(api.Alert{Resource: api.ResourceType_RESOURCE_TYPE_NODE, Severity: api.SeverityType_SEVERITY_TYPE_ALARM})

	err = kva.Clear(api.ResourceType_RESOURCE_TYPE_NODE, raiseAlert.Id)
	require.NoError(t, err, "Failed to clear alert")

	_, err = kv.GetVal(getResourceKey(api.ResourceType_RESOURCE_TYPE_NODE)+strconv.FormatInt(raiseAlert.Id, 10), &alert)
	require.Equal(t, true, alert.Cleared, "Failed to clear alert")

	err = kva.Erase(api.ResourceType_RESOURCE_TYPE_NODE, raiseAlert.Id)
}

func TestEnumerateAlert(t *testing.T) {
	// Raise a few alert
	raiseAlert1, err := kva.Raise(api.Alert{Resource: api.ResourceType_RESOURCE_TYPE_VOLUME, Severity: api.SeverityType_SEVERITY_TYPE_NOTIFY})
	raiseAlert2, err := kva.Raise(api.Alert{Resource: api.ResourceType_RESOURCE_TYPE_VOLUME, Severity: api.SeverityType_SEVERITY_TYPE_NOTIFY})
	raiseAlert3, err := kva.Raise(api.Alert{Resource: api.ResourceType_RESOURCE_TYPE_VOLUME, Severity: api.SeverityType_SEVERITY_TYPE_WARNING})
	raiseAlert4, err := kva.Raise(api.Alert{Resource: api.ResourceType_RESOURCE_TYPE_NODE, Severity: api.SeverityType_SEVERITY_TYPE_WARNING})

	enAlerts, err := kva.Enumerate(api.Alert{Resource: api.ResourceType_RESOURCE_TYPE_VOLUME})
	require.NoError(t, err, "Failed to enumerate alert")
	require.Equal(t, 3, len(enAlerts), "Enumerated incorrect number of alert")

	enAlerts, err = kva.Enumerate(api.Alert{Resource: api.ResourceType_RESOURCE_TYPE_VOLUME, Severity: api.SeverityType_SEVERITY_TYPE_WARNING})
	require.NoError(t, err, "Failed to enumerate alert")
	require.Equal(t, 1, len(enAlerts), "Enumerated incorrect number of alert")
	require.Equal(t, api.SeverityType_SEVERITY_TYPE_WARNING, enAlerts[0].Severity, "Severity mismatch")

	enAlerts, err = kva.Enumerate(api.Alert{})
	require.NoError(t, err, "Failed to enumerate alert")
	require.Equal(t, 4, len(enAlerts), "Enumerated incorrect number of alert")

	enAlerts, err = kva.Enumerate(api.Alert{Severity: api.SeverityType_SEVERITY_TYPE_WARNING})
	require.NoError(t, err, "Failed to enumerate alert")
	require.Equal(t, 2, len(enAlerts), "Enumerated incorrect number of alert")
	require.Equal(t, api.SeverityType_SEVERITY_TYPE_WARNING, enAlerts[0].Severity, "Severity mismatch")

	// Add a dummy event into kvdb two hours ago
	kv := kva.GetKvdbInstance()
	currentTime := time.Now()
	delayedTime := currentTime.Add(-1 * time.Duration(2) * time.Hour)

	alert := api.Alert{Timestamp: prototime.TimeToTimestamp(delayedTime), Id: 10, Resource: api.ResourceType_RESOURCE_TYPE_VOLUME}

	_, err = kv.Put(getResourceKey(api.ResourceType_RESOURCE_TYPE_VOLUME)+strconv.FormatInt(10, 10), &alert, 0)
	enAlerts, err = kva.EnumerateWithinTimeRange(currentTime.Add(-1*time.Duration(10)*time.Second), currentTime, api.ResourceType_RESOURCE_TYPE_VOLUME)
	require.NoError(t, err, "Failed to enumerate results")
	require.Equal(t, 3, len(enAlerts), "Enumerated incorrect number of alert")

	err = kva.Erase(api.ResourceType_RESOURCE_TYPE_VOLUME, raiseAlert1.Id)
	err = kva.Erase(api.ResourceType_RESOURCE_TYPE_VOLUME, raiseAlert2.Id)
	err = kva.Erase(api.ResourceType_RESOURCE_TYPE_VOLUME, raiseAlert3.Id)
	err = kva.Erase(api.ResourceType_RESOURCE_TYPE_NODE, raiseAlert4.Id)
}

func testAlertWatcher(alert *api.Alert, action api.AlertActionType, prefix string, key string) error {
	// A dummy callback function
	// Setting the global variables so that we can check them in our unit tests
	isWatcherCalled = 1
	if action != api.AlertActionType_ALERT_ACTION_TYPE_DELETE {
		watcherAlert = *alert
	} else {
		watcherAlert = api.Alert{}
	}
	watcherAction = action
	watcherPrefix = prefix
	watcherKey = key
	return nil

}

func TestWatch(t *testing.T) {
	isWatcherCalled = 0

	err := kva.Watch(clusterName, testAlertWatcher)
	require.NoError(t, err, "Failed to subscribe a watch function")

	raiseAlert1, err := kva.Raise(api.Alert{Resource: api.ResourceType_RESOURCE_TYPE_CLUSTER, Severity: api.SeverityType_SEVERITY_TYPE_NOTIFY})

	// Sleep for sometime so that we pass on some previous watch callbacks
	time.Sleep(time.Millisecond * 100)

	require.Equal(t, 1, isWatcherCalled, "Callback function not called")
	require.Equal(t, api.AlertActionType_ALERT_ACTION_TYPE_CREATE, watcherAction, "action mismatch for create")
	require.Equal(t, raiseAlert1.Id, watcherAlert.Id, "alert id mismatch")
	require.Equal(t, "alert/cluster/"+strconv.FormatInt(raiseAlert1.Id, 10), watcherKey, "key mismatch")

	err = kva.Clear(api.ResourceType_RESOURCE_TYPE_CLUSTER, raiseAlert1.Id)

	// Sleep for sometime so that we pass on some previous watch callbacks
	time.Sleep(time.Millisecond * 100)

	require.Equal(t, api.AlertActionType_ALERT_ACTION_TYPE_UPDATE, watcherAction, "action mismatch for update")
	require.Equal(t, "alert/cluster/"+strconv.FormatInt(raiseAlert1.Id, 10), watcherKey, "key mismatch")

	err = kva.Erase(api.ResourceType_RESOURCE_TYPE_CLUSTER, raiseAlert1.Id)

	// Sleep for sometime so that we pass on some previous watch callbacks
	time.Sleep(time.Millisecond * 100)

	require.Equal(t, api.AlertActionType_ALERT_ACTION_TYPE_DELETE, watcherAction, "action mismatch for delete")
	require.Equal(t, "alert/cluster/"+strconv.FormatInt(raiseAlert1.Id, 10), watcherKey, "key mismatch")

	// Watch on a new clusterID
	newClusterId := "2"
	isWatcherCalled = 0
	err = kva.Watch(newClusterId, testAlertWatcher)

	// Create a new alert instance for raising an alert in this new cluster id
	kvaNew, err := New("alert_kvdb_test", mem.Name, kvdbDomain, []string{}, newClusterId)
	if err != nil {
		t.Fatalf("Failed to create new Kvapi.Alert object %s", err.Error())
	}

	raiseAlertNew, err := kvaNew.Raise(api.Alert{Resource: api.ResourceType_RESOURCE_TYPE_NODE, Severity: api.SeverityType_SEVERITY_TYPE_ALARM})
	// Sleep for sometime so that we pass on some previous watch callbacks
	time.Sleep(time.Millisecond * 100)

	require.Equal(t, 1, isWatcherCalled, "Callback function not called")
	require.Equal(t, api.AlertActionType_ALERT_ACTION_TYPE_CREATE, watcherAction, "action mismatch for create")
	require.Equal(t, raiseAlertNew.Id, watcherAlert.Id, "alert id mismatch")
	require.Equal(t, "alert/node/"+strconv.FormatInt(raiseAlertNew.Id, 10), watcherKey, "key mismatch")
}


func mockGenerateId(clusterId string) (int64, error) {
	nextId = nextId + 1
	return nextId, nil
}
