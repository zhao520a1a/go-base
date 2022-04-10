package xcontext

import (
	"context"
	"errors"
	"time"
)

// 由于请求的上下文信息的 thrift 定义在 util 项目中，本模块主要为了避免循环依赖

//ContextCaller store caller info
type ContextCaller struct {
	Method string
}

const (
	//ContextKeyTraceID ...
	ContextKeyTraceID = "traceID"
	//ContextKeyHead ...
	ContextKeyHead = "Head"
	//ContextKeyHeadUID ...
	ContextKeyHeadUID = "uid"
	//ContextKeyHeadSource ...
	ContextKeyHeadSource = "source"
	//ContextKeyHeadIP ...
	ContextKeyHeadIP = "ip"
	//ContextKeyHeadRegion ...
	ContextKeyHeadRegion = "region"
	//ContextKeyHeadDt ...
	ContextKeyHeadDt = "dt"
	//ContextKeyHeadUnionID ...
	ContextKeyHeadUnionID = "unionid"
	//ContextKeyHeadDID
	ContextKeyHeadDID = "h_did"
	//ContextKeyHeadZone
	ContextKeyHeadZone = "zone"
	//ContextKeyHeadZoneName
	ContextKeyHeadZoneName = "zone_name"

	//ContextKeyControl ...
	ContextKeyControl = "Control"

	ContextKeyCaller = "Caller"
)

//DefaultGroup ...
const DefaultGroup = ""

//ErrInvalidContext ...
var ErrInvalidContext = errors.New("invalid context")

//ContextHeader ...
type ContextHeader interface {
	ToKV() map[string]interface{}
}

//ContextControlRouter ...
type ContextControlRouter interface {
	GetControlRouteGroup() (string, bool)
	SetControlRouteGroup(string) error
}

//ContextControlCaller ...
type ContextControlCaller interface {
	GetControlCallerServerName() (string, bool)
	SetControlCallerServerName(string) error
	GetControlCallerServerID() (string, bool)
	SetControlCallerServerID(string) error
	GetControlCallerMethod() (string, bool)
	SetControlCallerMethod(string) error
}

//GetControlRouteGroup ...
func GetControlRouteGroup(ctx context.Context) (group string, ok bool) {
	value := ctx.Value(ContextKeyControl)
	if value == nil {
		ok = false
		return
	}
	control, ok := value.(ContextControlRouter)
	if ok == false {
		return
	}
	return control.GetControlRouteGroup()
}

//SetControlRouteGroup ...
func SetControlRouteGroup(ctx context.Context, group string) (context.Context, error) {
	value := ctx.Value(ContextKeyControl)
	if value == nil {
		return ctx, ErrInvalidContext
	}
	control, ok := value.(ContextControlRouter)
	if !ok {
		return ctx, ErrInvalidContext
	}

	err := control.SetControlRouteGroup(group)
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, ContextKeyControl, control), nil
}

//GetControlRouteGroupWithDefault ...
func GetControlRouteGroupWithDefault(ctx context.Context, dv string) string {
	if group, ok := GetControlRouteGroup(ctx); ok {
		return group
	}
	return dv
}

func getHeaderByKey(ctx context.Context, key string) (val interface{}, ok bool) {
	head := ctx.Value(ContextKeyHead)
	if head == nil {
		ok = false
		return
	}

	var header ContextHeader
	if header, ok = head.(ContextHeader); ok {
		val, ok = header.ToKV()[key]
	}
	return
}

//GetUID ...
func GetUID(ctx context.Context) (uid int64, ok bool) {
	val, ok := getHeaderByKey(ctx, ContextKeyHeadUID)
	if ok {
		uid, ok = val.(int64)
	}
	return
}

//GetSource ...
func GetSource(ctx context.Context) (source int32, ok bool) {
	val, ok := getHeaderByKey(ctx, ContextKeyHeadSource)
	if ok {
		source, ok = val.(int32)
	}
	return
}

//GetIP ...
func GetIP(ctx context.Context) (ip string, ok bool) {
	val, ok := getHeaderByKey(ctx, ContextKeyHeadIP)
	if ok {
		ip, ok = val.(string)
	}
	return
}

//GetRegion ...
func GetRegion(ctx context.Context) (region string, ok bool) {
	val, ok := getHeaderByKey(ctx, ContextKeyHeadRegion)
	if ok {
		region, ok = val.(string)
	}
	return
}

//GetDt ...
func GetDt(ctx context.Context) (dt int32, ok bool) {
	val, ok := getHeaderByKey(ctx, ContextKeyHeadDt)
	if ok {
		dt, ok = val.(int32)
	}
	return
}

//GetUnionID ...
func GetUnionID(ctx context.Context) (unionID string, ok bool) {
	val, ok := getHeaderByKey(ctx, ContextKeyHeadUnionID)
	if ok {
		unionID, ok = val.(string)
	}
	return
}

func GetDID(ctx context.Context) (string, bool) {
	val, ok := getHeaderByKey(ctx, ContextKeyHeadDID)
	if !ok {
		return "", false
	}
	did, ok := val.(string)
	return did, ok
}

func GetZone(ctx context.Context) (int32, bool) {
	val, ok := getHeaderByKey(ctx, ContextKeyHeadZone)
	if !ok {
		return 0, false
	}
	zone, ok := val.(int32)
	return zone, ok
}

func GetZoneName(ctx context.Context) (string, bool) {
	val, ok := getHeaderByKey(ctx, ContextKeyHeadZoneName)
	if !ok {
		return "", false
	}
	zoneName, ok := val.(string)
	return zoneName, ok
}

func getControlCaller(ctx context.Context) (ContextControlCaller, error) {
	value := ctx.Value(ContextKeyControl)
	if value == nil {
		return nil, ErrInvalidContext
	}
	caller, ok := value.(ContextControlCaller)
	if !ok {
		return nil, ErrInvalidContext
	}
	return caller, nil
}

//GetControlCallerServerName ...
func GetControlCallerServerName(ctx context.Context) (serverName string, ok bool) {
	caller, ok := ctx.Value(ContextKeyControl).(ContextControlCaller)
	if !ok {
		return
	}
	return caller.GetControlCallerServerName()
}

//SetControlCallerServerName ...
func SetControlCallerServerName(ctx context.Context, serverName string) (context.Context, error) {
	caller, err := getControlCaller(ctx)
	if err != nil {
		return ctx, err
	}
	err = caller.SetControlCallerServerName(serverName)
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, ContextKeyControl, caller), nil
}

//GetControlCallerServerID ...
func GetControlCallerServerID(ctx context.Context) (serverID string, ok bool) {
	caller, ok := ctx.Value(ContextKeyControl).(ContextControlCaller)
	if !ok {
		return
	}
	return caller.GetControlCallerServerID()
}

//SetControlCallerServerID ...
func SetControlCallerServerID(ctx context.Context, serverID string) (context.Context, error) {
	caller, err := getControlCaller(ctx)
	if err != nil {
		return ctx, err
	}
	err = caller.SetControlCallerServerID(serverID)
	return context.WithValue(ctx, ContextKeyControl, caller), nil
}

//GetControlCallerMethod ...
func GetControlCallerMethod(ctx context.Context) (method string, ok bool) {
	caller, ok := ctx.Value(ContextKeyControl).(ContextControlCaller)
	if !ok {
		return
	}
	return caller.GetControlCallerMethod()
}

//SetControlCallerMethod ...
func SetControlCallerMethod(ctx context.Context, method string) (context.Context, error) {
	caller, err := getControlCaller(ctx)
	if err != nil {
		return ctx, err
	}
	err = caller.SetControlCallerMethod(method)
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, ContextKeyControl, caller), nil
}

func getCaller(ctx context.Context) ContextCaller {
	value := ctx.Value(ContextKeyCaller)
	caller, ok := value.(ContextCaller)
	if !ok {
		return ContextCaller{}
	}

	return caller
}

// SetCallerMethod ...
func SetCallerMethod(ctx context.Context, method string) context.Context {
	caller := getCaller(ctx)
	caller.Method = method
	return context.WithValue(ctx, ContextKeyCaller, caller)
}

// GetCallerMethod ...
func GetCallerMethod(ctx context.Context) (method string, ok bool) {
	caller, ok := ctx.Value(ContextKeyCaller).(ContextCaller)
	if !ok {
		return
	}
	return caller.Method, true
}

type ValueContext struct {
	ctx context.Context
}

func (c ValueContext) Deadline() (time.Time, bool)       { return time.Time{}, false }
func (c ValueContext) Done() <-chan struct{}             { return nil }
func (c ValueContext) Err() error                        { return nil }
func (c ValueContext) Value(key interface{}) interface{} { return c.ctx.Value(key) }

// NewValueContext returns a context that is never canceled.
func NewValueContext(ctx context.Context) context.Context {
	return ValueContext{ctx: ctx}
}
