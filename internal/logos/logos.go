// Package logos provides ASCII OS logos of different sizes
// At the moment we have: Arch
package logos

import (
	"strings"
)

var logosMap = map[string][]string{
	"arch": {
		`   /\   `,
		`  /\ \  `,
		` / .\ \ `,
		`/.'  '.\`,
		`        `,
	},

	"ubuntu": {
		`  /-'-( )`,
		`( )    | `,
		` \     / `,
		`   -.-( )`,
		`         `,
	},

	"linux": {
		`  .-, `,
		`  oo| `,
		` /--\ `,
		`(\_^/)`,
		`      `,
	},
}

const baseLogo = "linux"

func Get(name string) []string {
	nameF := strings.ToLower(strings.TrimSpace(name))

	if logo, ok := logosMap[nameF]; ok {
		return logo
	}

	return logosMap[baseLogo]
}
