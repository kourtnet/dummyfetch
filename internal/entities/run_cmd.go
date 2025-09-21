package entities

import "os/exec"

func runCmd(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)

	res, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(res), nil
}
