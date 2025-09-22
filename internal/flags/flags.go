// Package flags provides parser for input flags and a global Config structure to get parse results
package flags

import (
	"fmt"
	"os"
	"strings"

	"github.com/kourtnet/dummyfetch/internal/entities"
)

type Config struct {
	ArgsOrder []entities.Arg
	LogoName  string
}

var defaultArgs = []entities.Arg{
	entities.ArgsList["os"],
	entities.ArgsList["kernel"],
	entities.ArgsList["shell"],
	entities.ArgsList["uptime"],
}

func Parse() (Config, error) {
	cfg := Config{
		ArgsOrder: []entities.Arg{},
	}

	args := os.Args[1:]
	for _, v := range args {
		vF := strings.TrimLeft(v, "-")

		_, ok := entities.LogosMap[vF]
		if ok {
			cfg.LogoName = vF
			continue
		}

		arg, ok := entities.ArgsList[vF]
		if ok {
			cfg.ArgsOrder = append(cfg.ArgsOrder, arg)
			continue
		}

		return Config{}, fmt.Errorf("unknow flag: %s", v)
	}

	if len(cfg.ArgsOrder) == 0 {
		cfg.ArgsOrder = defaultArgs
	}

	return cfg, nil
}
