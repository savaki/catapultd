package agent

import "fmt"

type MockLogger struct {
	Logs []string
}

func (m *MockLogger) Log(format string, a ...interface{}) {
	text := fmt.Sprintf(format, a...)

	if m.Logs == nil {
		m.Logs = []string{}
	}

	m.Logs = append(m.Logs, text)
}
