package asana

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// A remote client to Asana
type Client struct {
	key string
	hc  http.Client
	Me  User
}

type Task struct {
	Id   int64
	Name string
}

type taskData struct {
	Data []Task
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

// Created a new client with a specified API Key
func NewClient(key string) (Client, error) {
	c := Client{key, http.Client{}, User{}}
	me, err := c.User("me")
	if err != nil {
		return c, err
	}

	c.Me = me
	return c, nil
}

func (c Client) request(method string, uri string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return req, fmt.Errorf("Cannot create request %s %s", method, uri, err)
	}
	req.SetBasicAuth(c.key, "")
	return req, nil
}

// Fetch a user by id
func (a Client) User(id interface{}) (User, error) {
	ud := userData{}
	uri := fmt.Sprintf("https://app.asana.com/api/1.0/users/%v", id)
	req, err := a.request("GET", uri, nil)
	if err != nil {
		return ud.Data, err
	}

	resp, err := a.hc.Do(req)
	if err != nil {
		return ud.Data, fmt.Errorf("Cannot execute request GET %s: %s", uri, err)
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

func (c Client) Tasks(w Workspace) ([]Task, error) {
	td := taskData{}
	url := fmt.Sprintf("https://app.asana.com/api/1.0/workspaces/%d/tasks?assignee=%d",
		w.Id, c.Me.Id)

	req, err := c.request("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return td.Data, fmt.Errorf("error reading body: %s", err)
	}
	err = json.Unmarshal(body, &td)
	return td.Data, err
}
