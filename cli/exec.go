package cliq

import (
	"os"
	"os/exec"
)

func Command(str string, s ...string) {
	cmd := exec.Command(str, s...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
func CommandOutput(str string, s ...string) string {
	output := exec.Command(str, s...)
	bit, err := output.Output()
	Check(err, "error command output "+str)
	return string(bit)
}
