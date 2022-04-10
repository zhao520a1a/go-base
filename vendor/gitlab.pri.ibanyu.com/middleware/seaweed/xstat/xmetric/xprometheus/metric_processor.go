package xprometheus

import (
	"net/http"
	"net/http/pprof"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	defaultMetricLocation = "/metrics"
)

// MetricProcessor 为了兼容roc注册metrics，后续框架整合时会拆分使各个库更加独立
type MetricProcessor struct {
}

// NewMetricProcessor constructor of MetricProcessor
func NewMetricProcessor() *MetricProcessor {
	return new(MetricProcessor)
}

// Init do nothing
func (mp *MetricProcessor) Init() error {
	return nil
}

// Driver return addr and http driver
func (mp *MetricProcessor) Driver() (string, interface{}) {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.GET(defaultMetricLocation, monitor())

	// TODO pprof所依赖handler，后期统一迁移到后门程序，以后程序只有一个服务端口、一个管理端口
	engine.GET("/debug/pprof/", pprofHandler(pprof.Index))
	engine.POST("/debug/pprof/", pprofHandler(pprof.Index))
	engine.GET("/debug/pprof/allocs", pprofHandler(pprof.Index))
	engine.POST("/debug/pprof/allocs", pprofHandler(pprof.Index))
	engine.GET("/debug/pprof/block", pprofHandler(pprof.Index))
	engine.POST("/debug/pprof/block", pprofHandler(pprof.Index))
	engine.GET("/debug/pprof/goroutine", pprofHandler(pprof.Index))
	engine.POST("/debug/pprof/goroutine", pprofHandler(pprof.Index))
	engine.GET("/debug/pprof/heap", pprofHandler(pprof.Index))
	engine.POST("/debug/pprof/heap", pprofHandler(pprof.Index))
	engine.GET("/debug/pprof/mutex", pprofHandler(pprof.Index))
	engine.POST("/debug/pprof/mutex", pprofHandler(pprof.Index))
	engine.GET("/debug/pprof/threadcreate", pprofHandler(pprof.Index))
	engine.POST("/debug/pprof/threadcreate", pprofHandler(pprof.Index))
	engine.GET("/debug/pprof/cmdline", pprofHandler(pprof.Cmdline))
	engine.POST("/debug/pprof/cmdline", pprofHandler(pprof.Cmdline))
	engine.GET("/debug/pprof/profile", pprofHandler(pprof.Profile))
	engine.POST("/debug/pprof/profile", pprofHandler(pprof.Profile))
	engine.GET("/debug/pprof/trace", pprofHandler(pprof.Trace))
	engine.POST("/debug/pprof/trace", pprofHandler(pprof.Trace))
	engine.GET("/debug/pprof/symbol", pprofHandler(pprof.Symbol))
	engine.POST("/debug/pprof/symbol", pprofHandler(pprof.Symbol))
	return "", engine
}

func monitor() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := promhttp.Handler()
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func pprofHandler(h http.HandlerFunc) gin.HandlerFunc {
	handler := http.HandlerFunc(h)
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
