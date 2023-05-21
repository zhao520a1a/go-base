package bpm

import (
	"context"
)

type NotificationMode int32

const (
	NotificationMode_DEFAULT   NotificationMode = 0
	NotificationMode_EMAIL     NotificationMode = 1
	NotificationMode_DING_TALK NotificationMode = 2
)

var NotificationMode_name = map[int32]string{
	0: "DEFAULT",
	1: "EMAIL",
	2: "DING_TALK",
}

var NotificationMode_value = map[string]int32{
	"DEFAULT":   0,
	"EMAIL":     1,
	"DING_TALK": 2,
}

type Notification struct {
	To      []string
	Title   string
	Content string
	OpUser  string
	Link    string
	Mode    NotificationMode
}

type Notifier interface {
	Notify(ctx context.Context, notification *Notification) error
}
