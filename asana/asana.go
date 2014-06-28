package asana

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type Client struct {
	key string
	hc  http.Client
}

func NewClient(key string) Client {
	return Client{key, http.Client{}}
}

func (a Client) Me() (User, error) {
	ud := userData {}
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
	Id int64
	Name string
}

type User struct {
	Id int
	Name string
	Email string
	Photos map[string]string `json:"photos"`
	Workspaces []Workspace
/*
{
  "data": {
    "id": 184808224339,
    "name": "Brian McCallister",
    "email": "brianm@skife.org",
    "photo": {
      "image_21x21": "https://s3.amazonaws.com/profile_photos/184808224339.tPSTUiUUM4eBivXaQtE8_21x21.png",
      "image_27x27": "https://s3.amazonaws.com/profile_photos/184808224339.tPSTUiUUM4eBivXaQtE8_27x27.png",
      "image_36x36": "https://s3.amazonaws.com/profile_photos/184808224339.tPSTUiUUM4eBivXaQtE8_36x36.png",
      "image_60x60": "https://s3.amazonaws.com/profile_photos/184808224339.tPSTUiUUM4eBivXaQtE8_60x60.png",
      "image_128x128": "https://s3.amazonaws.com/profile_photos/184808224339.tPSTUiUUM4eBivXaQtE8_huge.jpeg"
    },
    "workspaces": [
      {
        "id": 13438604102030,
        "name": "Brian"
      },
      {
        "id": 498346170860,
        "name": "Personal Projects"
      }
    ]
  }
}
*/
}
