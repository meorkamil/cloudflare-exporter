package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type CfSummary struct {
	Page struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		URL       string    `json:"url"`
		TimeZone  string    `json:"time_zone"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"page"`
	Components []struct {
		ID                 string    `json:"id"`
		Name               string    `json:"name"`
		Status             string    `json:"status"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		Position           int       `json:"position"`
		Description        any       `json:"description"`
		Showcase           bool      `json:"showcase"`
		StartDate          any       `json:"start_date"`
		GroupID            string    `json:"group_id"`
		PageID             string    `json:"page_id"`
		Group              bool      `json:"group"`
		OnlyShowIfDegraded bool      `json:"only_show_if_degraded"`
		Components         []string  `json:"components,omitempty"`
	} `json:"components"`
	Incidents             []any `json:"incidents"`
	ScheduledMaintenances []struct {
		ID              string    `json:"id"`
		Name            string    `json:"name"`
		Status          string    `json:"status"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
		MonitoringAt    any       `json:"monitoring_at"`
		ResolvedAt      any       `json:"resolved_at"`
		Impact          string    `json:"impact"`
		Shortlink       string    `json:"shortlink"`
		StartedAt       time.Time `json:"started_at"`
		PageID          string    `json:"page_id"`
		IncidentUpdates []struct {
			ID                 string    `json:"id"`
			Status             string    `json:"status"`
			Body               string    `json:"body"`
			IncidentID         string    `json:"incident_id"`
			CreatedAt          time.Time `json:"created_at"`
			UpdatedAt          time.Time `json:"updated_at"`
			DisplayAt          time.Time `json:"display_at"`
			AffectedComponents []struct {
				Code      string `json:"code"`
				Name      string `json:"name"`
				OldStatus string `json:"old_status"`
				NewStatus string `json:"new_status"`
			} `json:"affected_components"`
			DeliverNotifications bool `json:"deliver_notifications"`
			CustomTweet          any  `json:"custom_tweet"`
			TweetID              any  `json:"tweet_id"`
		} `json:"incident_updates"`
		Components []struct {
			ID                 string    `json:"id"`
			Name               string    `json:"name"`
			Status             string    `json:"status"`
			CreatedAt          time.Time `json:"created_at"`
			UpdatedAt          time.Time `json:"updated_at"`
			Position           int       `json:"position"`
			Description        any       `json:"description"`
			Showcase           bool      `json:"showcase"`
			StartDate          any       `json:"start_date"`
			GroupID            string    `json:"group_id"`
			PageID             string    `json:"page_id"`
			Group              bool      `json:"group"`
			OnlyShowIfDegraded bool      `json:"only_show_if_degraded"`
		} `json:"components"`
		ScheduledFor   time.Time `json:"scheduled_for"`
		ScheduledUntil time.Time `json:"scheduled_until"`
	} `json:"scheduled_maintenances"`
	Status struct {
		Indicator   string `json:"indicator"`
		Description string `json:"description"`
	} `json:"status"`
}

type CfStatus struct {
	Page struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		URL       string    `json:"url"`
		TimeZone  string    `json:"time_zone"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"page"`
	Status struct {
		Indicator   string `json:"indicator"`
		Description string `json:"description"`
	} `json:"status"`
}

func main() {

	// TODO move to interface

	// CF Summary
	dataSummary := getCfSummary("https://www.cloudflarestatus.com/api/v2/summary.json")

	for _, component := range dataSummary.Components {

		fmt.Printf("%s |  %s\n", component.Name, component.Status)
	}

	// CF Status
	cfStatus := getCfStatus("https://www.cloudflarestatus.com/api/v2/status.json")
	fmt.Println(cfStatus.Status.Indicator)

}

func getCfSummary(url string) CfSummary {

	res, err := http.Get(url)

	if err != nil {

		log.Fatal(err)
	}

	defer res.Body.Close()

	var summary CfSummary

	body, err := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(body, &summary); err != nil {
		fmt.Println("Failed to unmarshal JSON")
	}

	return summary

}

func getCfStatus(url string) CfStatus {

	res, err := http.Get(url)

	if err != nil {

		log.Fatal(err)
	}

	defer res.Body.Close()

	var cfstatus CfStatus

	body, err := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(body, &cfstatus); err != nil {
		fmt.Println("Failed to unmarshal JSON")
	}

	return cfstatus

}
