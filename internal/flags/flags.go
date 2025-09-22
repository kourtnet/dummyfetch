// Package flags provides parser for input flags and a global Config structure to get parse results
package flags

import (
	"fmt"
	"os"

	"github.com/kourtnet/dummyfetch/internal/entities"
)

type Config struct {
	ArgsOrder []string
	Logo      string
}

func (c *Config) appendArg(flag string) bool {
	if arg, ok := argsMap[flag]; ok {
		c.ArgsOrder = append(c.ArgsOrder, arg)
		return true
	}

	return false
}

func (c *Config) setLogo(flag string) bool {
	if logo, ok := logosMap[flag]; ok {
		c.Logo = logo
		return true
	}

	return false
}

var argsMap = map[string]string{
	"--os":         entities.OSName,
	"--kernel":     entities.KernelName,
	"--shell":      entities.ShellName,
	"--term":       entities.TerminalName,
	"--uptime":     entities.UptimeName,
	"--palette_bg": entities.PaletteBgName,
	"--palette_fg": entities.PaletteFgName,
	"--indent":     entities.IndentName,
	"--sep":        entities.SeparatorName,
}

var logosMap = map[string]string{
	"--arch":   entities.ArchName,
	"--ubuntu": entities.UbuntuName,
	"--tux":    entities.TuxName,
}

func Parse() (Config, error) {
	cfg := Config{
		ArgsOrder: []string{},
	}

	flags := os.Args[1:]
	for _, flag := range flags {
		if ok := cfg.appendArg(flag); ok {
			continue
		}

		if ok := cfg.setLogo(flag); ok {
			continue
		}

		return Config{}, fmt.Errorf("unknown flag: %s", flag)
	}

	return cfg, nil
}
