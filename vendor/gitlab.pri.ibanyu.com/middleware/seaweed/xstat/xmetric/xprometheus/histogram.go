package xprometheus

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"

	"gitlab.pri.ibanyu.com/middleware/seaweed/xstat/xmetric"
)

// HistogramVecOpts is histogram vector opts.
type HistogramVecOpts struct {
	Namespace  string
	Subsystem  string
	Name       string
	Help       string
	LabelNames []string
	Buckets    []float64
}

// Histogram prom histogram collection.
type promHistogramVec struct {
	hv  *prometheus.HistogramVec
	lvs xmetric.LabelValues
}

// NewHistogram constructs and registers a Prometheus HistogramVec,
// and returns a usable Histogram object.
func NewHistogram(cfg *HistogramVecOpts) xmetric.Histogram {
	if cfg == nil {
		return nil
	}
	vec := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: cfg.Namespace,
			Subsystem: cfg.Subsystem,
			Name:      cfg.Name,
			Help:      cfg.Help,
			Buckets:   cfg.Buckets,
		}, cfg.LabelNames)
	prometheus.MustRegister(vec)
	return &promHistogramVec{
		hv: vec,
	}
}

// With append k-v pairs to histogram lvs
func (h *promHistogramVec) With(labelValues ...string) xmetric.Histogram {
	return &promHistogramVec{
		hv:  h.hv,
		lvs: h.lvs.With(labelValues...),
	}
}

// Timing adds a single observation to the histogram.
func (h *promHistogramVec) Observe(v float64) {
	if err := h.lvs.Check(); err != nil {
		fmt.Printf("histogram label value invalid:%s\n", err.Error())
		return
	}
	h.hv.With(makeLabels(h.lvs...)).Observe(v)
}
