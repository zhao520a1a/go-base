package thriftutil

import (
	"time"
)

func NewDefaultControl() *Control {
	return &Control{
		Route:  &Route{""},
		Ct:     time.Now().Unix(),
		Et:     0,
		Caller: &Endpoint{"","",""},
	}
}

func CreateContextByUid(uid int64) *Context {
	return &Context{
		Head: &Head{
			Uid: uid,
		},
		Control: NewDefaultControl(),
	}
}

func CreateContext(uid int64, source, dt int32, unionid, ip, region string) *Context {
	return &Context{
		Head: &Head{
			Uid:     uid,
			Source:  source,
			Ip:      ip,
			Region:  region,
			Dt:      dt,
			Unionid: unionid,
		},
		Control: NewDefaultControl(),
	}
}

func (c *Context) ensureControl() {
	if c.Control == nil {
		c.Control = NewDefaultControl()
		return
	}

	if c.Control.Route == nil {
		c.Control.Route = &Route{}
	}

	if c.Control.Caller == nil {
		c.Control.Caller = &Endpoint{}
	}
}

// deprecated
func (c *Context) SetGroup(group string) {
	c.ensureControl()
	c.Control.Route.Group = group
}

// deprecated
func (c *Context) SetExpireTime(et int64) {
	c.ensureControl()
	c.Control.Et = et
}

func (c *Control) GetControlRouteGroup() (string, bool) {
	if c.Route == nil {
		return "", false
	}
	return c.Route.Group, true
}

func (c *Control) ensureRoute() {
	if c.Route == nil {
		c.Route = &Route{}
	}
}

func (c *Control) SetControlRouteGroup(group string) error {
	c.ensureRoute()
	c.Route.Group = group
	return nil
}

func (c *Control) ensureCaller() {
	if c.Caller == nil {
		c.Caller = &Endpoint{}
	}
}

func (c *Control) GetControlCallerServerName() (string, bool) {
	if c.Caller == nil {
		return "", false
	}
	return c.Caller.Sname, true
}

func (c *Control) SetControlCallerServerName(sname string) error {
	c.ensureCaller()
	c.Caller.Sname = sname
	return nil
}

func (c *Control) GetControlCallerServerId() (string, bool) {
	if c.Caller == nil {
		return "", false
	}
	return c.Caller.Sid, true
}

func (c *Control) SetControlCallerServerId(sid string) error {
	c.ensureCaller()
	c.Caller.Sid = sid
	return nil
}

func (c *Control) GetControlCallerMethod() (string, bool) {
	if c.Caller == nil {
		return "", false
	}
	return c.Caller.Method, true
}

func (c *Control) SetControlCallerMethod(method string) error {
	c.ensureCaller()
	c.Caller.Method = method
	return nil
}

func (h *Head) ToKV() map[string]interface{} {
	return map[string]interface{}{
		"uid":     h.Uid,
		"source":  h.Source,
		"ip":      h.Ip,
		"region":  h.Region,
		"dt":      h.Dt,
		"unionid": h.Unionid,
	}
}
