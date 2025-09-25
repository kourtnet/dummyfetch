// Package entities provides entities used in the project: system info arguments and OS logos
package entities

type Arg struct {
	Name     string
	Command  func() (string, error)
	Contents string
}

const (
	OSName        = "os"
	KernelName    = "kernel"
	ShellName     = "shell"
	TerminalName  = "terminal"
	UptimeName    = "uptime"
	PaletteBgName = "palette_bg"
	PaletteFgName = "palette_fg"
	PrintName     = "print"
)

var ArgsMap = map[string]Arg{
	OSName: {
		Name:    "OS",
		Command: getPrettyDistro,
	},

	KernelName: {
		Name:    "Kernel",
		Command: getKernel,
	},

	ShellName: {
		Name:    "Shell",
		Command: getShell,
	},

	TerminalName: {
		Name:    "Terminal",
		Command: getTerminal,
	},

	UptimeName: {
		Name:    "Uptime",
		Command: getUptime,
	},

	PaletteBgName: {
		Command: getPaletteBg,
	},

	PaletteFgName: {
		Command: getPaletteFg,
	},

	PrintName: {
		Command: getPrint,
	},
}

var BasicArgs = []string{
	OSName,
	KernelName,
	TerminalName,
	UptimeName,
}
