package xconfig

import (
	"fmt"
	"sort"
	"sync"
)

var (
	driversMu sync.RWMutex
	drivers   = make(map[ConfigerType]Driver)
)

// Register ...
func Register(ctype ConfigerType, driver Driver) {
	driversMu.Lock()
	defer driversMu.Unlock()

	if driver == nil {
		panic("xconfig: driver is nil")
	}

	if _, dup := drivers[ctype]; dup {
		panic("xconfig: driver can called Register only once")
	}

	drivers[ctype] = driver
}

// Drivers returns a sorted list of the names of the registered driver.
func Drivers() []string {
	driversMu.RLock()
	defer driversMu.RUnlock()

	var list []string
	for name := range drivers {
		list = append(list, string(name))
	}

	sort.Strings(list)
	return list
}

// GetDriver returns a driver implement by config type.
func GetDriver(ctype ConfigerType) (Driver, error) {
	driversMu.RLock()
	driveri, ok := drivers[ctype]
	driversMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("unkown config type:%s", string(ctype))
	}
	return driveri, nil
}
