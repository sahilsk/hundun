package avclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	log "github.com/sahilsk/hundun/logger"

	avs "github.com/sahilsk/hundun/avclient/schema"

	d "github.com/sahilsk/hundun/dialer"
)

var logger *log.Clogger

type AVClient struct {
	ApiKey string
	// eg. https://paxos-trust-company-llc.alienvault.cloud/api/1.0/
	Endpoint string
	Dialer   *d.Dialer
	Cookie   string
	Verbose  bool
}

func NewAVClient(apikey string, endpoint string, cookie string, verbose bool) *AVClient {

	avclient := &AVClient{
		ApiKey:   apikey,
		Endpoint: endpoint,
		Cookie:   cookie,
		Verbose:  verbose,
	}

	logger = log.NewLogger(verbose)

	avclient.Init()
	return avclient
}

//Init initializes pagerduty client with headers and pg endpoint
func (av *AVClient) Init() {

	av.Dialer = &d.Dialer{
		HeaderList: http.Header{
			//			"Authorization": []string{fmt.Sprintf("Token token=%s", av.ApiKey)},
			"Content-Type": []string{"application/json"},
			"Accept":       []string{"application/json", "text/plain", "*/*"},
			"User-Agent":   []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36"},
			"Cookie":       []string{av.Cookie},
		},
		Url:     av.Endpoint,
		Verbose: av.Verbose,
	}
	av.Dialer.InitClient()
}

//Get sends a GET request to pagerduty to fetch the specified entity
func (p *AVClient) Get(entity string, id string) (interface{}, error) {

	requestURL := fmt.Sprintf("%s/%s/%s", p.Endpoint, entity, id)

	if entity == "" {
		return nil, errors.New("no entity passed")
	}

	body, err := p.Dialer.Get(requestURL, url.Values{})
	if err != nil {
		return nil, err
	}

	type AlarmResponse map[string]avs.Alarm

	var alarmResponse AlarmResponse
	if err := json.Unmarshal(body, &alarmResponse); err != nil {
		logger.Fatal("%s", err)
	}
	avalarm := alarmResponse[id]
	logger.Info("%v", avalarm)
	return avalarm, err

}
