package manager

import (
	"sync"
	"time"

	"gitlab.pri.ibanyu.com/middleware/seaweed/internal/dbrouter"
)

const (
	// MysqlConfNamespace mysql apollo conf namespace
	MysqlConfNamespace = "mysql"
	// MaxIdleConnsKey ...
	MaxIdleConnsKey = "maxIdleConns"
	// MaxOpenConnsKey ...
	MaxOpenConnsKey = "maxOpenConns"
	// MaxLifeTimeSecKey ...
	MaxLifeTimeSecKey = "maxLifeTimeSec"
	// TimeoutMsecKey ...
	TimeoutMsecKey = "timeoutMsec"
	// ReadTimeoutMsecKey ...
	ReadTimeoutMsecKey = "readTimeoutMsec"
	// WriteTimeoutMsecKey ...
	WriteTimeoutMsecKey = "writeTimeoutMsec"
	// KeySep
	KeySep = "."

	defaultMaxIdleConns = 64
	defaultMaxOpenConns = 128
	defaultReadTimeout  = time.Second * 10
	defaultWriteTimeout = time.Second * 10
	defaultMaxLifeTime  = time.Hour * 6
	defaultTimeout      = time.Second * 3
)

// MysqlConf ...
type MysqlConf struct {
	MaxIdleConns     int `properties:"maxIdleConns"`
	MaxOpenConns     int `properties:"maxOpenConns"`
	MaxLifeTimeSec   int `properties:"maxLifeTimeSec"`
	TimeoutMsec      int `properties:"timeoutMsec"`
	ReadTimeoutMsec  int `properties:"readTimeoutMsec"`
	WriteTimeoutMsec int `properties:"writeTimeoutMsec"`
}

type cfg struct {
	Confs map[string]MysqlConf `properties:"confs"`
}

// DynamicConf ...
type DynamicConf struct {
	Timeout        time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxLifeTimeSec time.Duration
	MaxIdleConns   int
	MaxOpenConns   int
}

// DynamicConfiger ...
type DynamicConfiger struct {
	mysqlConf     *cfg
	mu            sync.RWMutex
	closeInsChan  chan dbrouter.ChangeIns
	reloadInsChan chan dbrouter.ChangeIns
}

// loadDynamicConf ...
func (c *DynamicConfiger) loadDynamicConf(insName string) *DynamicConf {
	config := &DynamicConf{
		Timeout:        defaultTimeout,
		ReadTimeout:    defaultReadTimeout,
		WriteTimeout:   defaultWriteTimeout,
		MaxLifeTimeSec: defaultMaxLifeTime,
		MaxIdleConns:   defaultMaxIdleConns,
		MaxOpenConns:   defaultMaxOpenConns,
	}
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.mysqlConf == nil {
		return config
	}
	if v, ok := c.mysqlConf.Confs[insName]; ok {
		if v.MaxIdleConns != 0 {
			config.MaxIdleConns = v.MaxIdleConns
		}
		if v.MaxOpenConns != 0 {
			config.MaxOpenConns = v.MaxOpenConns
		}
		if v.MaxLifeTimeSec != 0 {
			config.MaxLifeTimeSec = time.Duration(v.MaxLifeTimeSec) * time.Second
		}
		if v.TimeoutMsec != 0 {
			config.Timeout = time.Duration(v.TimeoutMsec) * time.Millisecond
		}
		if v.ReadTimeoutMsec != 0 {
			config.ReadTimeout = time.Duration(v.ReadTimeoutMsec) * time.Millisecond
		}
		if v.WriteTimeoutMsec != 0 {
			config.WriteTimeout = time.Duration(v.WriteTimeoutMsec) * time.Millisecond
		}
	}
	return config
}

// reloadConf ...
func (c *DynamicConfiger) reloadConf(cfg *cfg) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.mysqlConf = cfg
}
