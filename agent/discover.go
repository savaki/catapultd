package agent

import (
	"encoding/json"
	"net/http"
)

//---------------------------------------------------------------

type Link struct {
	Href string `json:"href"`
}

type Hosts struct {
	LogServer Link `json:"log-server"`
	ApiServer Link `json:"api-server"`
}

func discover(url string) (*Hosts, error) {
	hosts := &Hosts{}

	// fetch the directory info
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// parse the results
	err = json.NewDecoder(resp.Body).Decode(hosts)
	if err != nil {
		return nil, err
	}

	return hosts, nil
}
