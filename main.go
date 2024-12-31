package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/swldmnn/presto/data"
	"github.com/swldmnn/presto/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	requests := data.BuildRequests()
	requestBinding := binding.NewUntypedList()
	for _, t := range requests {
		requestBinding.Append(t)
	}

	responseText := widget.NewLabel("")
	responseText.Wrapping = fyne.TextWrapBreak

	requestList := widget.NewListWithData(
		requestBinding,
		func() fyne.CanvasObject {
			return container.NewBorder(
				nil,
				nil,
				widget.NewLabel(""),
				widget.NewButton("presto", func() {}),
				widget.NewLabel(""),
			)
		},
		func(dataItem binding.DataItem, o fyne.CanvasObject) {
			ctr, _ := o.(*fyne.Container)

			labelUrl := ctr.Objects[0].(*widget.Label)
			labelMethod := ctr.Objects[1].(*widget.Label)
			buttonSend := ctr.Objects[2].(*widget.Button)

			dataItemUntyped, _ := dataItem.(binding.Untyped).Get()
			request := dataItemUntyped.(model.HttpRequest)

			labelMethod.SetText(request.Method)
			labelUrl.SetText(request.Url)
			buttonSend.SetIcon(theme.MediaFastForwardIcon())
			buttonSend.OnTapped = func() { sendRequest(request.Url, responseText) }
		})

	myApp := app.New()
	myWindow := myApp.NewWindow("presto")
	myWindow.Resize(fyne.NewSize(500, 500))

	myWindow.SetContent(container.New(layout.NewGridLayout(1), requestList, responseText))
	myWindow.ShowAndRun()
}

func sendRequest(url string, textField *widget.Label) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	textField.SetText(sb)
}
