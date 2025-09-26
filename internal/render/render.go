// Package render provides tools for rendering fetch information
package render

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/kourtnet/dummyfetch/internal/entities"
	"github.com/kourtnet/dummyfetch/internal/flags"
	"github.com/kourtnet/dummyfetch/internal/sysinfo"
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

		// predefined style vars
		if ansi, ok := styleVars[varName]; ok {
			return ansi
		}

		// user defined vars
		if ansi, ok := flags.Config.Vars[varName]; ok {
			return ansi
		}

		return match
	})

	return res
}

func countLogoLength() {
	const format = "\x1b[%dC"

	logo := entities.LogosMap[entities.CustomName]

	if len(entities.LogosMap[entities.CustomName].Logo) == 0 {
		logo.BlankRow = fmt.Sprintf(format, 0)
	} else {
		clearRow := styleVarRegex.ReplaceAllString(logo.Logo[0], "")
		length := len([]rune(clearRow))

		logo.BlankRow = fmt.Sprintf(format, length)
	}

	entities.LogosMap[entities.CustomName] = logo
}

func prepareLogo() error {
	if err := sysinfo.FetchLogo(); err != nil {
		return err
	}

	if flags.Config.LogoName != entities.CustomName {
		return nil
	}

	countLogoLength()

	for i, str := range entities.LogosMap[entities.CustomName].Logo {
		entities.LogosMap[entities.CustomName].Logo[i] = prepareStr(str)
	}

	return nil
}

func prepareArg(str string) (string, string, error) {
	str = prepareStr(str)

	var title string
	var err error

	res := styleVarRegex.ReplaceAllStringFunc(str, func(match string) string {
		argName := strings.TrimPrefix(strings.TrimSuffix(match, "}"), "${")

		var argContents string
		argContents, err = sysinfo.FetchArg(argName)
		if err != nil {
			return ""
		}

		if title == "" {
			title = entities.ArgsMap[argName].DefaultTitle
		}

		return argContents
	})

	if err != nil {
		return "", "", err
	}

	return res, title, nil
}

func prepareModules() error {
	for i, mod := range flags.Config.Modules {
		flags.Config.Modules[i].Title = prepareStr(mod.Title)

		strTmp, title, err := prepareArg(mod.Arg)
		if err != nil {
			return err
		}

		if flags.Config.Modules[i].Title == "" {
			flags.Config.Modules[i].Title = title
		}

		flags.Config.Modules[i].Arg = strTmp
	}

	return nil
}

func prepare() error {
	if err := prepareModules(); err != nil {
		return err
	}

	if err := prepareLogo(); err != nil {
		return err
	}

	return nil
}

func renderLogo() {
	logoJoined := strings.Join(entities.LogosMap[flags.Config.LogoName].Logo, "\n")
	os.Stdout.WriteString(logoJoined)
}

func render() {
	renderLogo()

	logoHeight := len(entities.LogosMap[flags.Config.LogoName].Logo)
	os.Stdout.WriteString("\033[0m\033[" + strconv.Itoa(logoHeight-1) + "A")

	textColor := entities.LogosMap[flags.Config.LogoName].TextColor
	for _, module := range flags.Config.Modules {
		// TODO: ADD MODULE SEPARATOR VAR AND FLAG
		os.Stdout.WriteString(flags.Config.Separator)
		os.Stdout.WriteString("\033[" + strconv.Itoa(textColor) + "m" + module.Title + "\033[0m")
		os.Stdout.WriteString(": " + module.Arg)

		os.Stdout.WriteString("\033[1B\r" + entities.LogosMap[flags.Config.LogoName].BlankRow)
	}

	modulesNum := len(flags.Config.Modules)
	if logoHeight > modulesNum {
		os.Stdout.WriteString("\033[" + strconv.Itoa(logoHeight-modulesNum) + "B")
	}
	os.Stdout.WriteString("\n")
}

func PrepareAndRender() {
	if err := prepare(); err != nil {
		fmt.Println(err)
		return
	}

	render()
}
