package render

import "regexp"

var styleVars = map[string]string{
	// Foreground colors
	// dark
	"fgdbla": "\033[30m", // black
	"fgdr":   "\033[31m", // red
	"fgdg":   "\033[32m", // green
	"fgdy":   "\033[33m", // yellow
	"fgdblu": "\033[34m", // blue
	"fgdm":   "\033[35m", // magenta
	"fgdc":   "\033[36m", // cyan
	"fgdw":   "\033[37m", // white
	// light
	"fglbla": "\033[90m", // black
	"fglr":   "\033[91m", // red
	"fglg":   "\033[92m", // green
	"fgly":   "\033[93m", // yellow
	"fglblu": "\033[94m", // blue
	"fglm":   "\033[95m", // magenta
	"fglc":   "\033[96m", // cyan
	"fglw":   "\033[97m", // white

	// Background colors
	// dark
	"bgdbla": "\033[40m",  // black
	"bgdr":   "\033[41m",  // red
	"bgdg":   "\033[42m",  // green
	"bgdy":   "\033[43m",  // yellow
	"bgdblu": "\033[44m",  // blue
	"bgdm":   "\033[45m",  // magenta
	"bgdc":   "\033[46m",  // cyan
	"bgdw":   "\033[107m", // white
	// light
	"bglbla": "\033[100m", // black
	"bglr":   "\033[101m", // red
	"bglg":   "\033[102m", // green
	"bgly":   "\033[103m", // yellow
	"bglblu": "\033[104m", // blue
	"bglm":   "\033[105m", // magenta
	"bglc":   "\033[106m", // cyan
	"bglw":   "\033[107m", // white

	// Text styles
	"b":  "\033[1m",  // bold
	"d":  "\033[2m",  // dim
	"i":  "\033[3m",  // italic
	"u":  "\033[4m",  // underline
	"du": "\033[21m", // double underline
	"r":  "\033[7m",  // reverse
	"s":  "\033[9m",  // strikethrough

	// Reset all styles
	"reset": "\033[0m",
}

var moveVars = map[string]string{
	"up":    "A",
	"down":  "B",
	"left":  "D",
	"right": "C",
}

var (
	// regexes for escape sequences
	escapeRegex     = regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	moveEscapeRegex = regexp.MustCompile(`\x1b\[[0-9;]*[ABCDG]`)

	// regexes for dummyfetch variables
	styleVarRegex = regexp.MustCompile(`\$\{([^}]+)\}`)
	moveVarRegex  = regexp.MustCompile(`^(up|down|left|right)(\d+)$`)

	// regex to count offset of modules considering cursor offset escape sequences
	offsetCountingRegex = regexp.MustCompile(`\x1b\[(\d+)([AB])|([^\x1b]+)|(\x1b\[[0-9;]*m)`)
)

const (
	seqStart = "\033["
	seqReset = "\033[0m"
)
