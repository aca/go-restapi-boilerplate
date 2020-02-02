package api

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	opsProcessing = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "api_processing_ops_total",
		Help: "The number of events processing",
	})
)

// mwMetrics is simple middleware to count ongoing requests.
func mwMetrics(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		defer opsProcessing.Dec()
		opsProcessing.Inc()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(f)
}
