package app

import (
	"editor/internal/editor"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func Start() {
	fmt.Println("starting app")

	a := app.New()
	w := a.NewWindow("Text Editor")

	textBox := editor.NewTextBox()
	content := container.NewVBox(textBox)
	w.SetContent(content)
	w.Resize(fyne.NewSize(200, 200))
	w.ShowAndRun()
}
