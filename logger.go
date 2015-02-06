package main

import "log"

type LogFunc func(format string, a ...interface{}) error

func DefaultLog(format string, a ...interface{}) error {
	_, err := log.Printf(format, a...)
	return err
}
