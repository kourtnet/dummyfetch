// Package render provides tools for rendering fetch information
package render

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/kourtnet/dummyfetch/internal/entities"
	"github.com/kourtnet/dummyfetch/internal/flags"
)

var styleVars = map[string]string{
	// Background colors
	"bg0": "\033[30m", // black
	"bg1": "\033[31m", // red
	"bg2": "\033[32m", // green
	"bg3": "\033[33m", // yellow
	"bg4": "\033[34m", // blue
	"bg5": "\033[35m", // magenta
	"bg6": "\033[36m", // cyan
	"bg7": "\033[37m", // white

	// Foreground colors
	"fg0": "\033[90m", // black
	"fg1": "\033[91m", // red
	"fg2": "\033[92m", // green
	"fg3": "\033[93m", // yellow
	"fg4": "\033[94m", // blue
	"fg5": "\033[95m", // magenta
	"fg6": "\033[96m", // cyan
	"fg7": "\033[97m", // white

	// Text styles
	"b": "\033[1m", // bold
	"d": "\033[2m", // dim
	"i": "\033[3m", // italic
	"u": "\033[4m", // underline
	"r": "\033[7m", // reverse
	"s": "\033[9m", // strikethrough

	// Reset all styles
	"reset": "\033[0m",
}

var styleVarRegex = regexp.MustCompile(`\$\{([^}]+)\}`)

func prepareStr(str string) string {
	res := styleVarRegex.ReplaceAllStringFunc(str, func(match string) string {
		varName := strings.TrimPrefix(strings.TrimSuffix(match, "}"), "${")
		if ansi, ok := styleVars[varName]; ok {
			return ansi
		}

		return match
	})

	return res
}

func prepareLogo() {
	if flags.Config.LogoName != entities.CustomName {
		return
	}

	for i, str := range entities.LogosMap[entities.CustomName].Logo {
		entities.LogosMap[entities.CustomName].Logo[i] = prepareStr(str)
	}
}

func prepareModules() {
	for i, mod := range flags.Config.Modules {
		flags.Config.Modules[i].Title = prepareStr(mod.Title)
	}
}

func prepare() {
	prepareModules()
	prepareLogo()
}

func Render() {
	prepare()

	logo := entities.LogosMap[flags.Config.LogoName]

	maxLen := max(len(logo.Logo), len(flags.Config.Modules))

	for i := range maxLen {
		if i < len(logo.Logo) {
			fmt.Print(logo.Logo[i])
		} else {
			fmt.Print(logo.BlankRow)
		}

		fmt.Print("  ")

		if i < len(flags.Config.Modules) {
			arg := entities.ArgsMap[flags.Config.Modules[i].Arg]

			titleToPrint := arg.DefaultTitle
			if flags.Config.Modules[i].IsTitleSet {
				titleToPrint = flags.Config.Modules[i].Title
			}

			fmt.Printf("\033[%dm%s\033[0m", logo.TextColor, titleToPrint)
			fmt.Printf(": %s\n", arg.Contents)
		} else {
			fmt.Println()
		}
	}
}
