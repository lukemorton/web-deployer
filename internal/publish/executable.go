package publish

import (
	"fmt"
	"io"
	"os/exec"
)

func ensureExecutableInstalled(name string) error {
	_, err := exec.LookPath(name)

	if err != nil {
		return fmt.Errorf("Could not find %s executable, is it installed?", name)
	}

	return nil
}

func runExecutable(writer *io.PipeWriter, executable string, args ...string) error {
	cmd := exec.Command(executable, args...)
	cmd.Stdout = writer
	cmd.Stderr = writer
	return cmd.Run()
}

func runExecutableAndReturnOutput(writer *io.PipeWriter, executable string, args ...string) ([]byte, error) {
	cmd := exec.Command(executable, args...)
	cmd.Stderr = writer
	return cmd.Output()
}
