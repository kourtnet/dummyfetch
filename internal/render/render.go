// Package render provides tools for rendering fetch information
package render

import (
	"fmt"

	"github.com/kourtnet/dummyfetch/internal/entities"
	"github.com/kourtnet/dummyfetch/internal/flags"
)

func Render() {
	logo := entities.LogosMap[flags.Config.LogoName]

	maxLen := max(len(logo.Logo), len(flags.Config.Modules))

	for i := range maxLen {
		if i < len(logo.Logo) {
			fmt.Print(logo.Logo[i])
		} else {
			fmt.Print(logo.BlankRow)
		}

		fmt.Print("  ")

		if i < len(flags.Config.Modules) {
			arg := entities.ArgsMap[flags.Config.Modules[i].Arg]

			titleToPrint := arg.BasicTitle
			if flags.Config.Modules[i].IsTitleSet {
				titleToPrint = flags.Config.Modules[i].Title
			}

			fmt.Printf("\033[%dm%s\033[0m", logo.TextColor, titleToPrint)
			fmt.Printf(": %s\n", arg.Contents)
		} else {
			fmt.Println()
		}
	}
}
