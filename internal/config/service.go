package config

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Gromp-13/IronGym/internal/db"
)

func Service() fyne.CanvasObject {
	searchlastname := widget.NewEntry()
	searchlastname.Resize(fyne.NewSize(150, 36))
	searchlastname.Move(fyne.NewPos(10, 5))
	searchlastname.SetPlaceHolder("Поиск по фамилии")

	searchbarcode := widget.NewEntry()
	searchbarcode.Resize(fyne.NewSize(150, 36))
	searchbarcode.Move(fyne.NewPos(170, 5))
	searchbarcode.SetPlaceHolder("Поиск по карте")

	clients, _ := db.Repo.GetClient()

	ClientList := widget.NewList(
		func() int {
			return len(clients)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			client := clients[i]
			o.(*widget.Label).SetText(client.LastName + " " + client.FirstName + " " + client.MiddleName)
		},
	)
	ClientList.Resize(fyne.NewSize(500, 400))
	ClientList.Move(fyne.NewPos(10, 50))

	btn1 := widget.NewButton(">", func() { fmt.Println(">") })
	btn1.Resize(fyne.NewSize(20, 20))
	btn1.Move(fyne.NewPos(540, 170))

	btn2 := widget.NewButton("<", func() { fmt.Println("<") })
	btn2.Resize(fyne.NewSize(20, 20))
	btn2.Move(fyne.NewPos(540, 210))

	scrollclientstrening := container.NewVScroll(
		container.NewVBox(
			widget.NewLabel("1"),
			widget.NewLabel("2"),
			widget.NewLabel("3"),
			widget.NewLabel("4"),
			widget.NewLabel("5"),
			widget.NewLabel("6"),
			widget.NewLabel("7"),
			widget.NewLabel("8"),
			widget.NewLabel("9"),
			widget.NewLabel("10"),
			widget.NewLabel("11"),
			widget.NewLabel("12"),
			widget.NewLabel("13"),
			widget.NewLabel("14"),
			widget.NewLabel("15"),
			widget.NewLabel("16"),
			widget.NewLabel("17"),
			widget.NewLabel("18"),
			widget.NewLabel("19"),
			widget.NewLabel("20"),
			widget.NewLabel("21"),
			widget.NewLabel("22"),
			widget.NewLabel("23"),
			widget.NewLabel("24"),
			widget.NewLabel("25"),
			widget.NewLabel("26"),
		),
	)
	scrollclientstrening.Resize(fyne.NewSize(500, 400))
	scrollclientstrening.Move(fyne.NewPos(580, 50))

	onetimesub := widget.NewButton("Разовый абонемент", func() { fmt.Println("Разовый") })
	onetimesub.Resize(fyne.NewSize(175, 40))
	onetimesub.Move(fyne.NewPos(875, 475))

	cont := container.NewWithoutLayout(
		searchbarcode,
		searchlastname,
		ClientList,
		btn1,
		btn2,
		scrollclientstrening,
		onetimesub,
	)

	return cont
}
