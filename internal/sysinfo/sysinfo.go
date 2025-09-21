// Package sysinfo provides function for gathering system info
package sysinfo

import (
	"os/exec"

	"github.com/kourtnet/dummyfetch/internal/args"
	"github.com/kourtnet/dummyfetch/internal/logos"
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
			name: `Uptime`,
			command: `awk '
				function plural(n, singular, plural_form) {
  			return (n == 1) ? singular : plural_form
				}
				{
  				s=int($1);
  				d=int(s/86400);
  				h=int((s%86400)/3600);
  				m=int((s%3600)/60);
  				str = ""
  				if(d > 0) str = str d " " plural(d, "day", "days") ", "
  				if(h > 0) str = str h " " plural(h, "hour", "hours") ", "
  				if(m > 0) str = str m " " plural(m, "minute", "minutes") ", "
  				sub(/, $/, "", str)
  				print str
				}' /proc/uptime`,
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

func ResolveIcon() ([]string, error) {
	distro, err := runCmd(`grep '^ID=' /etc/os-release | cut -d= -f2 | tr -d '"'`)
	if err != nil {
		return nil, err
	}

	return logos.Get(distro), nil
}
