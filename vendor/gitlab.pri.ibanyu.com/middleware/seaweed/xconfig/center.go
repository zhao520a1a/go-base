package xconfig

import (
	"context"
	"fmt"
)

// Option ...
type Option func(ConfigCenter)

// ConfigerType ...
type ConfigerType string

type ConfigCenter interface {
	// RegisterObserver register observer return recall func to cancel observer
	RegisterObserver(ctx context.Context, observer *ConfigObserver) (recall func())
	// Stop stop client include cancel client ctx, cancel longpoller ctx, close updateChan
	Stop(ctx context.Context) error
	// SubscribeNamespaces subscibe new namespaces if init not set
	SubscribeNamespaces(ctx context.Context, namespaceNames []string) error
	// GetString get string value form default namespace application
	GetString(ctx context.Context, key string) (string, bool)
	// GetStringWithNamespace get string value form specified namespace
	GetStringWithNamespace(ctx context.Context, namespace, key string) (string, bool)
	// GetBool get bool value form default namespace application
	GetBool(ctx context.Context, key string) (bool, bool)
	// GetBoolWithNamespace get bool value form specified namespace
	GetBoolWithNamespace(ctx context.Context, namespace, key string) (bool, bool)
	// GetInt get int value form default namespace application
	GetInt(ctx context.Context, key string) (int, bool)
	// GetIntWithNamespace get int value form specified namespace
	GetIntWithNamespace(ctx context.Context, namespace, key string) (int, bool)
	// GetAllKeys get all keys from default namespace application
	GetAllKeys(ctx context.Context) []string
	// GetAllKeysWithNamespace get all keys form specified namespace
	GetAllKeysWithNamespace(ctx context.Context, namespace string) []string
	// Unmarshal unmarshal from default namespace application
	Unmarshal(ctx context.Context, v interface{}) error
	// UnmarshalWithNamespace unmarshal form specified namespace
	UnmarshalWithNamespace(ctx context.Context, namespace string, v interface{}) error
	// UnmarshalKey unmarshal key from default namespace application
	UnmarshalKey(ctx context.Context, key string, v interface{}) error
	// UnmarshalKeyWithNamespace unmarshal key form specified namespace
	UnmarshalKeyWithNamespace(ctx context.Context, namespace string, key string, v interface{}) error
	// SetCluster set center cluster
	SetCluster(cluster string)
	// SetCacheDir set center cache dir
	SetCacheDir(cacheDir string)
	// SetIPHost set center remote host
	SetIPHost(ipHost string)
}

// NewConfigCenter ...
func NewConfigCenter(ctx context.Context, ctype ConfigerType, serviceName string, namespaceNames []string, options ...Option) (ConfigCenter, error) {
	driver, err := GetDriver(ctype)
	if err != nil {
		return nil, fmt.Errorf("new config center err:%s", err.Error())
	}
	return driver.New(ctx, serviceName, namespaceNames, options...)
}
