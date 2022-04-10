package manager

import (
	"database/sql"

	xprom "gitlab.pri.ibanyu.com/middleware/seaweed/xstat/xmetric/xprometheus"
)

const namespace = "xsql"

var (
	msBuckets     = []float64{1, 3, 5, 10, 25, 50, 100, 200, 300, 500, 1000, 3000, 5000, 10000, 15000}
	_metricReqDur = xprom.NewHistogram(&xprom.HistogramVecOpts{
		Namespace:  namespace,
		Subsystem:  "requests",
		Name:       "duration_ms",
		Help:       "mysql client requests duration(ms).",
		Buckets:    msBuckets,
		LabelNames: []string{"cluster", "table", "command"},
	})

	_metricReqErrTotal = xprom.NewCounter(&xprom.CounterVecOpts{
		Namespace:  namespace,
		Subsystem:  "requests",
		Name:       "err_total",
		Help:       "mysql client err requests total.",
		LabelNames: []string{"cluster", "table", "command"},
	})

	_metricConnTotal = xprom.NewGauge(&xprom.GaugeVecOpts{
		Namespace:  namespace,
		Subsystem:  "connections",
		Name:       "total",
		Help:       "mysql client connections total count.",
		LabelNames: []string{"dbname"},
	})

	_metricConnInUse = xprom.NewGauge(&xprom.GaugeVecOpts{
		Namespace:  namespace,
		Subsystem:  "connections",
		Name:       "in_use",
		Help:       "mysql client connections in use.",
		LabelNames: []string{"dbname"},
	})

	_metricConnIdle = xprom.NewGauge(&xprom.GaugeVecOpts{
		Namespace:  namespace,
		Subsystem:  "connections",
		Name:       "idle",
		Help:       "mysql client connections idle.",
		LabelNames: []string{"dbname"},
	})
)

func statMetricReqErrTotal(cluster, table, command string, err error) {
	if err != nil && err != sql.ErrNoRows {
		_metricReqErrTotal.With("cluster", cluster, "table", table, "command", command).Inc()
	}
}
