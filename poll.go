package main

type Command struct {
}

type PollFunc func() (*Command, error)
