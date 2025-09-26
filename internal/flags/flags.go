// Package flags provides parser for input flags and a global Config structure to get parse results
package flags

import (
	"flag"
	"strconv"
	"strings"

	"github.com/kourtnet/dummyfetch/internal/entities"
)

func Parse() {
	flag.Func("module", "define a module to print in a form \"title\":\"module\" or just \"module\"", addModule)
	flag.Func("mod", "-module alias", addModule)

	flag.Func("variable", "define a style variable to use in modules formatting. Predefined variables do not not work here, so as user variables", addVar)
	flag.Func("var", "-variable alias", addVar)

	flag.StringVar(&Config.LogoName, "distro", entities.AutoName, "define distro logo to print. Use already predefined name or path to a text file with custom logo in it")

	flag.Func("logo_separator", "define separator between logo and modules", setLogoSeparator)
	flag.Func("logo_sep", "-separator alias", setLogoSeparator)

	flag.Func("module_separator", "define separator between module title and contents", setModuleSeparator)
	flag.Func("mod_sep", "-module_separator alias", setModuleSeparator)

	flag.Parse()

	if len(Config.Modules) == 0 {
		Config.Modules = defaultArgs
	}
}

var defaultArgs = []Module{
	{Arg: "${" + entities.OSName + "}"},
	{Arg: "${" + entities.KernelName + "}"},
	{Arg: "${" + entities.TerminalName + "}"},
	{Arg: "${" + entities.UptimeName + "}"},
}

type Module struct {
	IsTitleSet bool
	Title      string
	Arg        string
}

var Config = struct {
	Modules         []Module
	LogoName        string
	Vars            map[string]string
	LogoSeparator   string
	ModuleSeparator string
}{
	Vars:            map[string]string{},
	LogoSeparator:   entities.DefaultLogoSeparator,
	ModuleSeparator: entities.DefaultModuleSeparator,
}

func setLogoSeparator(str string) error {
	if Config.LogoSeparator == entities.DefaultLogoSeparator {
		Config.LogoSeparator = str
	}

	return nil
}

func setModuleSeparator(str string) error {
	if Config.ModuleSeparator == entities.DefaultModuleSeparator {
		Config.ModuleSeparator = str
	}

	return nil
}

func unescapeStr(str string) (string, error) {
	quoted := `"` + str + `"`

	unquoted, err := strconv.Unquote(quoted)
	if err != nil {
		return "", err
	}

	return unquoted, nil
}

func addVar(str string) error {
	str, err := unescapeStr(str)
	if err != nil {
		return err
	}

	Config.Vars[strconv.Itoa(len(Config.Vars))] = str
	return nil
}

func addModule(str string) error {
	str, err := unescapeStr(str)
	if err != nil {
		return err
	}

	colonIndex := strings.LastIndex(str, ":")

	var title, arg string
	var isTitleSet bool

	if colonIndex != -1 {
		title = str[:colonIndex]
		arg = str[colonIndex+1:]
		isTitleSet = true
	} else {
		arg = str
	}

	Config.Modules = append(Config.Modules, Module{IsTitleSet: isTitleSet, Title: title, Arg: arg})
	return nil
}
