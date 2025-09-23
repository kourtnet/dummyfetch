// Package sysinfo provides functions for gathering system and logo info
package sysinfo

import (
	"github.com/kourtnet/dummyfetch/internal/entities"
)

func fetchArg(argName string) error {
	arg := entities.ArgsMap[argName]
	if arg.Contents != "" {
		return nil
	}

	var err error
	arg.Contents, err = arg.Command()
	if err != nil {
		return err
	}

	entities.ArgsMap[argName] = arg

	return nil
}

func Fetch(argNames []string) error {
	for _, v := range argNames {
		err := fetchArg(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func FetchLogo(name string) (string, error) {
	if name == "" {
		var err error
		name, err = entities.GetDistro()
		if err != nil {
			return "", nil
		}
	}

	if _, ok := entities.LogosMap[name]; ok {
		return name, nil
	}

	return entities.BasicLogoName, nil
}
