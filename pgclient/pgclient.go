package pgclient

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	d "github.com/sahilsk/hundun/dialer"
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

/**
 * List.
 *
 * @author	Unknown
 * @since	v0.0.1
 * @version	v1.0.0	Sunday, December 29th, 2019.
 * @global
 * @param	entity	string
 * @return	void
 */
func (p *PgClient) List(entity string) error {
	if entity == "" {
		return errors.New("no entity defined")
	}

	err := p.Dialer.Get(fmt.Sprintf("%s/%s", p.Endpoint, "incidents"))
	if err != nil {
		log.Fatal(err)
	}

	return nil

}
