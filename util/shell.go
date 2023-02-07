package util

import (
	"bytes"
	"os/exec"
)

func ShellCommand(appName string, args []string) (out string, errmsg string, err error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	app, _ := exec.LookPath("dot")

	cmd := exec.Command(app, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cmderr := cmd.Run()
	if cmderr != nil {
		return "", "", cmderr
	}

	return stdout.String(), stderr.String(), nil
}
