// Package entities provides entities used in the project: system info arguments and OS logos
package entities

import (
	"bufio"
	"os"
	"strings"
)

type Arg struct {
	Name     string
	Command  func() (string, error)
	Contents string
}

var ArgsList = map[string]Arg{
	"os": {
		Name: `OS`,
		Command: func() (string, error) {
			file, err := os.Open("/etc/os-release")
			if err != nil {
				panic(err)
			}

			defer file.Close()

			scanner := bufio.NewScanner(file)
			var res string

			for scanner.Scan() {
				text := scanner.Text()
				if strings.HasPrefix(text, "PRETTY_NAME=") {
					suffix := text[len("PRETTY_NAME="):]
					res = strings.Trim(suffix, `"`) + "\n"

					break
				}
			}

			if err := scanner.Err(); err != nil {
				return "", err
			}

			return res, nil
		},
	},

	"kernel": {
		Name: `Kernel`,
		Command: func() (string, error) {
			output, err := runCmd(`uname -r`)
			if err != nil {
				return "", err
			}

			return output, nil
		},
	},

	"shell": {
		Name: `Shell`,
		Command: func() (string, error) {
			output, err := runCmd(`basename $SHELL`)
			if err != nil {
				return "", err
			}

			return output, nil
		},
	},

	"terminal": {
		Name: `Terminal`,
		Command: func() (string, error) {
			output, err := runCmd(`echo $TERM`)
			if err != nil {
				return "", err
			}

			return output, nil
		},
	},

	"uptime": {
		Name: `Uptime`,
		Command: func() (string, error) {
			output, err := runCmd(`uptime -p | sed 's/^up //'`)
			if err != nil {
				return "", err
			}

			return output, nil
		},
	},

	"palette_bg": {
		Name: ``,
		Command: func() (string, error) {
			output := "\033[2D\033[40m   \033[41m   \033[42m   \033[43m   \033[44m   \033[45m   \033[46m   \033[47m   \033[0m\n"
			return output, nil
		},
	},

	"palette_fg": {
		Name: ``,
		Command: func() (string, error) {
			output := "\033[2D\033[90m███\033[91m███\033[92m███\033[93m███\033[94m███\033[95m███\033[96m███\033[97m███\033[0m\n"
			return output, nil
		},
	},

	"indent": {
		Name: ``,
		Command: func() (string, error) {
			return "\033[2D \n", nil
		},
	},
}
