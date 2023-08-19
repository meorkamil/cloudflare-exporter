package main

import (
	"fmt"
)

func main() {

	fetchMetrics()
}

func fetchMetrics() {

	// CF Summary
	dataComponents := getCfComponents("https://www.cloudflarestatus.com/api/v2/summary.json")

	for _, component := range dataComponents.Components {

		fmt.Printf("%s |  %s\n", component.Name, component.Status)
	}

	// CF Status
	cfStatus := getCfStatus("https://www.cloudflarestatus.com/api/v2/status.json")
	fmt.Println(cfStatus.Status.Indicator)

	// CF Unresolve
	cfUnresolve := getCfUnresolve("https://www.cloudflarestatus.com/api/v2/incidents/unresolved.json")
	fmt.Println(cfUnresolve)

	// CF Schedules
	cfSchedules := getCfSchedules("https://www.cloudflarestatus.com/api/v2/scheduled-maintenances/upcoming.json")

	for _, schedules := range cfSchedules.ScheduledMaintenances {

		for _, incidentupdates := range schedules.IncidentUpdates {
			fmt.Printf("%s | %s \n", incidentupdates.Body, incidentupdates.Status)
		}
	}

}
