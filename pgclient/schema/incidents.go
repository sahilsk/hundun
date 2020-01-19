package schema

import (
	"encoding/json"
	"log"
	"time"
)

type IncidentResponse struct {
	Incident Incident `json:"incident"`
}

type IncidentsResponse struct {
	Incidents []Incident `json:"incidents"`
	Pagination
}

type Incident struct {
	Entity
	Incident_number       uint                 `json:"incident_number"`
	Created_at            time.Time            `json:"created_at"`
	Status                string               `json:"status"`
	Title                 string               `json:"title"`
	ResolveReason         string               `json:"resolve_reason"`
	AlertCounts           AlertCounts          `json:"alert_counts"`
	IsMergeable           bool                 `json:"is_mergeable"`
	PendingActions        []PendingAction      `json:"pending_actions"`
	IncidentKey           string               `json:"incident_key"`
	Service               Service              `json:"service"`
	Priority              Priority             `json:"priority"`
	Assigned_via          string               `json:"assigned_via"`
	Assignments           []Assignment         `json:"assignments"`
	Acknowledgements      []Acknowledgement    `json:"acknowledgemnts"`
	Last_status_change_at time.Time            `json:"last_status_change_at"`
	LastStatusChangeBy    LastStatusChangeBy   `json:"last_status_change_by"`
	FirstTriggerLogEntry  FirstTriggerLogEntry `json:"first_trigger_log_entry"`
	EscalationPolicy      EscalationPolicy     `json:"escalation_policy"`
	Teams                 []Team               `json:"teams"`
	Urgency               string               `json:"urgency"`
}

type AlertCounts struct {
	All       uint `json:"all"`
	Triggered uint `json:"triggered"`
	Resolved  uint `json:"resolved"`
}

type PendingAction struct {
	Type string `json:"type"`
	At   string `json:"at"`
}

type Service struct {
	Entity
}

type Assignment struct {
	At       string   `json:"at"`
	Assignee Assignee `json:"assignee"`
}

type Assignee struct {
	Entity
}

type Acknowledgement struct {
	At           string `json:"at`
	Acknowledger Acknowledger
}

type Acknowledger struct {
	Entity
}

type LastStatusChangeBy struct {
	Entity
}

type FirstTriggerLogEntry struct {
	Entity
}

type EscalationPolicy struct {
	Entity
}

type Team struct {
	Entity
}

func (ir *IncidentsResponse) ToPrettyString() string {
	b, err := json.MarshalIndent(*ir, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func (ir *IncidentsResponse) ToString() string {
	b, err := json.Marshal(*ir)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func (ir *IncidentResponse) ToPrettyString() string {
	b, err := json.MarshalIndent(*ir, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func (ir *IncidentResponse) ToString() string {
	b, err := json.Marshal(*ir)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
