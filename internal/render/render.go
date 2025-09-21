// Package render provides tools for rendering fetch information
package render

import (
	"fmt"

	"github.com/kourtnet/dummyfetch/internal/args"
	"github.com/kourtnet/dummyfetch/internal/logos"
)

func Render(logo logos.LogoInfo, arguments []args.Arg) {
	maxLen := max(len(logo.Logo), len(arguments))

	for i := range maxLen {
		if i < len(logo.Logo) {
			fmt.Print(logo.Logo[i])
		} else {
			fmt.Print(logo.Logo[len(logo.Logo)-1])
		}

		fmt.Print("\t")

		if i < len(arguments) {
			fmt.Printf("\033[%dm%s\033[0m", logo.Color, arguments[i].Name)
			fmt.Printf(": %s", arguments[i].Contents)
		} else {
			fmt.Printf("\n")
		}
	}
}
