package metrics

import (
	"cloudflare-status/internal/cloudflare"
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	OpsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
	CfSumMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "cloudflare_exporter_summary",
		Help: "Current cloudflare summary",
	})
	CfIncMetric = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "cloudflare_exporter_incident",
		Help: "Current cloudflare incident",
	},
		[]string{"name"},
	)
	CfComMetric = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "cloudflare_exporter_component",
		Help: "Current cloudflare component status",
	},
		[]string{"name"},
	)
)

func RecordMetrics() {
	go func() {
		for {
			incidents := cloudflare.CfIncidents()
			components := cloudflare.CfComponents()
			CfSumMetric.Set(cloudflare.CfSummaries())

			for _, s := range components.Components {

				log.Println("cloudflare component", s.Name)

				if s.Status != "operational" {
					CfComMetric.With(prometheus.Labels{"name": s.Name}).Set(1)
				} else {
					CfComMetric.With(prometheus.Labels{"name": s.Name}).Set(0)
				}
			}

			for _, s := range incidents.Incidents {

				log.Println("cloudflare incident:", s.Name)

				if s.Status != "resolved" {
					CfIncMetric.With(prometheus.Labels{"name": s.Name}).Set(1)
				} else {
					CfIncMetric.With(prometheus.Labels{"name": s.Name}).Set(0)
				}
			}

			time.Sleep(5 * time.Second)
		}
	}()
}
