package avschema

import (
	"encoding/json"
	"log"
	"time"
)

type Alarm struct {
	TimeStamp uint64  `json:"timeStamp"`
	Enriched  bool    `json:"enriched"`
	Message   Message `json:"message"`
}

type Message struct {
	RuleIntent               string        `json:"rule_intent"`
	AppType                  string        `json:"app_type"`
	AlarmEventsCount         uint16        `json:"alarm_events_count"`
	Sources                  []Source      `json:"sources"`
	AlarmSourceCountries     []string      `json:"alarm_source_countries"`
	SourceUsername           string        `json:"source_username"`
	AlarmSensorSources       []string      `json:"alarm_sensor_sources"`
	TimestampOccured         string        `json:"timestamp_occured"`
	UUID                     string        `json:"uuid"`
	NeedsEnrichment          bool          `json:"needs_enrichment"`
	SourceOrganisation       string        `json:"source_organisation"`
	AlarmSourceCities        []string      `json:"alarm_source_cities"`
	EventType                string        `json:"event_type"`
	AccountName              string        `json:"account_name"`
	RuleMethod               string        `json:"rule_method"`
	PriorityLabel            string        `json:"priority_label"`
	Suppressed               string        `json:"suppressed"`
	AppId                    string        `json:"app_id"`
	HasAlarm                 string        `json:"has_alarm"`
	SourceName               string        `json:"source_name"`
	Events                   []Event       `json:"events"`
	TimestampReceived        string        `json:"timestamp_received"`
	RuleStrategy             string        `json:"rule_strategy"`
	RuleName                 string        `json:"rule_name"`
	TimestampReceivedIso8601 time.Time     `json:"timestamp_received_iso8601"`
	PacketData               []string      `json:"packet_data"`
	Destinations             []Destination `json:"destinations"`
	AlarmSources             []string      `json:"alarm_sources"`
	AlarmSourceNames         []string      `json:"alarm_source_names"`
	HighlightFields          []string      `json:"highlight_fields"`
	Priority                 string        `json:"priority"`
	AlarmSourceLongitudes    []string      `json:"alarm_source_longitudes"`
	AlarmSourceLatitudes     []string      `json:"alarm_source_latitudes"`
	RuleId                   string        `json:"rule_id"`
	AlarmSourceOrganisations []string      `json:"alarm_source_organisations"`
	SensorUUID               string        `json:"sensor_uuid"`
	TimestampOccuredIso8601  time.Time     `json:"timestamp_occured_iso8601"`
	Transient                bool          `json:"transient"`
	EventName                string        `json:"event_name"`
	AlarmSourceZones         []string      `json:"alarm_source_zones"`
	PacketType               string        `json:"packet_type"`
	Status                   string        `json:"status"`
}

type Destination struct {
	DestinationCanonical string `json:"destination_canonical"`
	DestinationAddress   string `json:"destination_address"`
	DestinationHostname  string `json:"destination_hostname"`
	EventCount           uint32 `json:"event_count"`
	DestinationName      string `json:"destination_name"`
}

type Source struct {
	SourceOrganisation string `json:"source_organisation"`
	SourceCountry      string `json:"source_country"`
	EventCount         int32  `json:"event_count"`
	SourceAddress      string `json:"source_address"`
	SourceCanonical    string `json:"source_canonical"`
	SourceName         string `json:"source_name"`
}

type Event struct {
	Source
	WasFuzzied               bool      `json:"was_fuzzied"`
	AppType                  string    `json:"app_type"`
	EventSeverity            string    `json:"event_severity"`
	TimestampOccurred        string    `json:"timestamp_occured"`
	UUID                     string    `json:"uuid"`
	BaseEventCount           uint      `json:"base_event_count"`
	UsedHint                 bool      `json:"used_hint"`
	AppId                    string    `json:"app_id"`
	SourceASN                string    `json:"source_asn"`
	WasGuessed               bool      `json:"was_guessed"`
	TimestampReceived        string    `json:"timestamp_received"`
	SourceLatitude           string    `json:"source_latitude"`
	TimestampReceivedIso8601 time.Time `json:"timestamp_received_iso8601"`
	SourceZone               string    `json:"source_zone"`
	UserResource             string    `json:"user_resource"`
	SensorUUID               string    `json:"sensor_uuid"`
	Transient                bool      `json:"transient"`
	SourceLongitude          string    `json:"source_longitude"`
	RepDeviceRuleId          string    `json:"rep_device_rule_id"`
	EventName                string    `json:"event_name"`
	SourceRegisteredCountry  string    `json:"source_registered_country"`
	PacketType               string    `json:"packet_type"`
	PluginVersion            string    `json:"plugin_version"`
	Log                      string    `json:"log"`
	TimeStart                string    `json:"time_start"`
	EventDescription         string    `json:"event_description"`
	SourceUsername           string    `json:"source_username"`
	NeedsEnrichment          bool      `json:"needs_enrichment"`
	SourceUserId             string    `json:"source_userid"`
	AccountName              string    `json:"account_name"`
	TimeEnd                  string    `json:"time_end"`
	Suppressed               string    `json:"suppressed"`
	HasAlarm                 string    `json:"has_alarm"`
	PluginDeviceType         string    `json:"plugin_device_type"`
	SourceCity               string    `json:"source_city"`
	PluginDevice             string    `json:"plugin_device"`
	HighlightFields          []string  `json:"highlight_fields"`
	AppName                  string    `json:"app_name"`
	EventAction              string    `json:"event_action"`
	TimestampOccuredIso8601  time.Time `json:"timestamp_occured_iso8601"`
	AccountId                string    `json:"account_id"`
	Plugin                   string    `json:"plugin"`
	Application              string    `json:"application"`
	InAlarms                 []string  `json:"in_alarms"`
	RepDeviceVersion         string    `json:"rep_device_version"`
	SourceUserPrivileges     string    `json:"source_user_privileges"`
}

func (ir *Alarm) ToPrettyString() string {
	b, err := json.MarshalIndent(*ir, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func (ir *Alarm) ToString() string {
	b, err := json.Marshal(*ir)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
