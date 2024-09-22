package main

import (
	"cloudflare-status/cloudflare"
	"cloudflare-status/metrics"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	reg := prometheus.NewRegistry()
	m := metrics.NewMetrics(reg)
	incidents := cloudflare.CfIncidents()
	components := cloudflare.CfComponents()

	// Register metrics
	m.CfSummaryMetric.Set(cloudflare.CfSummaries())
	log.Println("Creating cloudflare summary metrics")

	for _, s := range incidents.Incidents {

		log.Println("Creating cloudflare incident:", s.Name)

		if s.Status != "resolved" {
			m.CfIncidentMetric.With(prometheus.Labels{"name": s.Name}).Set(1)
		} else {
			m.CfIncidentMetric.With(prometheus.Labels{"name": s.Name}).Set(0)
		}
	}

	for _, s := range components.Components {

		log.Println("Creating cloudflare component", s.Name)

		if s.Status != "operational" {
			m.CfComponentMetric.With(prometheus.Labels{"name": s.Name}).Set(1)
		} else {
			m.CfComponentMetric.With(prometheus.Labels{"name": s.Name}).Set(0)
		}
	}

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	log.Fatal(http.ListenAndServe(":5001", nil))

}
