package dialer

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
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
func (d *Dialer) Get(url string) ([]byte, error) {
	log.Printf("(GET)Dailing url: %s", url)
	req, err := http.NewRequest("GET", url, nil)
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

	//log.Print(string(body))

	return body, err
	//return nil
}
