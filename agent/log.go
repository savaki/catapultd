package agent

import "log"

type LogFunc func(format string, a ...interface{}) error

func StdoutLogFunc(format string, a ...interface{}) error {
	log.Printf(format, a...)
	return nil
}
