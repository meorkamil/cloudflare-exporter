package metrics

import (
	"cloudflare-status/internal/cloudflare"
	"cloudflare-status/internal/models"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
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
	cf := cloudflare.NewCloudFlare("https://www.cloudflarestatus.com/api/v2")
	for {
		sumchan := make(chan float64, 1)
		indchan := make(chan models.Incidents, 1)
		comchan := make(chan models.Components, 1)

		go cf.CfSummaries(sumchan)
		go cf.CfIncidents(indchan)
		go cf.CfComponents(comchan)

		select {
		case v := <-sumchan:
			CfSumMetric.Set(v)
			close(sumchan)
		case v := <-indchan:
			for _, s := range v.Incidents {
				if s.Status != "resolved" {
					CfIncMetric.With(prometheus.Labels{"name": s.Name}).Set(1)
				} else {
					CfIncMetric.With(prometheus.Labels{"name": s.Name}).Set(0)
				}
			}
			close(indchan)
		case v := <-comchan:
			for _, s := range v.Components {
				if s.Status != "operational" {
					CfComMetric.With(prometheus.Labels{"name": s.Name}).Set(1)
				} else {
					CfComMetric.With(prometheus.Labels{"name": s.Name}).Set(0)
				}
			}
			close(comchan)
		}

		time.Sleep(5 * time.Second)
	}
}
