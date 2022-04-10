package manager

import (
	"context"
	"fmt"

	"gitlab.pri.ibanyu.com/middleware/seaweed/internal/dbrouter"
)

func parseKey(key string) (instance string) {

	return key
}

func generateKey(instance string) string {

	return instance
}

func factory(ctx context.Context, insName, group string, configer dbrouter.Configer, dynamicConfiger *DynamicConfiger) (in dbrouter.Instancer, err error) {
	instance := parseKey(insName)
	config := configer.GetInstanceConfig(ctx, instance, group)
	if len(config.DBAddr) == 0 {

		return nil, fmt.Errorf("config.DBAddr err, key: %s", insName)
	}

	return NewDBInstance(insName, group, config, dynamicConfiger)
}
