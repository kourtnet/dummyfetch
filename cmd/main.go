package main

import (
	"fmt"
	"os"

	"github.com/kourtnet/dummyfetch/internal/flags"
	"github.com/kourtnet/dummyfetch/internal/render"
	"github.com/kourtnet/dummyfetch/internal/sysinfo"
)

func main() {
	cfg, err := flags.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	args, err := sysinfo.Fetch(cfg.ArgsOrder)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logo, err := sysinfo.ResolveIcon(cfg.LogoName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	render.Render(logo, args)
}
