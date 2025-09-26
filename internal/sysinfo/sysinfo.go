// Package sysinfo provides functions for gathering system and logo info
package sysinfo

import (
	"github.com/kourtnet/dummyfetch/internal/entities"
	"github.com/kourtnet/dummyfetch/internal/flags"
)

// TODO: rewrite this crap
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

func SetDefaultModules() {
	for _, str := range entities.DefaultArgs {
		flags.Config.Modules = append(flags.Config.Modules, flags.Module{Arg: str})
	}
}

func Fetch() error {
	if len(flags.Config.Modules) == 0 {
		SetDefaultModules()
	}

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
