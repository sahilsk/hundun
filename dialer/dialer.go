package dialer

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Dialer struct {
	HeaderList http.Header
	Url        string
	Client     *http.Client
}

/**
 * @var		d	*Diale
 * @global
 */
func (d *Dialer) InitClient() {
	d.Client = &http.Client{}
}

/**
 * @var		d	*Diale
 * @global
 */
func (d *Dialer) AddHeader(k, v string) {
	d.HeaderList.Add(k, v)
	log.Printf("Key: %s:%s", k, d.HeaderList.Get(k))
}

/**
* jsonFactory.
*
* @author	Unknown
* @since	v0.0.1
* @version	v1.0.0	Wednesday, January 1st, 2020.
* @global
* @param	s	string
* @return	mixed
* @example:
* req, _ := http.NewRequest("POST", "/", jsonFactory(`{"something": "hello"}`))
* 	req, _ = http.NewRequest("POST", "/", jsonFactory(`{}`))
	req, _ = http.NewRequest("POST", "/", jsonFactory(`{"something": ["test", "data"]}`))

*/
func jsonFactory(s string) io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader(s))
}

/**
 * @var		d	*Dialer
 * @global
 */
func (d *Dialer) Get(endpoint string, params url.Values) ([]byte, error) {
	log.Printf("(GET)Dailing url: %s", endpoint)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", endpoint, params.Encode()), nil)
	if err != nil {
		return nil, err
	}
	req.Header = d.HeaderList
	res, err := d.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	log.Print(string(body))

	if res.StatusCode > 399 {
		return nil, errors.New(string(body))
	}

	return body, err
	//return nil
}

/**
 * @var		d	*Diale
 * @global
 */
func (d *Dialer) Put(endpoint string, params url.Values, body []byte) ([]byte, error) {
	log.Printf("(PUT)Dailing url: %s", fmt.Sprintf("%s", endpoint))
	log.Printf("Payload: %s", string(body))
	//req, err := http.NewRequest("PUT", fmt.Sprintf("%s?%s", endpoint, params.Encode()), bytes.NewReader(body))
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s", endpoint), bytes.NewReader(body))

	if err != nil {
		return nil, err
	}
	req.Header = d.HeaderList
	res, err := d.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	log.Printf("Status Code: %d", res.StatusCode)

	body, err = ioutil.ReadAll(res.Body)
	log.Printf("body: %s", body)

	if res.StatusCode > 399 {
		return nil, errors.New(string(body))
	}
	return body, err

}
