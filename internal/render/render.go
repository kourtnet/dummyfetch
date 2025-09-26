// Package render provides tools for rendering fetch information
package render

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/kourtnet/dummyfetch/internal/entities"
	"github.com/kourtnet/dummyfetch/internal/flags"
	"github.com/kourtnet/dummyfetch/internal/sysinfo"
)

const (
	seqStart = "\033["
	seqReset = "\033[0m"
)

var styleVars = map[string]string{
	// Foreground colors
	"fgdbla": "\033[30m", // black
	"fgdr":   "\033[31m", // red
	"fgdg":   "\033[32m", // green
	"fgdy":   "\033[33m", // yellow
	"fgdblu": "\033[34m", // blue
	"fgdm":   "\033[35m", // magenta
	"fgdc":   "\033[36m", // cyan
	"fgdw":   "\033[37m", // white
	"fglbla": "\033[90m", // black
	"fglr":   "\033[91m", // red
	"fglg":   "\033[92m", // green
	"fgly":   "\033[93m", // yellow
	"fglblu": "\033[94m", // blue
	"fglm":   "\033[95m", // magenta
	"fglc":   "\033[96m", // cyan
	"fglw":   "\033[97m", // white

	// Background colors
	"bgdbla": "\033[40m",  // black
	"bgdr":   "\033[41m",  // red
	"bgdg":   "\033[42m",  // green
	"bgdy":   "\033[43m",  // yellow
	"bgdblu": "\033[44m",  // blue
	"bgdm":   "\033[45m",  // magenta
	"bgdc":   "\033[46m",  // cyan
	"bgdw":   "\033[107m", // white
	"bglbla": "\033[100m", // black
	"bglr":   "\033[101m", // red
	"bglg":   "\033[102m", // green
	"bgly":   "\033[103m", // yellow
	"bglblu": "\033[104m", // blue
	"bglm":   "\033[105m", // magenta
	"bglc":   "\033[106m", // cyan
	"bglw":   "\033[107m", // white

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
	logo := entities.LogosMap[entities.CustomName]

	if len(entities.LogosMap[entities.CustomName].Logo) == 0 {
		logo.BlankRow = seqStart + "0" + "C"
	} else {
		clearRow := styleVarRegex.ReplaceAllString(logo.Logo[0], "")
		length := len([]rune(clearRow))

		logo.BlankRow = seqStart + strconv.Itoa(length) + "C"
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

		textColor := entities.LogosMap[flags.Config.LogoName].TextColor
		flags.Config.Modules[i].Title = "\033[" + textColor + "m" + flags.Config.Modules[i].Title

		flags.Config.Modules[i].Arg = strTmp
	}

	return nil
}

func prepare() error {
	if err := prepareLogo(); err != nil {
		return err
	}

	if err := prepareModules(); err != nil {
		return err
	}

	return nil
}

func render() {
	logoJoined := strings.Join(entities.LogosMap[flags.Config.LogoName].Logo, "\n")
	os.Stdout.WriteString(logoJoined + seqReset)

	logoHeight := len(entities.LogosMap[flags.Config.LogoName].Logo)
	modsNum := len(flags.Config.Modules)

	if modsNum > logoHeight {
		for range modsNum - logoHeight {
			os.Stdout.WriteString("\n")
		}
		os.Stdout.WriteString(seqStart + strconv.Itoa(modsNum-1) + "A")
	} else {
		os.Stdout.WriteString(seqStart + strconv.Itoa(logoHeight-1) + "A")
	}

	os.Stdout.WriteString(entities.LogosMap[flags.Config.LogoName].BlankRow)

	for _, module := range flags.Config.Modules {
		os.Stdout.WriteString(flags.Config.LogoSeparator + seqReset)

		os.Stdout.WriteString(module.Title + seqReset)

		os.Stdout.WriteString(flags.Config.ModuleSeparator + seqReset)

		os.Stdout.WriteString(module.Arg + seqReset)

		os.Stdout.WriteString(seqStart + "1B\r" + entities.LogosMap[flags.Config.LogoName].BlankRow)
	}

	modulesNum := len(flags.Config.Modules)
	if logoHeight > modulesNum {
		os.Stdout.WriteString(seqStart + strconv.Itoa(logoHeight-modulesNum) + "B")
	}

	os.Stdout.WriteString("\n\n")
}

func PrepareAndRender() {
	if err := prepare(); err != nil {
		os.Stdout.WriteString(err.Error() + "\n")
		return
	}

	render()
}
