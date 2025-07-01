package config

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func ConfScreen() {
	a := app.New()
	window := a.NewWindow("IronGym")
	window.Resize(fyne.NewSize(1000, 600))
	a.Settings().SetTheme(theme.DarkTheme())

	label := widget.NewLabel("hello")

	window.SetContent(container.NewVBox(
		label,
	))

	window.ShowAndRun()
}
