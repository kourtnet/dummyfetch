// Package render provides tools for rendering fetch information
package render

import (
	"fmt"

	"github.com/kourtnet/dummyfetch/internal/entities"
)

func Render(logoName string, args []string) {
	logo := entities.LogosMap[logoName]

	maxLen := max(len(logo.Logo), len(args))

	for i := range maxLen {
		if i < len(logo.Logo) {
			fmt.Print(logo.Logo[i])
		} else {
			fmt.Print(logo.BlankRow)
		}

		fmt.Print("  ")

		if i < len(args) {
			arg := entities.ArgsMap[args[i]]
			fmt.Printf("\033[%dm%s\033[0m", logo.TextColor, arg.Name)
			fmt.Printf(": %s\n", arg.Contents)
		} else {
			fmt.Println()
		}
	}
}
