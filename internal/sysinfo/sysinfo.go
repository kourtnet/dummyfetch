// Package sysinfo provides function for gathering system info
package sysinfo

import (
	"github.com/kourtnet/dummyfetch/internal/entities"
)

func Fetch(args []entities.Arg) ([]entities.Arg, error) {
	for i := range args {
		output, err := args[i].Command()
		if err != nil {
			return nil, err
		}

		args[i].Contents = output
	}

	return args, nil
}

func ResolveIcon() (entities.LogoInfo, error) {
	logo, err := entities.GetLogo()
	if err != nil {
		return entities.LogoInfo{}, err
	}

	return logo, err
}
