package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/thomasf/lg"
)

func init() {
	prometheus.MustRegister(prometheus.NewCounterFunc(prometheus.CounterOpts{
		Name: "lg_info_lines",
		Help: "Number of entries lg error lines logged",
	}, func() float64 {
		return float64(lg.Stats.Info.Lines())
	}))

	prometheus.MustRegister(prometheus.NewCounterFunc(prometheus.CounterOpts{
		Name: "lg_warning_lines",
		Help: "Number of entries lg error lines logged",
	}, func() float64 {
		return float64(lg.Stats.Warning.Lines())
	}))

	prometheus.MustRegister(prometheus.NewCounterFunc(prometheus.CounterOpts{
		Name: "lg_error_lines",
		Help: "Number of entries lg error lines logged",
	}, func() float64 {
		return float64(lg.Stats.Error.Lines())
	}))
}
