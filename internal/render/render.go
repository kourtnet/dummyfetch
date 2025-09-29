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
	// dark
	"fgdbla": "\033[30m", // black
	"fgdr":   "\033[31m", // red
	"fgdg":   "\033[32m", // green
	"fgdy":   "\033[33m", // yellow
	"fgdblu": "\033[34m", // blue
	"fgdm":   "\033[35m", // magenta
	"fgdc":   "\033[36m", // cyan
	"fgdw":   "\033[37m", // white
	// light
	"fglbla": "\033[90m", // black
	"fglr":   "\033[91m", // red
	"fglg":   "\033[92m", // green
	"fgly":   "\033[93m", // yellow
	"fglblu": "\033[94m", // blue
	"fglm":   "\033[95m", // magenta
	"fglc":   "\033[96m", // cyan
	"fglw":   "\033[97m", // white

	// Background colors
	// dark
	"bgdbla": "\033[40m",  // black
	"bgdr":   "\033[41m",  // red
	"bgdg":   "\033[42m",  // green
	"bgdy":   "\033[43m",  // yellow
	"bgdblu": "\033[44m",  // blue
	"bgdm":   "\033[45m",  // magenta
	"bgdc":   "\033[46m",  // cyan
	"bgdw":   "\033[107m", // white
	// light
	"bglbla": "\033[100m", // black
	"bglr":   "\033[101m", // red
	"bglg":   "\033[102m", // green
	"bgly":   "\033[103m", // yellow
	"bglblu": "\033[104m", // blue
	"bglm":   "\033[105m", // magenta
	"bglc":   "\033[106m", // cyan
	"bglw":   "\033[107m", // white

	// Text styles
	"b":  "\033[1m",  // bold
	"d":  "\033[2m",  // dim
	"i":  "\033[3m",  // italic
	"u":  "\033[4m",  // underline
	"du": "\033[21m", // double underline
	"r":  "\033[7m",  // reverse
	"s":  "\033[9m",  // strikethrough

	// Reset all styles
	"reset": "\033[0m",
}

var moveVars = map[string]string{
	"up":    "A",
	"down":  "B",
	"left":  "D",
	"right": "C",
}

var (
	ansiEscapeRegex = regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	styleVarRegex   = regexp.MustCompile(`\$\{([^}]+)\}`)
	moveVarRegex    = regexp.MustCompile(`^(up|down|left|right)(\d+)$`)
	moveEscapeRegex = regexp.MustCompile(`\x1b\[[0-9;]*[ABCDG]`)
)

func prepare() error {
	prepareUserVars()

	if err := prepareLogo(); err != nil {
		return err
	}

	prepareSeparators()

	if err := prepareModules(); err != nil {
		return err
	}

	return nil
}

func prepareUserVars() {
	for k, v := range flags.Config.Vars {
		flags.Config.Vars[k] = resolvePredefinedVars(v)
	}
}

func resolvePredefinedVars(str string) string {
	res := styleVarRegex.ReplaceAllStringFunc(str, func(match string) string {
		varName := strings.TrimPrefix(strings.TrimSuffix(match, "}"), "${")

		// style vars
		if ansi, ok := styleVars[varName]; ok {
			return ansi
		}

		// move vars
		if moveMatch := moveVarRegex.FindStringSubmatch(str); moveMatch != nil {
			direction, offset := moveMatch[1], moveMatch[2]

			if directionLetter, ok := moveVars[direction]; ok {
				return seqStart + offset + directionLetter
			}
		}

		return match
	})

	return res
}

func removeOffsets(str string) string {
	return moveEscapeRegex.ReplaceAllString(str, "")
}

// TODO: forbid move vars in logo
func prepareLogo() error {
	if err := sysinfo.FetchLogo(); err != nil {
		return err
	}

	// there's no need to prepare predefined logos
	if flags.Config.LogoName != entities.CustomName {
		return nil
	}

	for i, str := range entities.LogosMap[entities.CustomName].Logo {
		entities.LogosMap[entities.CustomName].Logo[i] = resolvePredefinedVars(str)
		entities.LogosMap[entities.CustomName].Logo[i] = resolveUserVars(str)
		entities.LogosMap[entities.CustomName].Logo[i] = removeOffsets(str)
	}

	setLogoOffset()

	return nil
}

func visibleRowLen(row string) int {
	cleanRow := ansiEscapeRegex.ReplaceAllString(row, "")
	return len(cleanRow)
}

func setLogoOffset() {
	logo := entities.LogosMap[entities.CustomName]

	maxLen := 0
	for _, row := range logo.Logo {
		maxLen = max(maxLen, visibleRowLen(row))
	}

	logo.BlankRow = seqStart + strconv.Itoa(maxLen) + "C"

	entities.LogosMap[entities.CustomName] = logo
}

func resolveUserVars(str string) string {
	res := styleVarRegex.ReplaceAllStringFunc(str, func(match string) string {
		varName := strings.TrimPrefix(strings.TrimSuffix(match, "}"), "${")

		if ansi, ok := flags.Config.Vars[varName]; ok {
			return ansi
		}

		return match
	})

	return res
}

// TODO: remake prepare arg and module funcs
func prepareArg(str string) (string, string, error) {
	str = resolvePredefinedVars(str)
	str = resolveUserVars(str)

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
		flags.Config.Modules[i].Title = resolvePredefinedVars(mod.Title)
		flags.Config.Modules[i].Title = resolveUserVars(mod.Title)

		strTmp, title, err := prepareArg(mod.Arg)
		if err != nil {
			return err
		}

		if flags.Config.Modules[i].Title == "" {
			flags.Config.Modules[i].Title = title
		}

		textColor := entities.LogosMap[flags.Config.LogoName].TextColor
		flags.Config.Modules[i].Title = seqStart + textColor + "m" + flags.Config.Modules[i].Title

		flags.Config.Modules[i].Arg = strTmp
	}

	return nil
}

func prepareSeparators() {
	flags.Config.LogoSeparator = resolvePredefinedVars(flags.Config.LogoSeparator)
	flags.Config.LogoSeparator = resolveUserVars(flags.Config.LogoSeparator)

	flags.Config.ModuleSeparator = resolvePredefinedVars(flags.Config.ModuleSeparator)
	flags.Config.ModuleSeparator = resolveUserVars(flags.Config.ModuleSeparator)
}

func renderModule(module flags.Module) {
	os.Stdout.WriteString(seqStart + "1B\r" + entities.LogosMap[flags.Config.LogoName].BlankRow)

	os.Stdout.WriteString(flags.Config.LogoSeparator + seqReset)

	os.Stdout.WriteString(module.Title + seqReset)

	os.Stdout.WriteString(flags.Config.ModuleSeparator + seqReset)

	os.Stdout.WriteString(module.Arg + seqReset)
}

func renderLogo() {
	logoJoined := strings.Join(entities.LogosMap[flags.Config.LogoName].Logo, "\n")
	os.Stdout.WriteString(logoJoined + seqReset)
}

func render() {
	renderLogo()

	// TODO: DELETE ALL THIS CRAP AFTER A PROPER BUFFER FILLER IS DONE
	logoHeight := len(entities.LogosMap[flags.Config.LogoName].Logo)
	os.Stdout.WriteString(seqStart + strconv.Itoa(logoHeight) + "A")

	for _, module := range flags.Config.Modules {
		renderModule(module)
	}

	modulesNum := len(flags.Config.Modules)
	offset := logoHeight - modulesNum + len(flags.Config.LogoSeparator)
	os.Stdout.WriteString(seqStart + strconv.Itoa(offset) + "B")

	os.Stdout.WriteString("\n\n")
}

func PrepareAndRender() {
	if err := prepare(); err != nil {
		os.Stdout.WriteString(err.Error() + "\n")
		return
	}

	render()
}
