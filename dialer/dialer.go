package dialer

import (
	"io/ioutil"
	"log"
	"net/http"
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
 * @var		d	*Dialer
 * @global
 */
func (d *Dialer) Get(url string) error {
	log.Printf("(GET)Dailing url: %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header = d.HeaderList
	res, err := d.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	log.Print(string(body))

	//return res, err
	return nil
}
