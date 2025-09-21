// Package render provides tools for rendering fetch information
package render

import (
	"fmt"

	"github.com/kourtnet/dummyfetch/internal/args"
)

func Render(logo []string, arguments []args.Arg) {
	maxLen := max(len(logo), len(arguments))

	for i := range maxLen {
		if i < len(logo) {
			fmt.Printf(logo[i])
		} else {
			fmt.Printf(logo[len(logo)-1])
		}

		fmt.Printf("\t")

		if i < len(arguments) {
			fmt.Printf("%s", arguments[i].Name)
			fmt.Printf(": %s", arguments[i].Contents)
		} else {
			fmt.Printf("\n")
		}
	}
}
