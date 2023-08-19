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
