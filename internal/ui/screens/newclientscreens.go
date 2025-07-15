package screens

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewClientScreen(a fyne.App) {

	windowNCS := a.NewWindow("Новый клиент")
	windowNCS.Resize(fyne.NewSize(400, 500))
	a.Settings().SetTheme(theme.DarkTheme())

	label := widget.NewLabel("РЕГИСТРАЦИЯ")
	label.TextStyle = fyne.TextStyle{Bold: true}
	label.Move(fyne.NewPos(135, 3))

	userlname := widget.NewEntry()
	userlname.Resize(fyne.NewSize(300, 40))
	userlname.Move(fyne.NewPos(50, 50))
	userlname.SetPlaceHolder("Фамилия")

	userfname := widget.NewEntry()
	userfname.Resize(fyne.NewSize(300, 40))
	userfname.Move(fyne.NewPos(50, 110))
	userfname.SetPlaceHolder("Имя")

	patronymic := widget.NewEntry()
	patronymic.Resize(fyne.NewSize(300, 40))
	patronymic.Move(fyne.NewPos(50, 170))
	patronymic.SetPlaceHolder("Отчество")

	phone := widget.NewEntry()
	phone.Resize(fyne.NewSize(300, 40))
	phone.Move(fyne.NewPos(50, 230))
	phone.SetPlaceHolder("Номер телефона")

	birthdate := widget.NewEntry()
	birthdate.Resize(fyne.NewSize(300, 40))
	birthdate.Move(fyne.NewPos(50, 290))
	birthdate.SetPlaceHolder("Штрихкод карты")

	selectpas := widget.NewSelect(
		[]string{
			"Мужской",
			"Женский",
			"Студенчиски",
			"Пенсионный",
		},

		func(s string) {
			fmt.Printf("Select is %s", s)
		},
	)
	selectpas.Resize(fyne.NewSize(300, 40))
	selectpas.Move(fyne.NewPos(50, 350))
	selectpas.PlaceHolder = "Выберити абонемент"

	save := widget.NewButton("Сохранить", func() { fmt.Println("hi") })
	save.Resize(fyne.NewSize(200, 40))
	save.Move(fyne.NewPos(100, 410))

	cont := container.NewWithoutLayout(
		label,
		userlname,
		userfname,
		patronymic,
		phone,
		birthdate,
		selectpas,
		save,
	)

	windowNCS.SetContent(cont)
	windowNCS.Show()

}
