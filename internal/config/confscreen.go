package config

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/Gromp-13/IronGym/internal/ui/screens"
)

func ConfScreen() {
	a := app.New()
	windowMaster := a.NewWindow("IronGym")
	windowMaster.Resize(fyne.NewSize(1000, 600))
	a.Settings().SetTheme(theme.DarkTheme())

	file_item1 := fyne.NewMenuItem("новый клиет", func() {
		screens.NewClientScreen()
	})
	file_item2 := fyne.NewMenuItem("продажа абонемента", func() {
		fmt.Println("продажа абонемента")
	})
	menu1 := fyne.NewMenu("Клиенты", file_item1, file_item2)

	main_menu := fyne.NewMainMenu(menu1)
	windowMaster.SetMainMenu(main_menu)

	windowMaster.Show()
	windowMaster.SetMaster()
}
