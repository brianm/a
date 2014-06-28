package asana

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// A remote client to Asana
type Client struct {
	key string
	hc  http.Client
}

// Created a new client with a specified API Key
func NewClient(key string) Client {
	return Client{key, http.Client{}}
}

// Fetch the user owning the API key this
// client was created with
func (a Client) Me() (User, error) {
	ud := userData{}
	url := "https://app.asana.com/api/1.0/users/me"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ud.Data, fmt.Errorf("Cannot create request GET %s",
			url, err)
	}
	req.SetBasicAuth(a.key, "")
	resp, err := a.hc.Do(req)
	if err != nil {
		return ud.Data, fmt.Errorf("Cannot execute request GET %s: %s",
			url, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ud.Data, fmt.Errorf("error reading body: %s", err)
	}
	err = json.Unmarshal(body, &ud)
	return ud.Data, err
}

// Fetch a user by id
func (a Client) User(id int64) (User, error) {
	ud := userData{}
	url := fmt.Sprintf("https://app.asana.com/api/1.0/users/%d", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ud.Data, fmt.Errorf("Cannot create request GET %s",
			url, err)
	}
	req.SetBasicAuth(a.key, "")
	resp, err := a.hc.Do(req)
	if err != nil {
		return ud.Data, fmt.Errorf("Cannot execute request GET %s: %s",
			url, err)
	}

	if resp.StatusCode != 200 {
		return ud.Data, fmt.Errorf("Error response: %d", resp.StatusCode)
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
	Image_21x21   string
	Image_27x27   string
	Image_36x36   string
	Image_60x60   string
	Image_128x128 string
}

type User struct {
	Id         int
	Name       string
	Email      string
	Photos     Photos `json:"photo"`
	Workspaces []Workspace
}
