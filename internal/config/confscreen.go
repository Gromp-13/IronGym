package config

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/Gromp-13/IronGym/internal/ui/screens"
)

func ConfScreen() {
	myApp := app.New()
	windowMaster := myApp.NewWindow("IronGym")
	windowMaster.Resize(fyne.NewSize(1000, 600))
	myApp.Settings().SetTheme(theme.DarkTheme())

	file_item1 := fyne.NewMenuItem("новый клиет", func() {
		screens.NewClientScreen(myApp)
	})
	file_item2 := fyne.NewMenuItem("продажа абонемента", func() {
		fmt.Println("продажа абонемента")
	})
	menu1 := fyne.NewMenu("Клиенты", file_item1, file_item2)

	main_menu := fyne.NewMainMenu(menu1)
	windowMaster.SetMainMenu(main_menu)

	windowMaster.Show()
	windowMaster.SetMaster()
	myApp.Run()

}
