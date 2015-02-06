package agent

import "fmt"

type MockPoller struct {
	Commands []*Command
}

func (m *MockPoller) Add(command *Command) *MockPoller {
	if m.Commands == nil {
		m.Commands = []*Command{}
	}

	m.Commands = append(m.Commands, command)

	return m
}

func (m *MockPoller) Poll() (*Command, error) {
	if len(m.Commands) == 0 {
		return nil, fmt.Errorf("no more commands left in MockPoller")
	}

	command := m.Commands[0]
	m.Commands = m.Commands[1:]

	return command, nil
}
