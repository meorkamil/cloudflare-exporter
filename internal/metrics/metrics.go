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

type Metrics struct {
	Timeout   int
	Interval  int
	API       string
	ChanOpen  chan bool
	ChanClose chan bool
}

func NewMetrics(t int, i int, a string) (*Metrics, error) {
	return &Metrics{
		Timeout:  t,
		Interval: i,
		API:      a,
	}, nil
}

func (m *Metrics) RecordMetrics() {
	cf := cloudflare.NewCloudFlare(m.API)
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
				switch {
				case s.Status == "investigating":
					CfIncMetric.With(prometheus.Labels{"name": s.Name}).Set(2)
				case s.Status == "identified":
					CfIncMetric.With(prometheus.Labels{"name": s.Name}).Set(2)
				case s.Status == "monitoring":
					CfIncMetric.With(prometheus.Labels{"name": s.Name}).Set(1)
				case s.Status == "resolved":
					CfIncMetric.With(prometheus.Labels{"name": s.Name}).Set(0)
				}
			}
			close(indchan)
		case v := <-comchan:
			for _, s := range v.Components {
				switch {
				case s.Status == "degraded_performance":
					CfComMetric.With(prometheus.Labels{"name": s.Name}).Set(1)
				case s.Status == "partial_outage":
					CfComMetric.With(prometheus.Labels{"name": s.Name}).Set(2)
				case s.Status == "major_outage":
					CfComMetric.With(prometheus.Labels{"name": s.Name}).Set(3)
				case s.Status == "operational":
					CfComMetric.With(prometheus.Labels{"name": s.Name}).Set(0)
				}
			}
			close(comchan)
		}

		time.Sleep(time.Duration(m.Interval) * time.Second)
	}
}
