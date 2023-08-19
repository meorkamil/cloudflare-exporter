package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type CfComponents struct {
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

type CfUnresolve struct {
	Page struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		URL       string    `json:"url"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"page"`
	Incidents []struct {
		CreatedAt       string `json:"created_at"`
		ID              string `json:"id"`
		Impact          string `json:"impact"`
		IncidentUpdates []struct {
			Body       string `json:"body"`
			CreatedAt  string `json:"created_at"`
			DisplayAt  string `json:"display_at"`
			ID         string `json:"id"`
			IncidentID string `json:"incident_id"`
			Status     string `json:"status"`
			UpdatedAt  string `json:"updated_at"`
		} `json:"incident_updates"`
		MonitoringAt any    `json:"monitoring_at"`
		Name         string `json:"name"`
		PageID       string `json:"page_id"`
		ResolvedAt   any    `json:"resolved_at"`
		Shortlink    string `json:"shortlink"`
		Status       string `json:"status"`
		UpdatedAt    string `json:"updated_at"`
	} `json:"incidents"`
}

type CfUpcomingSchedules struct {
	Page struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		URL       string    `json:"url"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"page"`
	ScheduledMaintenances []struct {
		CreatedAt       string `json:"created_at"`
		ID              string `json:"id"`
		Impact          string `json:"impact"`
		IncidentUpdates []struct {
			Body       string `json:"body"`
			CreatedAt  string `json:"created_at"`
			DisplayAt  string `json:"display_at"`
			ID         string `json:"id"`
			IncidentID string `json:"incident_id"`
			Status     string `json:"status"`
			UpdatedAt  string `json:"updated_at"`
		} `json:"incident_updates"`
		MonitoringAt   any    `json:"monitoring_at"`
		Name           string `json:"name"`
		PageID         string `json:"page_id"`
		ResolvedAt     any    `json:"resolved_at"`
		ScheduledFor   string `json:"scheduled_for"`
		ScheduledUntil string `json:"scheduled_until"`
		Shortlink      string `json:"shortlink"`
		Status         string `json:"status"`
		UpdatedAt      string `json:"updated_at"`
	} `json:"scheduled_maintenances"`
}

func main() {

	// TODO move to interface

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

	// CF Upcoming Schedules
	cfUpcomingScedules := getCfUpcomingSchedules("https://www.cloudflarestatus.com/api/v2/scheduled-maintenances/upcoming.json")
	fmt.Println(cfUpcomingSchedules)

}

func getCfComponents(url string) CfComponents {

	res, err := http.Get(url)

	if err != nil {

		log.Fatal(err)
	}

	defer res.Body.Close()

	var component CfComponents

	body, err := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(body, &component); err != nil {
		fmt.Println("Failed to unmarshal JSON")
	}

	return component

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

func getCfUnresolve(url string) CfUnresolve {

	res, err := http.Get(url)

	if err != nil {

		log.Fatal(err)
	}

	defer res.Body.Close()

	var cfunresolve CfUnresolve

	body, err := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(body, &cfunresolve); err != nil {
		fmt.Println("Failed to unmarshal JSON")
	}

	return cfunresolve

}

func getCfUpcomingSchedule(url string) CfUpcomingSchedules {

	res, err := http.Get(url)

	if err != nil {

		fmt.Println("failed to scrape endpoint")

	}

	defer res.Body.Close()

	var cfupcomingschedules CfUpComingSchedules

	body, err := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(body, &cfupcomingschedules); err != nil {
		fmt.Println("Failed to unmarshal JSON")
	}

	return cfupcomingschedules
}
