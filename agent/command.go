package agent

import (
	"encoding/json"
	"fmt"
	"os"
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

func (c *Context) Path(command string) (string, error) {
	if strings.Contains(command, "/") {
		return "", ErrInvalidCommand
	}

	filename, err := filepath.Abs(fmt.Sprintf("%s/%s", c.Dir, command))
	if err != nil {
		return "", ErrInvalidCommand
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return "", ErrCommandNotFound
	}

	return filename, nil
}

type Command struct {
	RequestId string          `json:"request-id"`
	Command   string          `json:"command"`
	Payload   json.RawMessage `json:"payload"`
}

func (c *Command) Execute(ctx *Context) error {
	return nil
}
