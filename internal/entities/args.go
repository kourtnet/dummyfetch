// Package entities provides entities used in the project: system info arguments and OS logos
package entities

type Arg struct {
	Name     string
	Command  string
	Contents string
}

var ArgsList = map[string]Arg{
	"os": {
		Name:    `OS`,
		Command: `grep '^PRETTY_NAME=' /etc/os-release | cut -d= -f2 | tr -d '"'`,
	},

	"kernel": {
		Name:    `Kernel`,
		Command: `uname -r`,
	},

	"shell": {
		Name:    `Shell`,
		Command: `basename $SHELL`,
	},

	"terminal": {
		Name:    `Terminal`,
		Command: `echo $TERM`,
	},

	"uptime": {
		Name:    `Uptime`,
		Command: `uptime -p | sed 's/^up //'`,
	},
}
