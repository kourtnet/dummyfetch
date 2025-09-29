package render

import "os"

func PrepareAndRender() {
	if err := prepare(); err != nil {
		os.Stdout.WriteString(err.Error() + "\n")
		return
	}

	render()
}
