// Package render provides tools for rendering fetch information
package render

import (
	"os"
	"strconv"
	"strings"

	"github.com/kourtnet/dummyfetch/internal/entities"
	"github.com/kourtnet/dummyfetch/internal/flags"
)

func render() {
	modsOffset := countModsOffset()
	logoOffset := countLogoOffset()
	maxOffset := max(logoOffset, modsOffset)

	fillTerminal(maxOffset)
	cursorUp(maxOffset)

	renderLogo()
	cursorUp(logoOffset)

	for _, module := range flags.Config.Modules {
		renderModule(module)
	}
	cursorUp(modsOffset)

	cursorDown(maxOffset)
	os.Stdout.WriteString("\n")
}

func fillTerminal(offset int) {
	os.Stdout.WriteString(strings.Repeat("\n", offset))
}

func cursorUp(offset int) {
	os.Stdout.WriteString(seqStart + strconv.Itoa(offset) + "A")
}

func cursorDown(offset int) {
	os.Stdout.WriteString(seqStart + strconv.Itoa(offset) + "B")
}

func renderLogo() {
	logoJoined := strings.Join(entities.LogosMap[flags.Config.LogoName].Logo, "\n")
	os.Stdout.WriteString(logoJoined + seqReset)
}

func renderModule(module flags.Module) {
	cursorDown(1)

	os.Stdout.WriteString("\r" + entities.LogosMap[flags.Config.LogoName].BlankRow)

	os.Stdout.WriteString(flags.Config.LogoSeparator + seqReset)

	os.Stdout.WriteString(module.Title + seqReset)

	os.Stdout.WriteString(flags.Config.ModuleSeparator + seqReset)

	os.Stdout.WriteString(module.Arg + seqReset)
}
