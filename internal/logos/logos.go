// Package logos provides ASCII OS logos of different sizes
// At the moment we have: Arch
package logos

import "strings"

type LogoInfo struct {
	Color int
	Logo  []string
}

var logosMap = map[string]LogoInfo{
	"arch": {
		Color: 34,
		Logo: []string{
			"   \033[34m/\\\033[0m\t",
			"  \033[34m/\\ \\\033[0m\t",
			" \033[34m/ .\\ \\\033[0m\t",
			"\033[34m/.'  '.\\\033[0m",
			"\t",
		},
	},

	"ubuntu": {
		Color: 31,
		Logo: []string{
			"  \033[33m/\033[31m-'-( )\033[0m",
			"\033[31m( )    \033[31m|\033[0m",
			" \033[33m\\     \033[31m/\033[0m",
			"   \033[31m-.-\033[33m( )\033[0m",
			"\t",
		},
	},

	"linux": {
		Color: 37,
		Logo: []string{
			"  \033[37m.-,\033[0m\t",
			"  \033[37moo\033[37m|\033[0m\t",
			" /\033[33mv \033[37m\\\033[0m\t",
			"\033[33m(\\\033[37m_^\033[33m/)\033[0m\t",
			"\t",
		},
	},
}

var colorsMap = map[string]int{
	"arch":   34,
	"ubuntu": 31,
	"linux":  37,
}

const baseLogo = "linux"

func GetLogo(name string) LogoInfo {
	nameF := strings.ToLower(strings.TrimSpace(name))

	if logo, ok := logosMap[nameF]; ok {
		return logo
	}

	return logosMap[baseLogo]
}
