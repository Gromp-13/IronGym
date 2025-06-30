package screens

import "fyne.io/fyne/v2/app"

func Screens() {
	A := app.New()
	Window := A.NewWindow("IronGym")

	Window.ShowAndRun()
}
