package metrics

import (
	"cloudflare-status/internal/cloudflare"
	"cloudflare-status/internal/models"
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

	sumchan := make(chan float64)
	indchan := make(chan models.Incidents)
	comchan := make(chan models.Components)
	cf := cloudflare.NewCloudFlare("https://www.cloudflarestatus.com/api/v2")

	go func() {
		for {
			go cf.CfSummaries(sumchan)
			go cf.CfIncidents(indchan)
			go cf.CfComponents(comchan)

			select {
			case v := <-sumchan:
				log.Println("cloudflare summary:", v)
				CfSumMetric.Set(v)
			case v := <-indchan:
				for _, s := range v.Incidents {
					log.Println("cloudflare incident:", s.Name)

					if s.Status != "resolved" {
						CfIncMetric.With(prometheus.Labels{"name": s.Name}).Set(1)
					} else {
						CfIncMetric.With(prometheus.Labels{"name": s.Name}).Set(0)
					}
				}
			case v := <-comchan:
				for _, s := range v.Components {

					log.Println("cloudflare component:", s.Name)

					if s.Status != "operational" {
						CfComMetric.With(prometheus.Labels{"name": s.Name}).Set(1)
					} else {
						CfComMetric.With(prometheus.Labels{"name": s.Name}).Set(0)
					}
				}

			}

			time.Sleep(5 * time.Second)
		}
	}()
}
