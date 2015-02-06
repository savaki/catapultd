package agent

import "fmt"

type PollFunc func() (*Command, error)

func NoopPollFunc() (*Command, error) {
	return nil, fmt.Errorf("not-implemented yet")
}
