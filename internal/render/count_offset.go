package render

import (
	"strconv"

	"github.com/kourtnet/dummyfetch/internal/entities"
	"github.com/kourtnet/dummyfetch/internal/flags"
)

func countStrOffset(str string, startOffset int) (int, int) {
	currOffset, maxOffset := startOffset, 0

	for _, match := range offsetCountingRegex.FindAllStringSubmatch(str, -1) {
		if match[1] != "" {
			n, _ := strconv.Atoi(match[1])

			if match[2] == "B" {
				currOffset += n
			} else if match[2] == "A" {
				currOffset -= n
			}
		} else if match[3] != "" {
			if currOffset > 0 {
				maxOffset = max(maxOffset, currOffset)
			}
		}
	}

	return currOffset, maxOffset
}

func countModsOffset() int {
	var maxOffset, currOffset int

	for _, mod := range flags.Config.Modules {
		str := flags.Config.LogoSeparator + mod.Title + flags.Config.ModuleSeparator + mod.Arg

		var offset int
		offset, currOffset = countStrOffset(str, currOffset)

		maxOffset = max(maxOffset, offset)
		currOffset++
	}

	return maxOffset
}

func countLogoOffset() int {
	return len(entities.LogosMap[flags.Config.LogoName].Logo)
}
