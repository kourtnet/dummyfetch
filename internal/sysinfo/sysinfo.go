// Package sysinfo provides functions for gathering system and logo info
package sysinfo

import (
	"bufio"
	"os"

	"github.com/kourtnet/dummyfetch/internal/entities"
	"github.com/kourtnet/dummyfetch/internal/flags"
)

func FetchArg(argName string) (string, error) {
	arg, ok := entities.ArgsMap[argName]
	if !ok {
		return argName, nil
	}

	if arg.Contents != "" {
		return arg.Contents, nil
	}

	var err error
	arg.Contents, err = arg.Command()
	if err != nil {
		return "", err
	}

	entities.ArgsMap[argName] = arg

	return arg.Contents, nil
}

func fetchAutoLogo() error {
	distroName, err := entities.GetDistro()
	if err != nil {
		return err
	}

	if _, ok := entities.LogosMap[distroName]; ok {
		flags.Config.LogoName = distroName
	} else {
		flags.Config.LogoName = entities.DefaultLogoName
	}

	return nil
}

func fetchFileLogo() error {
	file, err := os.Open(flags.Config.LogoName)
	if err != nil {
		return err
	}

	defer file.Close()

	ascii := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ascii = append(ascii, string(scanner.Text()))
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	customLogo := entities.LogosMap[entities.CustomName]
	customLogo.Logo = ascii

	entities.LogosMap[entities.CustomName] = customLogo
	flags.Config.LogoName = entities.CustomName

	return nil
}

func FetchLogo() error {
	logoName := flags.Config.LogoName
	if _, ok := entities.LogosMap[logoName]; ok && logoName != entities.CustomName {
		return nil
	}

	if logoName == entities.AutoName {
		if err := fetchAutoLogo(); err != nil {
			return err
		}

		return nil
	}

	if err := fetchFileLogo(); err != nil {
		return err
	}

	return nil
}
