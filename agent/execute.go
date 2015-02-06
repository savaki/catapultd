package agent

type ExecuteFunc func(*Command) error

func NoopExecuteFunc(command *Command) error {
	return nil
}
