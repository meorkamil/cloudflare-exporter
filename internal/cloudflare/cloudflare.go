package cloudflare

import (
	"cloudflare-status/internal/api"
	"cloudflare-status/internal/models"
	"log"
)

type CfConfig struct {
	endpoint string
}

func NewCloudFlare(e string) *CfConfig {
	return &CfConfig{
		endpoint: e,
	}
}

func (c CfConfig) CfSummaries(ch chan<- float64) {
	summaryPayload, err := api.GetAPI(c.endpoint + "/status.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	var summary models.Summary
	api.UnmarshalJson(summaryPayload, &summary)

	switch summary.Status.Indicator {
	case "minor":
		ch <- 1
	case "major":
		ch <- 2
	default:
		ch <- 0
	}
}
func (c CfConfig) CfIncidents(ch chan<- models.Incidents) {
	incidentsPayload, err := api.GetAPI(c.endpoint + "/incidents.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	var incidents models.Incidents
	api.UnmarshalJson(incidentsPayload, &incidents)

	ch <- incidents
}

func (c CfConfig) CfComponents(ch chan<- models.Components) {
	componentsPayload, err := api.GetAPI(c.endpoint + "/components.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	var components models.Components
	api.UnmarshalJson(componentsPayload, &components)

	ch <- components
}
