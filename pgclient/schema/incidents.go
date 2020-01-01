package schema

type IncidentResponse struct {
	Incident Incident
}

type IncidentsResponse struct {
	Incidents []Incident
	Limit     uint
	Offset    uint
	Total     uint
	More      bool
}

type Incident struct {
	Id                    string               `json:id`
	Type                  string               `json:type`
	Summary               string               `json:summary`
	Self                  string               `json:self`
	Html_url              string               `json:html_url`
	Incident_number       string               `json:incident_number`
	Created_at            string               `json:created_at`
	Status                string               `json:status`
	Title                 string               `json:title`
	ResolveReason         string               `json:resolve_reason`
	AlertCounts           AlertCounts          `json:alert_counts`
	PendingActions        []PendingAction      `json:pending_actions`
	Incident_key          string               `json:incident_key`
	Service               Service              `json:service`
	Priority              string               `json:priority`
	Assigned_via          string               `json:assigned_via`
	Assignments           []Assignment         `json:assignments`
	Acknowledgements      []Acknowledgement    `json:acknowledgemnts`
	Last_status_change_at string               `json:last_status_change_at`
	LastStatusChangeBy    LastStatusChangeBy   `json:last_status_change_by`
	FirstTriggerLogEntry  FirstTriggerLogEntry `json:first_trigger_log_entry`
	EscalationPolicy      EscalationPolicy     `json:escalation_policy`
	Teams                 []Team               `json:teams`
	Urgency               string               `json:urgency`
}

type AlertCounts struct {
	All       uint   `json:all`
	Triggered string `json:triggered`
	Resolved  uint   `json:resolved`
}

type PendingAction struct {
	Type string
	at   string
}

type Entity struct {
	Id      string
	Type    string
	Summary string
	Self    string
	HtmlUrl string
}

type Service struct {
	Entity
}

type Priority struct {
	Entity
}

type Assignment struct {
	At       string
	Assignee Assignee
}

type Assignee struct {
	Entity
}

type Acknowledgement struct {
	At           string
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
