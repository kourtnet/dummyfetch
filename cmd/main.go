package main

import (
	"fmt"
	"time"

	"github.com/kourtnet/dummyfetch/internal/flags"
	"github.com/kourtnet/dummyfetch/internal/render"
	"github.com/kourtnet/dummyfetch/internal/sysinfo"
)

func main() {
	start := time.Now()
	cfg, err := flags.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := sysinfo.Fetch(cfg.ArgsOrder); err != nil {
		fmt.Println(err)
		return
	}

	logo, err := sysinfo.FetchLogo(cfg.Logo)
	if err != nil {
		fmt.Println(err)
		return
	}

	render.Render(logo, cfg.ArgsOrder)

	fmt.Println(time.Since(start))
}
