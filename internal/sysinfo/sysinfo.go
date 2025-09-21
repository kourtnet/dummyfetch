// Package sysinfo provides function for gathering system info
package sysinfo

import (
	"os/exec"

	"github.com/kourtnet/dummyfetch/internal/args"
)

func runCmd(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)

	res, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func Fetch() ([]args.Arg, error) {
	commands := []struct {
		name    string
		command string
	}{
		{
			name:    `OS`,
			command: `grep '^PRETTY_NAME=' /etc/os-release | cut -d= -f2 | tr -d '"'`,
		},
		{
			name:    `Kernel`,
			command: `uname -r`,
		},
		{
			name:    `Shell`,
			command: `basename $SHELL`,
		},
		{
			name:    `Terminal`,
			command: `echo $TERM`,
		},
		{
			name:    `Uptime`,
			command: `uptime -p | sed 's/^up //'`,
		},
	}

	res := make([]args.Arg, len(commands))

	for i := range commands {
		output, err := runCmd(commands[i].command)
		if err != nil {
			return nil, err
		}

		res[i].Name = commands[i].name
		res[i].Contents = output
	}

	return res, nil
}
