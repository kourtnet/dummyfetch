// Package sysinfo provides functions for gathering system and logo info
package sysinfo

import (
	"github.com/kourtnet/dummyfetch/internal/entities"
	"github.com/kourtnet/dummyfetch/internal/flags"
)

func fetchArg(argName string) error {
	var err error

	arg := entities.ArgsMap[argName]

	if arg.Contents != "" {
		return nil
	}

	arg.Contents, err = arg.Command()
	if err != nil {
		return err
	}

	entities.ArgsMap[argName] = arg

	return nil
}

func Fetch() error {
	for _, v := range flags.Config.Modules {
		err := fetchArg(v.Arg)
		if err != nil {
			return err
		}
	}

	return nil
}

func FetchLogo() error {
	if flags.Config.LogoName != "" {
		return nil
	}

	distroName, err := entities.GetDistro()
	if err != nil {
		return err
	}

	if _, ok := entities.LogosMap[distroName]; ok {
		flags.Config.LogoName = distroName
	} else {
		flags.Config.LogoName = entities.BasicLogoName
	}

	return nil
}
