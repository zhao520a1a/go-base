package mock

import "context"

type ConfigService struct {
	GetNotifierEmailFlagFn       func(ctx context.Context) (bool, error)
	GetNotNotifyConfigForEmailFn func(ctx context.Context) ([]string, error)
}

func (m *ConfigService) GetNotifierEmailFlag(ctx context.Context) (bool, error) {
	return m.GetNotifierEmailFlagFn(ctx)
}

func (m *ConfigService) GetNotNotifyConfigForEmail(ctx context.Context) ([]string, error) {
	return m.GetNotNotifyConfigForEmailFn(ctx)
}
