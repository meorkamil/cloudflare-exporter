package main

import (
	"cloudflare-status/api"
	"cloudflare-status/models"
	"fmt"
)

func main() {

	summaryPayload := api.GetAPI("https://www.cloudflarestatus.com/api/v2/status.json")
	componentsPayload := api.GetAPI("https://www.cloudflarestatus.com/api/v2/components.json")
	incidentsPayload := api.GetAPI("https://www.cloudflarestatus.com/api/v2/incidents.json")

	var summary models.Summary
	var components models.Components
	var incidents models.Incidents

	api.UnmarshalJson(summaryPayload, &summary)
	api.UnmarshalJson(componentsPayload, &components)
	api.UnmarshalJson(incidentsPayload, &incidents)

	switch summary.Status.Indicator {
	case "minor":
		fmt.Println("cloudflare_exporter_summary{status='minor'}1")
	case "major":
		fmt.Println("cloudflare_exporter_summary{status='major'}2")
	default:
		fmt.Println("cloudflare_exporter_summary{status='okay'}0")
	}

	for _, s := range components.Components {
		if s.Status != "operational" {
			fmt.Printf("cloudflare_exporter_component{name='%s'}1\n", s.Name)
		} else {
			fmt.Printf("cloudflare_exporter_component{name='%s'}0\n", s.Name)
		}
	}

	for _, s := range incidents.Incidents {
		if s.Status != "resolved" {
			fmt.Printf("cloudflare_exporter_incident{name='%s', status='%s'}1\n", s.Name, s.Status)
		}
	}

}
