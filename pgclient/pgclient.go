package pgclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	log "github.com/sahilsk/hundun/logger"

	d "github.com/sahilsk/hundun/dialer"
	ps "github.com/sahilsk/hundun/pgclient/schema"
)

var logger *log.Clogger

type PgClient struct {
	ApiKey   string
	Endpoint string
	Email    string
	Dialer   *d.Dialer
	Verbose  bool
}

func NewPgClient(apikey string, endpoint string, email string, verbose bool) *PgClient {

	pgclient := &PgClient{
		ApiKey:   apikey,
		Endpoint: endpoint,
		Email:    email,
		Verbose:  verbose,
	}

	logger = log.NewLogger(verbose)

	pgclient.Init()
	return pgclient
}

//Init initializes pagerduty client with headers and pg endpoint
func (p *PgClient) Init() {

	p.Dialer = &d.Dialer{
		HeaderList: http.Header{
			"Authorization": []string{fmt.Sprintf("Token token=%s", p.ApiKey)},
			"Accept":        []string{"application/vnd.pagerduty+json;version=2"},
			"From":          []string{p.Email},
			"Content-Type":  []string{"application/json"},
		},
		Url:     p.Endpoint,
		Verbose: p.Verbose,
	}
	p.Dialer.InitClient()
}

type RequestFilter struct {
	Ready       bool
	Since       string
	Until       string
	DateRange   string
	Statuses    []string
	IncidentKey string
	ServiceIds  []string
	TeamIds     []string
	UserIds     []string
	Urgencies   []string
	TimeZone    string
	SortBy      string
	Include     []string
}

//List sends a GET request to fetcha all listings of the specified entity
//filtered through query parameters
func (p *PgClient) List(entity string, filters url.Values) (interface{}, error) {

	if entity == "" {
		return nil, errors.New("no entity defined")
	}

	requestURL := fmt.Sprintf("%s/%s", p.Endpoint, entity)
	body, err := p.Dialer.Get(requestURL, filters)
	if err != nil {
		return nil, err
	}

	switch entity {
	case "incidents":
		var incidents ps.IncidentsResponse
		if err := json.Unmarshal(body, &incidents); err != nil {
			if p.Verbose {
				logger.Info("%s", err)
			}
		}
		return incidents, err
	case "priorities":
		var priorities ps.Priorities
		if err := json.Unmarshal(body, &priorities); err != nil {
			logger.Info("%s", err)
		}
		return priorities, err
	default:
		return nil, errors.New("Unsupported entity")
	}

}

func (p *PgClient) ListChild(parent_entity string, child_entity string,
	identifier string, filters url.Values) (interface{}, error) {

	if parent_entity == "" {
		return nil, errors.New("no parent entity defined")
	}
	if child_entity == "" {
		return nil, errors.New("no child entity defined")
	}

	requestURL := fmt.Sprintf("%s/%s/%s/%s", p.Endpoint, parent_entity, identifier, child_entity)
	body, err := p.Dialer.Get(requestURL, filters)
	if err != nil {
		return nil, err
	}

	switch child_entity {
	case "notes":
		var notes ps.Notes
		if err := json.Unmarshal(body, &notes); err != nil {
			logger.Info("%s", err)
		}
		return notes, err
	default:
		return nil, errors.New("Unsupported entity")
	}

}

//Get sends a GET request to pagerduty to fetch the specified entity
func (p *PgClient) Get(entity string, id string) (interface{}, error) {

	requestURL := fmt.Sprintf("%s/%s/%s", p.Endpoint, entity, id)

	if entity == "" {
		return nil, errors.New("no entity passed")
	}

	body, err := p.Dialer.Get(requestURL, url.Values{})
	if err != nil {
		return nil, err
	}

	switch entity {
	case "incidents":
		var incident ps.IncidentResponse
		if err := json.Unmarshal(body, &incident); err != nil {
			logger.Info("%s", err)
		}
		return incident, err
	case "priorities":
		var priority ps.PriorityResponse
		if err := json.Unmarshal(body, &priority); err != nil {
			logger.Info("%s", err)
		}
		return priority, nil
	default:
		return nil, errors.New("Unsupported entity")
	}

}

//Put send a PUT request to pg endpoints along with query params and json payload
//Use this to make a PUT request to update 'entity' with data provided
func (p *PgClient) Put(entity string, id string, params url.Values, payload []byte) (interface{}, error) {
	requestURL := fmt.Sprintf("%s/%s/%s", p.Endpoint, entity, id)

	if entity == "" {
		return nil, errors.New("no entity passed")
	}

	body, err := p.Dialer.Put(requestURL, params, payload)
	if err != nil {
		return nil, err
	}

	switch entity {
	case "incidents":
		var incident ps.IncidentResponse
		if err := json.Unmarshal(body, &incident); err != nil {
			logger.Info("%s", err)
		}
		return incident, err
	case "priorities":
		var priority ps.Priority
		if err := json.Unmarshal(body, &priority); err != nil {
			logger.Info("%s", err)
		}
		return priority, nil
	default:
		return nil, errors.New("Unsupported entity")
	}
}

func (p *PgClient) PostChild(parent_entity string, child_entity string, id string, params url.Values, payload []byte) (interface{}, error) {

	if parent_entity == "" {
		return nil, errors.New("no parent entity defined")
	}
	if child_entity == "" {
		return nil, errors.New("no child entity defined")
	}
	requestURL := fmt.Sprintf("%s/%s/%s/%s", p.Endpoint, parent_entity, id, child_entity)

	body, err := p.Dialer.Post(requestURL, params, payload)
	if err != nil {
		return nil, err
	}

	switch child_entity {
	case "notes":
		var note ps.Note
		if err := json.Unmarshal(body, &note); err != nil {
			logger.Info("%s", err)
		}
		return note, err
	default:
		return nil, errors.New("Unsupported entity")
	}
}
