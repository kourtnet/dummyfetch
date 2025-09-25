package main

import (
	"fmt"

	"github.com/kourtnet/dummyfetch/internal/flags"
	"github.com/kourtnet/dummyfetch/internal/render"
	"github.com/kourtnet/dummyfetch/internal/sysinfo"
)

func main() {
	err := flags.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := sysinfo.Fetch(); err != nil {
		fmt.Println(err)
		return
	}

	if err := sysinfo.FetchLogo(); err != nil {
		fmt.Println(err)
		return
	}

	render.Render()
}
