package xconfig

import "context"

// Driver implements by each config center(such as apollo)
// config center should call 'Register' to register itself
type Driver interface {
	// New ...
	New(ctx context.Context, serviceName string, namespaceNames []string, options ...Option) (ConfigCenter, error)
}
