package agent

type Agent struct {
	agentId   string
	authToken string
	log       LogFunc
	Poll      PollFunc
	Execute   ExecuteFunc
}

type Link struct {
	Href string `json:"href"`
}

type Options struct {
	Url string
	Dir string
}

func New(opts *Options) (*Agent, error) {
	hosts, err := discover(opts.Url)
	if err != nil {
		return nil, err
	}

	pollFunc, err := makePollFunc(hosts.ApiServer.Href)
	if err != nil {
		return nil, err
	}

	logFunc, err := makeLogFunc(hosts.LogServer.Href)
	if err != nil {
		return nil, err
	}

	executeFunc, err := makeExecuteFunc(opts.Dir)
	if err != nil {
		return nil, err
	}

	agentId, authToken, err := register(hosts.ApiServer.Href, opts.Dir)
	if err != nil {
		return nil, err
	}

	return &Agent{
		agentId:   agentId,
		authToken: authToken,
		log:       logFunc,
		Poll:      pollFunc,
		Execute:   executeFunc,
	}, nil
}

func makePollFunc(url string) (PollFunc, error) {
	return func() (*Command, error) {
		return nil, nil
	}, nil
}

func makeLogFunc(url string) (LogFunc, error) {
	return StdoutLogFunc, nil
}

func makeExecuteFunc(dir string) (ExecuteFunc, error) {
	return NoopExecuteFunc, nil
}
