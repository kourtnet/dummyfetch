// Package entities provides entities used in the project: system info arguments and OS logos
package entities

type Arg struct {
	Name     string
	Command  func() (string, error)
	Contents string
}

var ArgsList = map[string]Arg{
	"os": {
		Name:    "OS",
		Command: getOS,
	},

	"kernel": {
		Name:    "Kernel",
		Command: getKernel,
	},

	"shell": {
		Name:    "Shell",
		Command: getShell,
	},

	"terminal": {
		Name:    "Terminal",
		Command: getTerminal,
	},

	"uptime": {
		Name:    "Uptime",
		Command: getUptime,
	},

	"palette_bg": {
		Command: getPaletteBg,
	},

	"palette_fg": {
		Command: getPaletteFg,
	},

	"indent": {
		Command: getIndent,
	},

	"separator": {
		Command: getSeparator,
	},
}
