package entities

import "strings"

type LogoInfo struct {
	TextColor int
	Logo      []string
	BlankRow  string
}

var LogosMap = map[string]LogoInfo{
	"arch": {
		TextColor: 34,
		Logo: []string{
			"   \033[34m/\\\033[0m   ",
			"  \033[34m/\\ \\\033[0m  ",
			" \033[34m/ .\\ \\\033[0m ",
			"\033[34m/.'  '.\\\033[0m",
		},
		BlankRow: "\t\t",
	},

	"ubuntu": {
		TextColor: 31,
		Logo: []string{
			"  \033[33m/\033[31m--( )\033[0m",
			"\033[31m( )    \033[31m|\033[0m",
			" \033[33m\\     \033[31m/\033[0m",
			"   \033[31m--\033[33m( )\033[0m",
			"\t",
		},
		BlankRow: "\t",
	},

	"tux": {
		TextColor: 37,
		Logo: []string{
			"   \033[37m.-,\033[0m  ",
			"   \033[37moo\033[37m|\033[0m  ",
			"  /\033[33mv \033[37m\\\033[0m  ",
			" \033[33m(\\\033[37m_^\033[33m/)\033[0m ",
			"",
		},
		BlankRow: "\t",
	},
}

const baseLogo = "tux"

func GetLogo(logoName string) (LogoInfo, error) {
	if logoName == "" {
		distro, err := runCmd(`grep '^ID=' /etc/os-release | cut -d= -f2 | tr -d '"'`)
		if err != nil {
			return LogoInfo{}, err
		}

		logoName = strings.ToLower(strings.TrimSpace(distro))
	}

	if logo, ok := LogosMap[logoName]; ok {
		return logo, nil
	}

	return LogosMap[baseLogo], nil
}
