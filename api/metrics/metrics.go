package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	URLsAdded = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "urls_added",
			Help: "Number of URLs added",
		},
	)
	URLsRemoved = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "urls_removed",
			Help: "Number of URLs removed",
		},
	)
	ChecksTotal = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "checks_total",
			Help: "Total number of checks performed",
		},
	)
	ChecksFailed = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "checks_failed",
			Help: "Total number of failed checks",
		},
	)
	ChecksSucceeded = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "checks_succeeded",
			Help: "Total number of successful checks",
		},
	)
	ServiceStatus = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "service_status",
			Help: "Status of the service",
		},
		[]string{"service"},
	)
)

func RecordSuccess(service string) {
	ChecksTotal.Inc()
	ChecksSucceeded.Inc()
	ServiceStatus.WithLabelValues(service).Set(1)
}
func RecordFailure(service string) {
	ChecksTotal.Inc()
	ChecksFailed.Inc()
	ServiceStatus.WithLabelValues(service).Set(0)
}
func RecordUrlAdded() {
	URLsAdded.Inc()
}
func RecordUrlRemoved() {
	URLsRemoved.Inc()
}
