package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/savaki/catapultd/agent"
)

func Opts(c *cli.Context) *agent.Options {
	return &agent.Options{
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
	agent, err := agent.New(opts)
	check(err)

	for {
		command, err := agent.Poll()
		check(err)

		agent.Execute(command)
	}
}
