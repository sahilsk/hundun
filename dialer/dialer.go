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

func printAndReturn(r io.Reader, sc int) ([]byte, error) {

	body, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if sc > 399 {
		return nil, errors.New(string(body))
	} else {
		log.Print(string(body))
	}

	return body, err
}

/**
 * @var		d	*Dialer
 * @global
 */
func (d *Dialer) Get(endpoint string, params url.Values) ([]byte, error) {
	requestURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())
	log.Printf("(GET)Dailing url: %s", requestURL)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header = d.HeaderList
	res, err := d.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return printAndReturn(res.Body, res.StatusCode)
}

/**
 * @var		d	*Diale
 * @global
 */
func (d *Dialer) Put(endpoint string, params url.Values, body []byte) ([]byte, error) {
	requestURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())
	log.Printf("(PUT)Dailing url: %s", requestURL)
	log.Printf("Payload: %s", string(body))

	req, err := http.NewRequest(http.MethodPut, requestURL, bytes.NewReader(body))

	if err != nil {
		return nil, err
	}
	req.Header = d.HeaderList
	res, err := d.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return printAndReturn(res.Body, res.StatusCode)
}
