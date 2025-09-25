package main

import (
	"fmt"

	"github.com/kourtnet/dummyfetch/internal/flags"
)

func main() {
	// start := time.Now()
	err := flags.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(flags.Config)
	//
	//	if err := sysinfo.Fetch(cfg.ArgsOrder); err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	// logo, err := sysinfo.FetchLogo(cfg.Logo)
	//
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	// render.Render(logo, cfg.ArgsOrder)
	//
	// fmt.Println(time.Since(start))
}
