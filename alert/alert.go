package alert

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/libopenstorage/openstorage/api"
	"github.com/portworx/kvdb"
	"go.pedge.io/dlog"
)

var (
	// ErrNotSupported implemenation of a specific function is not supported.
	ErrNotSupported = errors.New("implementation not supported")
	// ErrNotFound raised if Key is not found.
	ErrNotFound = errors.New("Key not found")
	// ErrExist raised if key already exists.
	ErrExist = errors.New("Key already exists")
	// ErrUnmarshal raised if Get fails to unmarshal value.
	ErrUnmarshal = errors.New("Failed to unmarshal value")
	// ErrIllegal raised if object is not valid.
	ErrIllegal = errors.New("Illegal operation")
	// ErrNotInitialized raised if alert not initialized.
	ErrNotInitialized = errors.New("Alert not initialized")
	// ErrAlertClientNotFound raised if no client implementation found.
	ErrAlertClientNotFound = errors.New("Alert client not found")
	// ErrResourceNotFound raised if ResourceType is not found>
	ErrResourceNotFound = errors.New("Resource not found in Alert")

	instances = make(map[string]AlertClient)
	drivers   = make(map[string]InitFunc)

	lock sync.RWMutex
)

// InitFunc initialization function for alert.
type InitFunc func(string, string, []string, string) (AlertClient, error)

// AlertWatcherFunc is a function type used as a callback for KV WatchTree.
type AlertWatcherFunc func(*api.Alert, api.AlertActionType, string, string) error

// AlertClient interface for Alert API.
type AlertClient interface {
	fmt.Stringer

	// Shutdown.
	Shutdown()

	// GetKvdbInstance.
	GetKvdbInstance() kvdb.Kvdb

	// Raise raises an Alert.
	Raise(alert api.Alert) (api.Alert, error)

	// Retrieve retrieves specific Alert.
	Retrieve(resourceType api.ResourceType, id int64) (api.Alert, error)

	// Enumerate enumerates Alert.
	Enumerate(filter api.Alert) ([]*api.Alert, error)

	// EnumerateWithinTimeRange enumerates Alert between timeStart and timeEnd.
	EnumerateWithinTimeRange(timeStart time.Time, timeEnd time.Time, resourceType api.ResourceType) ([]*api.Alert, error)

	// Erase erases an Alert.
	Erase(resourceType api.ResourceType, alertID int64) error

	// Clear an Alert.
	Clear(resourceType api.ResourceType, alertID int64) error

	// Watch on all Alert>
	Watch(clusterId string, alertWatcher AlertWatcherFunc) error
}

type AlertInstance interface {
	// Clear clears an alert.
	Clear(resourceType api.ResourceType, resourceId string, alertID int64)

	// Alarm raises an alert with severity : ALARM.
	Alarm(name string, msg string, resourceType api.ResourceType, resourceId string) (int64, error)

	// Notify raises an alert with severity : NOTIFY.
	Notify(name string, msg string, resourceType api.ResourceType, resourceId string) (int64, error)

	// Warn raises an alert with severity : WARNING.
	Warn(name string, msg string, resourceType api.ResourceType, resourceId string) (int64, error)

	// Alert :  Keeping this function for backward compatibility
	// until we remove all calls to this function.
	Alert(name string, msg string) error
}

// Shutdown the alert instance.
func Shutdown() {
	lock.Lock()
	defer lock.Unlock()
	for _, v := range instances {
		v.Shutdown()
	}
}

// Get an alert instance.
func Get(name string) (AlertClient, error) {
	lock.RLock()
	defer lock.RUnlock()

	if v, ok := instances[name]; ok {
		return v, nil
	}
	return nil, ErrAlertClientNotFound
}

// New returns a new alert instance.
func New(name string, kvdbName string, kvdbBase string, kvdbMachines []string, clusterId string) (AlertClient, error) {
	lock.Lock()
	defer lock.Unlock()

	if _, ok := instances[name]; ok {
		return nil, ErrExist
	}
	if initFunc, exists := drivers[name]; exists {
		driver, err := initFunc(kvdbName, kvdbBase, kvdbMachines, clusterId)
		if err != nil {
			return nil, err
		}
		instances[name] = driver
		return driver, err
	}
	return nil, ErrNotSupported
}

// NewAlertInstance creates a new singleton istance of AlertInstance.
func NewAlertInstance(version, nodeId, clusterId, kvdbName, kvdbBase string, kvdbMachines []string) {
	kva, err := Get(Name)
	if err != nil {
		kva, err = New(Name, kvdbName, kvdbBase, kvdbMachines, clusterId)
		if err != nil {
			dlog.Errorf("Failed to initialize an AlertInstance ")
		}
	}
	newAlertInstance(nodeId, clusterId, version, kva)
}

// Instance returns the singleton AlertInstance.
func Instance() AlertInstance {
	return instance()
}

// Register an alert interface.
func Register(name string, initFunc InitFunc) error {
	lock.Lock()
	defer lock.Unlock()
	if _, exists := drivers[name]; exists {
		return ErrExist
	}
	drivers[name] = initFunc
	return nil
}
