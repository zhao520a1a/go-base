package util

import (
	"context"
	"gitlab.pri.ibanyu.com/middleware/seaweed/xlog"
	"reflect"
)

func CompareObject(ctx context.Context, x, y interface{}, fun string, params string) {
	if !reflect.DeepEqual(x, y) {
		xlog.Infof(ctx, "%s params: %s Not Equal %v - %v ", fun, params, x, y)
	}
}

func CompareInt64(ctx context.Context, x, y int64, fun string, params string) {
	if x != y {
		xlog.Infof(ctx, "%s params: %s Not Equal %v - %v ", fun, params, x, y)
	}
}
