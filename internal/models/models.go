package models

import (
	"time"
)

type Summary struct {
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

type Components struct {
	Page struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		URL       string    `json:"url"`
		TimeZone  string    `json:"time_zone"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"page"`
	Components []struct {
		ID                 string      `json:"id"`
		Name               string      `json:"name"`
		Status             string      `json:"status"`
		CreatedAt          time.Time   `json:"created_at"`
		UpdatedAt          time.Time   `json:"updated_at"`
		Position           int         `json:"position"`
		Description        interface{} `json:"description"`
		Showcase           bool        `json:"showcase"`
		StartDate          interface{} `json:"start_date"`
		GroupID            string      `json:"group_id"`
		PageID             string      `json:"page_id"`
		Group              bool        `json:"group"`
		OnlyShowIfDegraded bool        `json:"only_show_if_degraded"`
		Components         []string    `json:"components,omitempty"`
	} `json:"components"`
}

type Incidents struct {
	Incidents []struct {
		ID              string      `json:"id"`
		Name            string      `json:"name"`
		Status          string      `json:"status"`
		CreatedAt       time.Time   `json:"created_at"`
		UpdatedAt       time.Time   `json:"updated_at"`
		MonitoringAt    interface{} `json:"monitoring_at"`
		ResolvedAt      interface{} `json:"resolved_at"`
		Impact          string      `json:"impact"`
		Shortlink       string      `json:"shortlink"`
		StartedAt       time.Time   `json:"started_at"`
		PageID          string      `json:"page_id"`
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
			DeliverNotifications bool        `json:"deliver_notifications"`
			CustomTweet          interface{} `json:"custom_tweet"`
			TweetID              interface{} `json:"tweet_id"`
		} `json:"incident_updates"`
		Components []struct {
			ID                 string      `json:"id"`
			Name               string      `json:"name"`
			Status             string      `json:"status"`
			CreatedAt          time.Time   `json:"created_at"`
			UpdatedAt          time.Time   `json:"updated_at"`
			Position           int         `json:"position"`
			Description        interface{} `json:"description"`
			Showcase           bool        `json:"showcase"`
			StartDate          interface{} `json:"start_date"`
			GroupID            string      `json:"group_id"`
			PageID             string      `json:"page_id"`
			Group              bool        `json:"group"`
			OnlyShowIfDegraded bool        `json:"only_show_if_degraded"`
		} `json:"components"`
		ReminderIntervals interface{} `json:"reminder_intervals"`
	} `json:"incidents"`
}

type PromConf struct {
	Port int
}
