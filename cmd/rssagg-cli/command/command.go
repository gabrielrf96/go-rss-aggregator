package command

import (
	"os"
	"os/exec"
)

func runCmd(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)

	return run(cmd)
}

func run(cmd *exec.Cmd) error {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
