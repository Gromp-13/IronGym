package config

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/Gromp-13/IronGym/internal/ui/screens"
)

func Config() {

	myApp := app.New()
	windowMaster := myApp.NewWindow("IronGym")
	windowMaster.Resize(fyne.NewSize(1100, 700))
	myApp.Settings().SetTheme(theme.DarkTheme())

	newClientItem := fyne.NewMenuItem("новый клиет", func() {
		screens.NewClientScreen(myApp)
	})
	sellSubscriptionItem := fyne.NewMenuItem("продажа абонемента", func() {
		fmt.Println("продажа абонемента")
	})
	menu1 := fyne.NewMenu("Клиенты", newClientItem, sellSubscriptionItem)

	main_menu := fyne.NewMainMenu(menu1)
	windowMaster.SetMainMenu(main_menu)

	tabs := container.NewAppTabs(
		container.NewTabItem("Сервис", Service()),
		container.NewTabItem("Финансы", Finance()),
		container.NewTabItem("Абонементы", Subscriptions()),
	)

	tabs.SetTabLocation(container.TabLocationBottom)

	windowMaster.SetContent(tabs)
	windowMaster.Show()
	myApp.Run()

}
