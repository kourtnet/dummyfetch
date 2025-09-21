// Package sysinfo provides function for gathering system info
package sysinfo

import (
	"os/exec"

	"github.com/kourtnet/dummyfetch/internal/entities"
)

func runCmd(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)

	res, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func Fetch(args []entities.Arg) ([]entities.Arg, error) {
	for i := range args {
		output, err := runCmd(args[i].Command)
		if err != nil {
			return nil, err
		}

		args[i].Contents = output
	}

	return args, nil
}

func ResolveIcon() (entities.LogoInfo, error) {
	distro, err := runCmd(`grep '^ID=' /etc/os-release | cut -d= -f2 | tr -d '"'`)
	if err != nil {
		return entities.LogoInfo{}, err
	}

	return entities.GetLogo(distro), nil
}
