package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

//---------------------------------------------------------------

type Hosts struct {
	LogServer Link `json:"logger"`
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

//---------------------------------------------------------------

type RegisterRequest struct {
	AgentId string `json:"agent-id"`
	Token   string `json:"token"`
}

type RegisterResponse struct {
	AgentId string `json:"agent-id"`
	Token   string `json:"token"`
}

func (a *Agent) Register(url string) error {
	// Logic:
	// 1. use #discover to find the url of our api server
	// 2.

	hosts, err := discover(url)
	if err != nil {
		return err
	}
	a.Hosts = hosts

	request := RegisterRequest{
		AgentId: a.AgentId,
		Token:   a.Token,
	}
	data, err := json.Marshal(request)
	if err != nil {
		return err
	}

	// attempt to register on the network
	resp, err := http.Post(a.Hosts.ApiServer.Href, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// decode the response
	response := &RegisterResponse{}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return err
	}

	a.AgentId = response.AgentId
	a.Token = response.Token

	return nil
}
