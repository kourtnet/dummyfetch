// Package entities provides entities used in the project: system info arguments and OS logos
package entities

type Arg struct {
	DefaultTitle string
	Command      func() (string, error)
	Contents     string
}

const (
	OSName        = "os"
	KernelName    = "kernel"
	ShellName     = "shell"
	TerminalName  = "terminal"
	UptimeName    = "uptime"
	PaletteBgName = "palette_bg"
	PaletteFgName = "palette_fg"

	DefaultLogoSeparator   = "  "
	DefaultModuleSeparator = ": "
)

var ArgsMap = map[string]Arg{
	OSName: {
		DefaultTitle: "OS",
		Command:      getPrettyDistro,
	},

	KernelName: {
		DefaultTitle: "Kernel",
		Command:      getKernel,
	},

	ShellName: {
		DefaultTitle: "Shell",
		Command:      getShell,
	},

	TerminalName: {
		DefaultTitle: "Terminal",
		Command:      getTerminal,
	},

	UptimeName: {
		DefaultTitle: "Uptime",
		Command:      getUptime,
	},

	PaletteBgName: {
		Command: getPaletteBg,
	},

	PaletteFgName: {
		Command: getPaletteFg,
	},
}
