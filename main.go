package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/savaki/catapultd/config"
)

type Options struct {
	Dir string
	Url string
}

func Opts(c *cli.Context) *Options {
	return &Options{
		Dir: c.String(FlagDir),
		Url: c.String(FlagUrl),
	}
}

const (
	FlagUrl = "url"
	FlagDir = "dir"
)

func main() {
	app := cli.NewApp()
	app.Name = "catapultd"
	app.Usage = "launcher app for the docker cloud tool"
	app.Author = "Matt Ho"
	app.Version = "SNAPSHOT"
	app.Flags = []cli.Flag{
		cli.StringFlag{FlagDir, fmt.Sprintf("%s/.catapultd", os.Getenv("HOME")), "local storage directory", "DIR"},
		cli.StringFlag{FlagUrl, "http://localhost:8000", "yellow pages url", "URL"},
	}
	app.Action = Run
	app.Run(os.Args)
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Run(c *cli.Context) {
	opts := Opts(c)

	// Logic:
	// 1. fetch log server and controller urls from directory server
	// 2. push logs to log server
	// 3. register with controller (poll until registration successful)
	// 4. poll for commands

	cfg, err := config.LoadFile(opts.Dir + "/catapultd.conf")
	check(err)

	agent := NewAgent(cfg)

	err = agent.Register(opts.Url)
	check(err)

	for {
		command, err := agent.Poll()
		check(err)

		agent.Execute(command)
	}
}
