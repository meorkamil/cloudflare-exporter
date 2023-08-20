package main

import "github.com/prometheus/client_golang/prometheus"

type metrics struct {
	cfStatusMetric     prometheus.Gauge
	cfComponentsMetric prometheus.Gauge
	cfScheduleMetric   prometheus.Gauge
	cfUnresolveMetric  prometheus.Gauge
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		cfStatusMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "cloudflare_status",
			Help: "Cloudflare Status",
		}),
		cfComponentsMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "cloudflare_components",
			Help: "Cloudflare Components",
		}),
		cfScheduleMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "cloudflare_schedule",
			Help: "Cloudflare Schedule Maintenance",
		}),
		cfUnresolveMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "cloudflare_unresolve",
			Help: "Cloudflare Unresolve Incident",
		}),
	}

	reg.MustRegister(m.cfStatusMetric)
	reg.MustRegister(m.cfScheduleMetric)
	reg.MustRegister(m.cfComponentsMetric)
	reg.MustRegister(m.cfUnresolveMetric)
	return m
}
