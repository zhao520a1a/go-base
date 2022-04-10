package xprometheus

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"

	"gitlab.pri.ibanyu.com/middleware/seaweed/xstat/xmetric"
)

// GaugeVecOpts is an alias of VectorOpts.
type GaugeVecOpts xmetric.VectorOpts

// gaugeVec gauge vec.
type promGaugeVec struct {
	gv  *prometheus.GaugeVec
	lvs xmetric.LabelValues
}

// NewGauge constructs and registers a Prometheus GaugeVec,
// and returns a usable Gauge object.
func NewGauge(cfg *GaugeVecOpts) xmetric.Gauge {
	if cfg == nil {
		return nil
	}
	vec := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
			Name:      cfg.Name,
			Help:      cfg.Help,
		}, cfg.LabelNames)
	prometheus.MustRegister(vec)
	return &promGaugeVec{
		gv: vec,
	}
}

// With add label k-v pairs
func (g *promGaugeVec) With(labelValues ...string) xmetric.Gauge {
	return &promGaugeVec{
		gv:  g.gv,
		lvs: g.lvs.With(labelValues...),
	}
}

// Add Inc increments the counter by 1. Use Add to increment it by arbitrary.
func (g *promGaugeVec) Add(delta float64) {
	if err := g.lvs.Check(); err != nil {
		fmt.Printf("gauge label value invalid:%s\n", err.Error())
		return
	}
	g.gv.With(makeLabels(g.lvs...)).Add(delta)
}

// Set set the given value to the collection.
func (g *promGaugeVec) Set(v float64) {
	if err := g.lvs.Check(); err != nil {
		fmt.Printf("gauge label value invalid:%s\n", err.Error())
		return
	}
	g.gv.With(makeLabels(g.lvs...)).Set(v)
}
