// Package flags provides parser for input flags and a global Config structure to get parse results
package flags

import (
	"flag"

	"github.com/kourtnet/dummyfetch/internal/entities"
)

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

func Parse() error {
	flag.Func("module", "define a module to print in a form \"title\":\"module\" or just \"module\"", addModule)
	flag.Func("print", "define a string to print", addPrint)
	flag.Func("distro", "define distro logo to print. Use already predefined name or path to a text file with custom logo in it", setDistroLogo)

	flag.Parse()
	return nil
}
