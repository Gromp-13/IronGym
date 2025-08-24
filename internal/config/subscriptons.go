package config

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Subscriptions() fyne.CanvasObject {

	labelSub := widget.NewLabel("Абонементы")
	labelSub.TextStyle = fyne.TextStyle{Bold: true}
	labelSub.Move(fyne.NewPos(5, 3))

	labelNewPrice := widget.NewLabel("Новая цена")
	labelNewPrice.TextStyle = fyne.TextStyle{Bold: true}
	labelNewPrice.Move(fyne.NewPos(250, 3))

	labelOldPrice := widget.NewLabel("Старая цена")
	labelOldPrice.TextStyle = fyne.TextStyle{Bold: true}
	labelOldPrice.Move(fyne.NewPos(450, 3))

	labelOneTime := widget.NewLabel("Разовый")
	labelOneTime.TextStyle = fyne.TextStyle{Italic: true}
	labelOneTime.Move(fyne.NewPos(5, 70))

	OneTimeNewPrice := widget.NewEntry()
	OneTimeNewPrice.Resize(fyne.NewSize(110, 40))
	OneTimeNewPrice.Move(fyne.NewPos(245, 70))
	OneTimeNewPrice.SetPlaceHolder("Введите цену")

	OneTimeOldPrice := widget.NewLabel("00.00")
	OneTimeOldPrice.TextStyle = fyne.TextStyle{Italic: true}
	OneTimeOldPrice.Move(fyne.NewPos(480, 70))

	labelMale := widget.NewLabel("Мужской")
	labelMale.TextStyle = fyne.TextStyle{Italic: true}
	labelMale.Move(fyne.NewPos(5, 120))

	MaleNewPrice := widget.NewEntry()
	MaleNewPrice.Resize(fyne.NewSize(110, 40))
	MaleNewPrice.Move(fyne.NewPos(245, 120))
	MaleNewPrice.SetPlaceHolder("Введите цену")

	MaleOldPrice := widget.NewLabel("00.00")
	MaleOldPrice.TextStyle = fyne.TextStyle{Italic: true}
	MaleOldPrice.Move(fyne.NewPos(480, 120))

	labelFemale := widget.NewLabel("Женский")
	labelFemale.TextStyle = fyne.TextStyle{Italic: true}
	labelFemale.Move(fyne.NewPos(5, 170))

	FemaleNewPrice := widget.NewEntry()
	FemaleNewPrice.Resize(fyne.NewSize(110, 40))
	FemaleNewPrice.Move(fyne.NewPos(245, 170))
	FemaleNewPrice.SetPlaceHolder("Введите цену")

	FemaleOldPrice := widget.NewLabel("00.00")
	FemaleOldPrice.TextStyle = fyne.TextStyle{Italic: true}
	FemaleOldPrice.Move(fyne.NewPos(480, 170))

	labelSchool := widget.NewLabel("Школьный")
	labelSchool.TextStyle = fyne.TextStyle{Italic: true}
	labelSchool.Move(fyne.NewPos(5, 220))

	SchoolNewPrice := widget.NewEntry()
	SchoolNewPrice.Resize(fyne.NewSize(110, 40))
	SchoolNewPrice.Move(fyne.NewPos(245, 220))
	SchoolNewPrice.SetPlaceHolder("Введите цену")

	SchoolOldPrice := widget.NewLabel("00.00")
	SchoolOldPrice.TextStyle = fyne.TextStyle{Italic: true}
	SchoolOldPrice.Move(fyne.NewPos(480, 220))

	labelStudent := widget.NewLabel("Студенческий")
	labelStudent.TextStyle = fyne.TextStyle{Italic: true}
	labelStudent.Move(fyne.NewPos(5, 270))

	StudentNewPrice := widget.NewEntry()
	StudentNewPrice.Resize(fyne.NewSize(110, 40))
	StudentNewPrice.Move(fyne.NewPos(245, 270))
	StudentNewPrice.SetPlaceHolder("Введите цену")

	StudentOldPrice := widget.NewLabel("00.00")
	StudentOldPrice.TextStyle = fyne.TextStyle{Italic: true}
	StudentOldPrice.Move(fyne.NewPos(480, 270))

	labelMale_3 := widget.NewLabel("Мужской 3 месяца")
	labelMale_3.TextStyle = fyne.TextStyle{Italic: true}
	labelMale_3.Move(fyne.NewPos(5, 320))

	MaleNewPrice_3 := widget.NewEntry()
	MaleNewPrice_3.Resize(fyne.NewSize(110, 40))
	MaleNewPrice_3.Move(fyne.NewPos(245, 320))
	MaleNewPrice_3.SetPlaceHolder("Введите цену")

	MaleOldPrice_3 := widget.NewLabel("00.00")
	MaleOldPrice_3.TextStyle = fyne.TextStyle{Italic: true}
	MaleOldPrice_3.Move(fyne.NewPos(480, 320))

	labelMale_6 := widget.NewLabel("Мужской пол года")
	labelMale_6.TextStyle = fyne.TextStyle{Italic: true}
	labelMale_6.Move(fyne.NewPos(5, 370))

	MaleNewPrice_6 := widget.NewEntry()
	MaleNewPrice_6.Resize(fyne.NewSize(110, 40))
	MaleNewPrice_6.Move(fyne.NewPos(245, 370))
	MaleNewPrice_6.SetPlaceHolder("Введите цену")

	MaleOldPrice_6 := widget.NewLabel("00.00")
	MaleOldPrice_6.TextStyle = fyne.TextStyle{Italic: true}
	MaleOldPrice_6.Move(fyne.NewPos(480, 370))

	labelMaleYear := widget.NewLabel("Мужской на год")
	labelMaleYear.TextStyle = fyne.TextStyle{Italic: true}
	labelMaleYear.Move(fyne.NewPos(5, 420))

	MaleNewPriceYear := widget.NewEntry()
	MaleNewPriceYear.Resize(fyne.NewSize(110, 40))
	MaleNewPriceYear.Move(fyne.NewPos(245, 420))
	MaleNewPriceYear.SetPlaceHolder("Введите цену")

	MaleOldPriceYear := widget.NewLabel("00.00")
	MaleOldPriceYear.TextStyle = fyne.TextStyle{Italic: true}
	MaleOldPriceYear.Move(fyne.NewPos(480, 420))

	labelFemale_3 := widget.NewLabel("Женский 3 месяца")
	labelFemale_3.TextStyle = fyne.TextStyle{Italic: true}
	labelFemale_3.Move(fyne.NewPos(5, 470))

	FemaleNewPrice_3 := widget.NewEntry()
	FemaleNewPrice_3.Resize(fyne.NewSize(110, 40))
	FemaleNewPrice_3.Move(fyne.NewPos(245, 470))
	FemaleNewPrice_3.SetPlaceHolder("Введите цену")

	FemaleOldPrice_3 := widget.NewLabel("00.00")
	FemaleOldPrice_3.TextStyle = fyne.TextStyle{Italic: true}
	FemaleOldPrice_3.Move(fyne.NewPos(480, 470))

	labelFemale_6 := widget.NewLabel("Женский пол года")
	labelFemale_6.TextStyle = fyne.TextStyle{Italic: true}
	labelFemale_6.Move(fyne.NewPos(5, 520))

	FemaleNewPrice_6 := widget.NewEntry()
	FemaleNewPrice_6.Resize(fyne.NewSize(110, 40))
	FemaleNewPrice_6.Move(fyne.NewPos(245, 520))
	FemaleNewPrice_6.SetPlaceHolder("Введите цену")

	FemaleOldPrice_6 := widget.NewLabel("00.00")
	FemaleOldPrice_6.TextStyle = fyne.TextStyle{Italic: true}
	FemaleOldPrice_6.Move(fyne.NewPos(480, 520))

	labelFemaleYear := widget.NewLabel("Женский на год")
	labelFemaleYear.TextStyle = fyne.TextStyle{Italic: true}
	labelFemaleYear.Move(fyne.NewPos(5, 570))

	FemaleNewPriceYear := widget.NewEntry()
	FemaleNewPriceYear.Resize(fyne.NewSize(110, 40))
	FemaleNewPriceYear.Move(fyne.NewPos(245, 570))
	FemaleNewPriceYear.SetPlaceHolder("Введите цену")

	FemaleOldPriceYear := widget.NewLabel("00.00")
	FemaleOldPriceYear.TextStyle = fyne.TextStyle{Italic: true}
	FemaleOldPriceYear.Move(fyne.NewPos(480, 570))

	cont := container.NewWithoutLayout(
		labelSub,
		labelNewPrice,
		labelOldPrice,
		labelOneTime,
		OneTimeNewPrice,
		OneTimeOldPrice,
		labelMale,
		MaleNewPrice,
		MaleOldPrice,
		labelFemale,
		FemaleNewPrice,
		FemaleOldPrice,
		labelSchool,
		SchoolNewPrice,
		SchoolOldPrice,
		labelStudent,
		StudentNewPrice,
		StudentOldPrice,
		labelMale_3,
		MaleNewPrice_3,
		MaleOldPrice_3,
		labelMale_6,
		MaleNewPrice_6,
		MaleOldPrice_6,
		labelMaleYear,
		MaleNewPriceYear,
		MaleOldPriceYear,
		labelFemale_3,
		FemaleNewPrice_3,
		FemaleOldPrice_3,
		labelFemale_6,
		FemaleNewPrice_6,
		FemaleOldPrice_6,
		labelFemaleYear,
		FemaleNewPriceYear,
		FemaleOldPriceYear,
	)

	return cont

}
