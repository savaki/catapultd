package agent

type Command struct {
}

type PollFunc func() (*Command, error)
