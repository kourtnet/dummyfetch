package entities

import "strings"

type LogoInfo struct {
	TextColor int
	Logo      []string
	BlankRow  string
}

var logosMap = map[string]LogoInfo{
	"arch": {
		TextColor: 34,
		Logo: []string{
			"   \033[34m/\\\033[0m\t\t",
			"  \033[34m/\\ \\\033[0m\t\t",
			" \033[34m/ .\\ \\\033[0m\t\t",
			"\033[34m/.'  '.\\\033[0m\t",
		},
		BlankRow: "\t\t",
	},

	"ubuntu": {
		TextColor: 31,
		Logo: []string{
			"  \033[33m/\033[31m-'-( )\033[0m",
			"\033[31m( )    \033[31m|\033[0m",
			" \033[33m\\     \033[31m/\033[0m",
			"   \033[31m-.-\033[33m( )\033[0m",
			"\t",
		},
		BlankRow: "\t",
	},

	"linux": {
		TextColor: 37,
		Logo: []string{
			"  \033[37m.-,\033[0m\t",
			"  \033[37moo\033[37m|\033[0m\t",
			" /\033[33mv \033[37m\\\033[0m\t",
			"\033[33m(\\\033[37m_^\033[33m/)\033[0m\t",
			"\t",
		},
		BlankRow: "\t",
	},
}

const baseLogo = "linux"

func GetLogo(name string) LogoInfo {
	nameF := strings.ToLower(strings.TrimSpace(name))

	if logo, ok := logosMap[nameF]; ok {
		return logo
	}

	return logosMap[baseLogo]
}
