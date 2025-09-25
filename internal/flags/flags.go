// Package flags provides parser for input flags and a global Config structure to get parse results
package flags

import (
	"flag"

	"github.com/kourtnet/dummyfetch/internal/entities"
)

type Config struct {
	modules Modules
}

var argsMap = map[string]struct{}{
	entities.OSName:        {},
	entities.KernelName:    {},
	entities.ShellName:     {},
	entities.TerminalName:  {},
	entities.UptimeName:    {},
	entities.PaletteBgName: {},
	entities.PaletteFgName: {},
	entities.PrintName:     {},
}

var logosMap = map[string]string{
	"--arch":   entities.ArchName,
	"--ubuntu": entities.UbuntuName,
	"--tux":    entities.TuxName,
}

func Parse() (Config, error) {
	cfg := Config{}

	flag.Var(&cfg.modules, "module", "define a module to print in a form \"title\":\"module\" or just \"module\"")
	flag.Var(&Print{&cfg.modules}, "print", "define a string to print")

	flag.Parse()
	return cfg, nil
}
