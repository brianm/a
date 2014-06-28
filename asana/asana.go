package asana

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	key string
	hc  http.Client
}

func NewClient(key string) Client {
	return Client{key, http.Client{}}
}

func (a Client) Me() ([]byte, error) {
	url := "https://app.asana.com/api/1.0/users/me"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Cannot create request GET %s", url, err)
	}
	req.SetBasicAuth(a.key, "")
	resp, err := a.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Cannot execute request GET %s: %s", url, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %s", err)
	}
	return body, nil
}
