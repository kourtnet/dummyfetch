package main

import (
	"github.com/kourtnet/dummyfetch/internal/flags"
	"github.com/kourtnet/dummyfetch/internal/render"
)

func main() {
	flags.Parse()
	render.PrepareAndRender()
}
