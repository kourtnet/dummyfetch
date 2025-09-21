package main

import (
	"github.com/kourtnet/dummyfetch/internal/logos"
	"github.com/kourtnet/dummyfetch/internal/render"
	"github.com/kourtnet/dummyfetch/internal/sysinfo"
)

func main() {
	stats, err := sysinfo.Fetch()
	if err != nil {
		panic(err)
	}

	render.Render(logos.ArchTiny, stats)
}
