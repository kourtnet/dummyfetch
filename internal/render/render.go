// Package render provides tools for rendering fetch information
package render

import (
	"fmt"

	"github.com/kourtnet/dummyfetch/internal/entities"
)

func Render(logo entities.LogoInfo, args []entities.Arg) {
	maxLen := max(len(logo.Logo), len(args))

	for i := range maxLen {
		if i < len(logo.Logo) {
			fmt.Print(logo.Logo[i])
		} else {
			fmt.Print(logo.BlankRow)
		}

		if i < len(args) {
			fmt.Printf("\033[%dm%s\033[0m", logo.TextColor, args[i].Name)
			fmt.Printf(": %s\n", args[i].Contents)
		} else {
			fmt.Printf("\n")
		}

	}
}
