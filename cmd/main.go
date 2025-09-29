package main

import (
	"os"
	"time"

	"github.com/kourtnet/dummyfetch/internal/flags"
	"github.com/kourtnet/dummyfetch/internal/render"
)

func main() {
	start := time.Now()

	flags.Parse()
	render.PrepareAndRender()

	os.Stdout.WriteString(time.Since(start).String())
}
