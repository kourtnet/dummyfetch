// Package flags provides parser for input flags and a global Config structure to get parse results
package flags

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kourtnet/dummyfetch/internal/entities"
)

func Parse() {
	flag.Func("module", "define a module to print in a form \"title\":\"module\" or just \"module\"", addModule)
	flag.Func("print", "define a string to print", addPrint)

	flag.Func("distro", "define distro logo to print. Use already predefined name or path to a text file with custom logo in it", setDistroLogo)

	flag.Func("variable", "define a style variable to use in modules formatting. Predefined variables do not not work here, so as user variables", addVar)
	flag.Func("var", "-variable alias", addVar)

	flag.Parse()
}

type Module struct {
	IsTitleSet bool
	Title      string
	Arg        string
}

var Config = struct {
	Modules  []Module
	LogoName string
	Vars     map[string]string
}{
	Vars: map[string]string{},
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
	const errStr = "unknown module name: %s"

	str, err := unescapeStr(str)
	if err != nil {
		return err
	}

	colonIndex := strings.LastIndex(str, ":")

	if colonIndex != -1 {
		title := str[:colonIndex]
		arg := str[colonIndex+1:]

		if _, ok := entities.ArgsMap[arg]; !ok {
			return fmt.Errorf(errStr, arg)
		}

		Config.Modules = append(Config.Modules, Module{IsTitleSet: true, Title: title, Arg: arg})
		return nil
	}

	if _, ok := entities.ArgsMap[str]; !ok {
		return fmt.Errorf(errStr, str)
	}

	Config.Modules = append(Config.Modules, Module{Arg: str})
	return nil
}

func addPrint(str string) error {
	return addModule(str + ":print")
}

func setDistroLogo(str string) error {
	if _, ok := entities.LogosMap[str]; ok && str != entities.CustomName {
		Config.LogoName = str
		return nil
	}

	file, err := os.Open(str)
	if err != nil {
		return fmt.Errorf("unknown distro name or invalid filepath: %s", str)
	}

	defer file.Close()

	ascii := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ascii = append(ascii, string(scanner.Text()))
	}

	if scanner.Err() != nil {
		return fmt.Errorf("failed to read file with custom logo")
	}

	customLogo := entities.LogosMap[entities.CustomName]
	customLogo.Logo = ascii

	entities.LogosMap[entities.CustomName] = customLogo
	Config.LogoName = entities.CustomName

	return nil
}
