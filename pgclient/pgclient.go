package pgclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

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
		},
		Url: p.Endpoint,
	}
	p.Dialer.InitClient()
}

type requestFilter struct {
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

type RequestFilter interface{}

func New(since, until, DateRange) RequestFilter {
	return requestFilter{name, 0}
}

func ConvertToQueryString(v RequestFilter) {
	var queryString []byte

}

/**
 * List.
 *
 * @author	Unknown
 * @var		p	*PgClien
 * @global
 */
func (p *PgClient) List(entity string, filters RequestFilter) (ps.IncidentsResponse, error) {
	var incident ps.IncidentsResponse

	if entity == "" {
		return incident, errors.New("no entity defined")
	}
	if filters {

	}

	body, err := p.Dialer.Get(fmt.Sprintf("%s/%s", p.Endpoint, entity))
	if err != nil {
		log.Fatal(err)
		return incident, err
	}

	json.Unmarshal(body, &incident)

	return incident, nil
}
