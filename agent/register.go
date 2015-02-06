package agent

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/savaki/catapultd/config"
)

//---------------------------------------------------------------

type RegisterRequest struct {
	AgentId   string `json:"agent-id"`
	AuthToken string `json:"auth-token"`
}

type RegisterResponse struct {
	AgentId   string `json:"agent-id"`
	AuthToken string `json:"auth-token"`
}

func register(url, dir string) (agentId string, authToken string, err error) {
	// create a request
	request := RegisterRequest{}

	// load our existing config if it exists
	if cfg, err := config.LoadFile(dir + "/catapultd.conf"); err == nil {
		request.AgentId = cfg.AgentId
		request.AuthToken = cfg.AuthToken
	}

	// send the request to the server
	var data []byte
	data, err = json.Marshal(request)
	if err != nil {
		return
	}

	var resp *http.Response
	resp, err = http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// decode the response
	response := &RegisterResponse{}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return
	}

	agentId = response.AgentId
	authToken = response.AuthToken
	return
}
