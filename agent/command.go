package agent

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	ErrInvalidCommand  = fmt.Errorf("invalid command")
	ErrCommandNotFound = fmt.Errorf("command not found")
)

type Context struct {
	Dir string
	Log LogFunc
}

func (c *Context) IsWithinContextDir(filename string) bool {
	prefix, _ := filepath.Abs(c.Dir)

	return strings.HasPrefix(filename, prefix)
}

//
func (c *Context) Filename(path string) (string, error) {
	filename, err := filepath.Abs(fmt.Sprintf("%s/%s", c.Dir, path))
	if err != nil {
		return filename, ErrInvalidCommand
	}

	if !c.IsWithinContextDir(filename) {
		// ignore requests for file outside the directory specified by context
		return filename, ErrInvalidCommand
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return filename, ErrCommandNotFound
	}

	return filename, nil
}

// find the absolute path to the command specified
func (c *Context) Command(command string, args ...string) (*exec.Cmd, error) {
	name, err := c.Filename("bin/" + command)
	if err != nil {
		return nil, err
	}
	return exec.Command(name, args...), nil
}

type Command struct {
	RequestId string          `json:"request-id"`
	Command   string          `json:"command"`
	Args      []string        `json:"args"`
	Payload   json.RawMessage `json:"payload"`
}

func (c *Command) Execute(ctx *Context) error {
	return nil
}
