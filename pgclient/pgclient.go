package pgclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

	d "github.com/sahilsk/hundun/dialer"
	ps "github.com/sahilsk/hundun/pgclient/schema"
)

type PgClient struct {
	ApiKey   string
	Endpoint string
	Dialer   *d.Dialer
}

func (p *PgClient) Init() {
	log.Print("init pg client")
	p.Dialer = &d.Dialer{
		HeaderList: http.Header{
			"Authorization": []string{fmt.Sprintf("Token token=%s", p.ApiKey)},
			"Accept":        []string{"application/vnd.pagerduty+json;version=2"},
			"From":          []string{"smeena@itbit.com"},
		},
		Url: p.Endpoint,
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

func (rf *RequestFilter) ToQueryString() string {
	log.Printf("ToQueryString called")
	return "sdf"
}

/**
 * List.
 *
 * @author	Unknown
 * @var		p	*PgClien
 * @global
 */
func (p *PgClient) List(entity string, filters url.Values) (ps.IncidentsResponse, error) {
	var incident ps.IncidentsResponse

	if entity == "" {
		return incident, errors.New("no entity defined")
	}

	body, err := p.Dialer.Get(fmt.Sprintf("%s/%s", p.Endpoint, entity), filters)
	if err != nil {
		log.Fatal(err)
		return incident, err
	}

	if err := json.Unmarshal(body, &incident); err != nil {
		log.Fatal(err)
	}

	return incident, nil
}

func (p *PgClient) Get(entity string, id string) (ps.IncidentResponse, error) {
	var incident ps.IncidentResponse

	if entity == "" {
		return incident, errors.New("no entity defined")
	}

	body, err := p.Dialer.Get(fmt.Sprintf("%s/%s/%s", p.Endpoint, entity, id), url.Values{})
	if err != nil {
		log.Fatal(err)
		return incident, err
	}

	if err := json.Unmarshal(body, &incident); err != nil {
		log.Fatal(err)
	}

	return incident, nil
}

func (p *PgClient) Put(entity string, id string, params url.Values, payload []byte) (ps.IncidentResponse, error) {
	var incident ps.IncidentResponse

	if entity == "" {
		return incident, errors.New("no entity defined")
	}

	body, err := p.Dialer.Put(fmt.Sprintf("%s/%s/%s", p.Endpoint, entity, id),
		params, payload)
	if err != nil {
		log.Fatal(err)
		return incident, err
	}

	if err := json.Unmarshal(body, &incident); err != nil {
		log.Fatal(err)
	}

	return incident, nil
}
