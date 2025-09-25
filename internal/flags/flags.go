// Package flags provides parser for input flags and a global Config structure to get parse results
package flags

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/kourtnet/dummyfetch/internal/entities"
)

func Parse() error {
	flag.Func("module", "define a module to print in a form \"title\":\"module\" or just \"module\"", addModule)
	flag.Func("print", "define a string to print", addPrint)
	flag.Func("distro", "define distro logo to print. Use already predefined name or path to a text file with custom logo in it", setDistroLogo)

	flag.Parse()
	return nil
}

type Module struct {
	IsTitleSet bool
	Title      string
	Arg        string
}

var Config = struct {
	Modules  []Module
	LogoName string
}{}

func addModule(str string) error {
	const errStr = "unknown module name: %s"

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
