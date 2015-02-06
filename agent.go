package main

import "github.com/savaki/catapultd/config"

type Agent struct {
	Hosts   *Hosts
	AgentId string
	Token   string
	Poll    PollFunc
	Log     LogFunc
}

type Link struct {
	Href string `json:"href"`
}

func NewAgent(cfg *config.Config) *Agent {
	agent := &Agent{}

	if cfg != nil {
		agent.AgentId = cfg.AgentId
		agent.Token = cfg.Token
	}

	return agent
}

func (a *Agent) Execute(command *Command) error {
	return nil
}
