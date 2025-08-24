package config

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Finance() fyne.CanvasObject {

	labelSub := widget.NewLabel("Абонементы")
	labelSub.TextStyle = fyne.TextStyle{Bold: true}
	labelSub.Move(fyne.NewPos(5, 3))

	labelMonth := widget.NewLabel("Месяц")
	labelMonth.TextStyle = fyne.TextStyle{Bold: true}
	labelMonth.Move(fyne.NewPos(200, 3))

	labelQuarter := widget.NewLabel("Квартал")
	labelQuarter.TextStyle = fyne.TextStyle{Bold: true}
	labelQuarter.Move(fyne.NewPos(400, 3))

	labeHalf := widget.NewLabel("Пол года")
	labeHalf.TextStyle = fyne.TextStyle{Bold: true}
	labeHalf.Move(fyne.NewPos(600, 3))

	labelYear := widget.NewLabel("Год")
	labelYear.TextStyle = fyne.TextStyle{Bold: true}
	labelYear.Move(fyne.NewPos(800, 3))

	labelTotal := widget.NewLabel("Всего")
	labelTotal.TextStyle = fyne.TextStyle{Bold: true}
	labelTotal.Move(fyne.NewPos(1000, 3))

	cont := container.NewWithoutLayout(
		labelSub,
		labelMonth,
		labelQuarter,
		labeHalf,
		labelYear,
		labelTotal,
	)

	return cont
}
