package main

type Agent struct {
	agentId   string
	authToken string
	log       LogFunc
}

type Link struct {
	Href string `json:"href"`
}

func NewAgent(opts *Options) (*Agent, error) {
	pollFunc, err := makePollFunc(opts.Url)
	if err != nil {
		return nil, err
	}

	logFunc, err := makeLogFunc(opts.Url)
	if err != nil {
		return nil, err
	}

	return &Agent{
		agentId:   agent,
		authToken: authToken,
		log:       logFunc,
		poll:      pollFunc,
	}
}

func makePollFunc(url string) (PollFunc, error) {
	return func() (*Command, error) {
		return nil, nil
	}, nil
}

func (a *Agent) Execute(command *Command) error {
	return nil
}
