package cloudflare

import (
	"cloudflare-status/internal/api"
	"cloudflare-status/internal/models"
)

func CfSummaries() float64 {

	summaryPayload := api.GetAPI("https://www.cloudflarestatus.com/api/v2/status.json")

	var summary models.Summary

	api.UnmarshalJson(summaryPayload, &summary)

	switch summary.Status.Indicator {
	case "minor":
		return 1
	case "major":
		return 2
	default:
		return 0
	}

}

func CfComponents() *models.Components {

	componentsPayload := api.GetAPI("https://www.cloudflarestatus.com/api/v2/components.json")

	var components models.Components

	api.UnmarshalJson(componentsPayload, &components)

	return &components

}

func CfIncidents() *models.Incidents {

	incidentsPayload := api.GetAPI("https://www.cloudflarestatus.com/api/v2/incidents.json")

	var incidents models.Incidents

	api.UnmarshalJson(incidentsPayload, &incidents)

	return &incidents

}
