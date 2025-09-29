package render

import (
	"strconv"
	"strings"

	"github.com/kourtnet/dummyfetch/internal/entities"
	"github.com/kourtnet/dummyfetch/internal/flags"
	"github.com/kourtnet/dummyfetch/internal/sysinfo"
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

func prepareModules() error {
	for i, mod := range flags.Config.Modules {
		mod.Title = resolvePredefinedVars(mod.Title)
		mod.Title = resolveUserVars(mod.Title)

		mod.Arg = resolvePredefinedVars(mod.Arg)
		mod.Arg = resolveUserVars(mod.Arg)

		arg, title, err := resolveModVars(mod.Arg)
		if err != nil {
			return err
		}

		mod.Arg = arg

		if mod.Title == "" {
			mod.Title = title
			textColor := entities.LogosMap[flags.Config.LogoName].TextColor
			mod.Title = seqStart + textColor + "m" + mod.Title
		}

		flags.Config.Modules[i] = mod
	}

	return nil
}

func resolveModVars(str string) (string, string, error) {
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

func prepareSeparators() {
	flags.Config.LogoSeparator = resolvePredefinedVars(flags.Config.LogoSeparator)
	flags.Config.LogoSeparator = resolveUserVars(flags.Config.LogoSeparator)

	flags.Config.ModuleSeparator = resolvePredefinedVars(flags.Config.ModuleSeparator)
	flags.Config.ModuleSeparator = resolveUserVars(flags.Config.ModuleSeparator)
}

func resolvePredefinedVars(str string) string {
	res := styleVarRegex.ReplaceAllStringFunc(str, func(match string) string {
		varName := strings.TrimPrefix(strings.TrimSuffix(match, "}"), "${")

		// style vars
		if ansi, ok := styleVars[varName]; ok {
			return ansi
		}

		// move vars
		if moveMatch := moveVarRegex.FindStringSubmatch(varName); moveMatch != nil {
			direction, offset := moveMatch[1], moveMatch[2]

			if directionLetter, ok := moveVars[direction]; ok {
				return seqStart + offset + directionLetter
			}
		}

		return match
	})

	return res
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

func removeOffsets(str string) string {
	return moveEscapeRegex.ReplaceAllString(str, "")
}

func visibleRowLen(row string) int {
	cleanRow := escapeRegex.ReplaceAllString(row, "")
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
