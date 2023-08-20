package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	//fetchMetrics()
	// Create a non-global registry.
	reg := prometheus.NewRegistry()

	// Create new metrics and register them using the custom registry.
	m := NewMetrics(reg)
	// Set values for the new created metrics.
	m.cfStatusMetric.Set(65)

	// Expose metrics and custom registry via an HTTP server
	// using the HandleFor function. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	log.Fatal(http.ListenAndServe(":50001", nil))

}

func fetchMetrics() {

	// CF Summary
	dataComponents := getCfComponents("https://www.cloudflarestatus.com/api/v2/summary.json")

	for _, component := range dataComponents.Components {

		if component.Status != "operational" {
			fmt.Printf("cf_components_status{componet='%s'}0\n", component.Name)
		} else {
			fmt.Printf("cf_components_status{componet='%s'}1\n", component.Name)
		}
	}

	// CF Status
	cfStatus := getCfStatus("https://www.cloudflarestatus.com/api/v2/status.json")

	if cfStatus.Status.Indicator == "minor" {
		fmt.Printf("cf_status{='https://www.cloudflarestatus.com/api/v2/status.json'}2\n")
	} else if cfStatus.Status.Indicator == "major" {
		fmt.Printf("cf_status{='https://www.cloudflarestatus.com/api/v2/status.json'}1\n")
	} else if cfStatus.Status.Indicator == "crtical" {
		fmt.Printf("cf_status{='https://www.cloudflarestatus.com/api/v2/status.json'}0\n")
	} else {
		fmt.Printf("cf_status{='https://www.cloudflarestatus.com/api/v2/status.json'}3\n")
	}

	// CF Unresolve
	cfUnresolve := getCfUnresolve("https://www.cloudflarestatus.com/api/v2/incidents/unresolved.json")
	fmt.Println(cfUnresolve)

	// CF Schedules
	cfSchedules := getCfSchedules("https://www.cloudflarestatus.com/api/v2/scheduled-maintenances/upcoming.json")

	for _, schedules := range cfSchedules.ScheduledMaintenances {

		for _, incidentupdates := range schedules.IncidentUpdates {
			fmt.Printf("cf_schedules{descripton='%s'}1 \n", incidentupdates.Body)
		}
	}

}
