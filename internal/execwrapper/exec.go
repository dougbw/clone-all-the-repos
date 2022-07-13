package execwrapper

import (
	"bytes"
	"fmt"
	"os/exec"
)

type cliCommandOutput struct {
	Stderr  []byte
	Stdout  []byte
	Command string
}

func (o *cliCommandOutput) Output() string {
	return string(o.Stdout)
}

func (o *cliCommandOutput) ErrorOutput() string {
	return string(o.Stderr)
}

type cliCommandError struct {
	exitCode int
	stderr   string
}

func (e *cliCommandError) Error() string {
	return fmt.Sprintf("%d - %s", e.exitCode, e.stderr)
}

func Exec(name string, args ...string) (out cliCommandOutput, error error) {

	cmd := exec.Command(name, args...)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err := cmd.Run()
	out.Stdout = outb.Bytes()
	out.Stderr = errb.Bytes()

	if err != nil {
		exitError, ok := err.(*exec.ExitError)
		if ok {
			// fmt.Printf("Exit code is %d\n", exitError.ExitCode())
			error = &cliCommandError{exitError.ExitCode(), errb.String()}
			return out, error
		}
	}

	return out, error

}
