package mock

import (
	"context"
	"github.com/zhao520a1a/go-base.git/unittest/mock/bpm"
)

type Notifier struct {
	NotifyFn           func(ctx context.Context, notif *bpm.Notification) error
	NotifyInvoked      bool
	NotifyInvokedTimes int
	LastNotification   *bpm.Notification
}

func (m *Notifier) Notify(ctx context.Context, notif *bpm.Notification) error {
	m.NotifyInvoked = true
	m.NotifyInvokedTimes += 1
	m.LastNotification = notif
	return m.NotifyFn(ctx, notif)
}
