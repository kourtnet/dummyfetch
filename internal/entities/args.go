// Package entities provides entities used in the project: system info arguments and OS logos
package entities

type Arg struct {
	BasicTitle string
	Command    func() (string, error)
	Contents   string
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
		BasicTitle: "OS",
		Command:    getPrettyDistro,
	},

	KernelName: {
		BasicTitle: "Kernel",
		Command:    getKernel,
	},

	ShellName: {
		BasicTitle: "Shell",
		Command:    getShell,
	},

	TerminalName: {
		BasicTitle: "Terminal",
		Command:    getTerminal,
	},

	UptimeName: {
		BasicTitle: "Uptime",
		Command:    getUptime,
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
