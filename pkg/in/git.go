package in

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type GitRunner interface {
	Run(args ...string) error
}

func NewRunner() GitRunner {
	return DefaultRunner{}
}

type DefaultRunner struct {
}

func (r DefaultRunner) Run(args ...string) error {
	cmd := "git"
	command := exec.Command(cmd, args...)
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		return fmt.Errorf("'git %s': %s", strings.Join(args, " "), err.Error())
	}
	return nil
}
