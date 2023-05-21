package config

import "context"

type ConfigService interface {
	GetNotifierEmailFlag(ctx context.Context) (bool, error)
	GetNotNotifyConfigForEmail(ctx context.Context) ([]string, error)
}
