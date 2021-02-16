package main

import (
	"cd-cpustats/pkg/sys"
	"fmt"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {
	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	stats := &sys.Stats{}
	fmt.Print("This works!!!")
	app := wails.CreateApp(&wails.AppConfig{
		Width:  512,
		Height: 512,
		Title:  "Local Agent",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	app.Bind(stats)
	app.Run()

}
