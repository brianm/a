package asana

import (
	"encoding/json"
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

func (a Client) Me() (User, error) {
	ud := userData{}
	url := "https://app.asana.com/api/1.0/users/me"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ud.Data, fmt.Errorf("Cannot create request GET %s", url, err)
	}
	req.SetBasicAuth(a.key, "")
	resp, err := a.hc.Do(req)
	if err != nil {
		return ud.Data, fmt.Errorf("Cannot execute request GET %s: %s", url, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ud.Data, fmt.Errorf("error reading body: %s", err)
	}
	err = json.Unmarshal(body, &ud)
	return ud.Data, err
}

type userData struct {
	Data User
}

type Workspace struct {
	Id   int64
	Name string
}

type Photos struct {
	Image_21x21 string `json:"image_21x21"`
	Image_27x27 string
	Image_36x36 string
	Image_60x60 string
	Image_128x128 string
}

type User struct {
	Id         int
	Name       string
	Email      string
	Photos     Photos `json:"photo"`
	Workspaces []Workspace
}
