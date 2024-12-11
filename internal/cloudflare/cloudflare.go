package cloudflare

import (
	"cloudflare-status/internal/api"
	"cloudflare-status/internal/models"
	"fmt"

	"golang.org/x/exp/slog"
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
		slog.Error(fmt.Sprintf("error api summaries: %s", err))
		return
	}

	var summary models.Summary
	if err := api.UnmarshalJson(summaryPayload, &summary); err != nil {
		slog.Error(fmt.Sprintf("error json summaries: %s", err))
		return
	}

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
		slog.Error(fmt.Sprintf("error api incidents: %s", err))
		return
	}

	var incidents models.Incidents
	if err := api.UnmarshalJson(incidentsPayload, &incidents); err != nil {
		slog.Error(fmt.Sprintf("error json incidents: %s", err))
		return
	}

	ch <- incidents
}

func (c CfConfig) CfComponents(ch chan<- models.Components) {
	componentsPayload, err := api.GetAPI(c.endpoint + "/components.json")
	if err != nil {
		slog.Error(fmt.Sprintf("error api components: %s", err))
		return
	}

	var components models.Components
	if err := api.UnmarshalJson(componentsPayload, &components); err != nil {
		slog.Error(fmt.Sprintf("error json components: %s", err))
		return
	}

	ch <- components
}
