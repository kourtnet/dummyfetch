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
}

var defaultArgs = []string{
	"os",
	"kernel",
	"shell",
	"uptime",
	"palette_bg",
	"palette_fg",
}

func Parse() (Config, error) {
	cfg := Config{
		[]entities.Arg{},
	}

	args := os.Args[1:]
	if len(os.Args[1:]) == 0 {
		args = defaultArgs
	}

	for _, v := range args {
		vF := strings.TrimLeft(v, "-")

		arg, ok := entities.ArgsList[vF]
		if !ok {
			return Config{}, fmt.Errorf("unknow flag: %s", v)
		}

		cfg.ArgsOrder = append(cfg.ArgsOrder, arg)
	}

	return cfg, nil
}
