package main

import (
	"image/color"
	"io/ioutil"
	"log"
	"net/http"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("presto")

	text1 := canvas.NewText("GET https://api.chucknorris.io/jokes/random", color.White)
	button1 := widget.NewButtonWithIcon("Presto!", theme.MediaFastForwardIcon(), func() {
		log.Println("Grazie!")

		resp, err := http.Get("https://api.chucknorris.io/jokes/random")
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		sb := string(body)
		log.Println(sb)
	})

	myWindow.SetContent(container.New(layout.NewVBoxLayout(), text1, button1))
	myWindow.ShowAndRun()
}
