package metrics

import "github.com/prometheus/client_golang/prometheus"

type metrics struct {
	CfSummaryMetric   prometheus.Gauge
	CfIncidentMetric  *prometheus.GaugeVec
	CfComponentMetric *prometheus.GaugeVec
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		CfSummaryMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "cloudflare_exporter_summary",
			Help: "Current Cloudflare Summary",
		}),
		CfIncidentMetric: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "cloudflare_exporter_incident",
				Help: "Current Cloudflare incident",
			},
			[]string{"name"},
		),
		CfComponentMetric: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "cloudflare_exporter_component",
				Help: "Current Cloudflare Component",
			},
			[]string{"name"},
		),
	}
	reg.MustRegister(m.CfSummaryMetric)
	reg.MustRegister(m.CfIncidentMetric)
	reg.MustRegister(m.CfComponentMetric)
	return m
}
