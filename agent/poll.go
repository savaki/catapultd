package agent

type PollFunc func() (*Command, error)
