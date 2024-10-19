package main

import (
	"cloudflare-status/internal/metrics"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	metrics.RecordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":5001", nil)
}
