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
	labelMonth.Move(fyne.NewPos(350, 3))

	labelQuarter := widget.NewLabel("Квартал")
	labelQuarter.TextStyle = fyne.TextStyle{Bold: true}
	labelQuarter.Move(fyne.NewPos(500, 3))

	labeHalf := widget.NewLabel("Пол года")
	labeHalf.TextStyle = fyne.TextStyle{Bold: true}
	labeHalf.Move(fyne.NewPos(650, 3))

	labelYear := widget.NewLabel("Год")
	labelYear.TextStyle = fyne.TextStyle{Bold: true}
	labelYear.Move(fyne.NewPos(800, 3))

	labelTotal := widget.NewLabel("Всего")
	labelTotal.TextStyle = fyne.TextStyle{Bold: true}
	labelTotal.Move(fyne.NewPos(950, 3))

	labelOneTime := widget.NewLabel("Разовый")
	labelOneTime.TextStyle = fyne.TextStyle{Italic: true}
	labelOneTime.Move(fyne.NewPos(5, 60))

	OneTimeMonth := widget.NewLabel("00.00")
	OneTimeMonth.TextStyle = fyne.TextStyle{Italic: true}
	OneTimeMonth.Move(fyne.NewPos(350, 60))

	OneTimeQuarter := widget.NewLabel("00.00")
	OneTimeQuarter.TextStyle = fyne.TextStyle{Italic: true}
	OneTimeQuarter.Move(fyne.NewPos(500, 60))

	OneTimeHalf := widget.NewLabel("00.00")
	OneTimeHalf.TextStyle = fyne.TextStyle{Italic: true}
	OneTimeHalf.Move(fyne.NewPos(650, 60))

	OneTimeYear := widget.NewLabel("00.00")
	OneTimeYear.TextStyle = fyne.TextStyle{Italic: true}
	OneTimeYear.Move(fyne.NewPos(800, 60))

	OneTimeTotal := widget.NewLabel("00.00")
	OneTimeTotal.TextStyle = fyne.TextStyle{Italic: true}
	OneTimeTotal.Move(fyne.NewPos(950, 60))

	labelMale := widget.NewLabel("Мужской")
	labelMale.TextStyle = fyne.TextStyle{Italic: true}
	labelMale.Move(fyne.NewPos(5, 100))

	MaleMonth := widget.NewLabel("00.00")
	MaleMonth.TextStyle = fyne.TextStyle{Italic: true}
	MaleMonth.Move(fyne.NewPos(350, 100))

	MaleQuarter := widget.NewLabel("00.00")
	MaleQuarter.TextStyle = fyne.TextStyle{Italic: true}
	MaleQuarter.Move(fyne.NewPos(500, 100))

	MaleHalf := widget.NewLabel("00.00")
	MaleHalf.TextStyle = fyne.TextStyle{Italic: true}
	MaleHalf.Move(fyne.NewPos(650, 100))

	MaleYear := widget.NewLabel("00.00")
	MaleYear.TextStyle = fyne.TextStyle{Italic: true}
	MaleYear.Move(fyne.NewPos(800, 100))

	MaleTotal := widget.NewLabel("00.00")
	MaleTotal.TextStyle = fyne.TextStyle{Italic: true}
	MaleTotal.Move(fyne.NewPos(950, 100))

	labelFemale := widget.NewLabel("Женский")
	labelFemale.TextStyle = fyne.TextStyle{Italic: true}
	labelFemale.Move(fyne.NewPos(5, 140))

	FemaleMonth := widget.NewLabel("00.00")
	FemaleMonth.TextStyle = fyne.TextStyle{Italic: true}
	FemaleMonth.Move(fyne.NewPos(350, 140))

	FemaleQuarter := widget.NewLabel("00.00")
	FemaleQuarter.TextStyle = fyne.TextStyle{Italic: true}
	FemaleQuarter.Move(fyne.NewPos(500, 140))

	FemaleHalf := widget.NewLabel("00.00")
	FemaleHalf.TextStyle = fyne.TextStyle{Italic: true}
	FemaleHalf.Move(fyne.NewPos(650, 140))

	FemaleYear := widget.NewLabel("00.00")
	FemaleYear.TextStyle = fyne.TextStyle{Italic: true}
	FemaleYear.Move(fyne.NewPos(800, 140))

	FemaleTotal := widget.NewLabel("00.00")
	FemaleTotal.TextStyle = fyne.TextStyle{Italic: true}
	FemaleTotal.Move(fyne.NewPos(950, 140))

	labelSchool := widget.NewLabel("Школьный")
	labelSchool.TextStyle = fyne.TextStyle{Italic: true}
	labelSchool.Move(fyne.NewPos(5, 180))

	SchoolMonth := widget.NewLabel("00.00")
	SchoolMonth.TextStyle = fyne.TextStyle{Italic: true}
	SchoolMonth.Move(fyne.NewPos(350, 180))

	SchoolQuarter := widget.NewLabel("00.00")
	SchoolQuarter.TextStyle = fyne.TextStyle{Italic: true}
	SchoolQuarter.Move(fyne.NewPos(500, 180))

	SchoolHalf := widget.NewLabel("00.00")
	SchoolHalf.TextStyle = fyne.TextStyle{Italic: true}
	SchoolHalf.Move(fyne.NewPos(650, 180))

	SchoolYear := widget.NewLabel("00.00")
	SchoolYear.TextStyle = fyne.TextStyle{Italic: true}
	SchoolYear.Move(fyne.NewPos(800, 180))

	SchoolTotal := widget.NewLabel("00.00")
	SchoolTotal.TextStyle = fyne.TextStyle{Italic: true}
	SchoolTotal.Move(fyne.NewPos(950, 180))

	labelStudent := widget.NewLabel("Студенческий")
	labelStudent.TextStyle = fyne.TextStyle{Italic: true}
	labelStudent.Move(fyne.NewPos(5, 220))

	StudentMonth := widget.NewLabel("00.00")
	StudentMonth.TextStyle = fyne.TextStyle{Italic: true}
	StudentMonth.Move(fyne.NewPos(350, 220))

	StudentQuarter := widget.NewLabel("00.00")
	StudentQuarter.TextStyle = fyne.TextStyle{Italic: true}
	StudentQuarter.Move(fyne.NewPos(500, 220))

	StudentHalf := widget.NewLabel("00.00")
	StudentHalf.TextStyle = fyne.TextStyle{Italic: true}
	StudentHalf.Move(fyne.NewPos(650, 220))

	StudentYear := widget.NewLabel("00.00")
	StudentYear.TextStyle = fyne.TextStyle{Italic: true}
	StudentYear.Move(fyne.NewPos(800, 220))

	StudentTotal := widget.NewLabel("00.00")
	StudentTotal.TextStyle = fyne.TextStyle{Italic: true}
	StudentTotal.Move(fyne.NewPos(950, 220))

	labelMale_3 := widget.NewLabel("Мужской 3 месяца")
	labelMale_3.TextStyle = fyne.TextStyle{Italic: true}
	labelMale_3.Move(fyne.NewPos(5, 260))

	Male_3_Month := widget.NewLabel("00.00")
	Male_3_Month.TextStyle = fyne.TextStyle{Italic: true}
	Male_3_Month.Move(fyne.NewPos(350, 260))

	Male_3_Quarter := widget.NewLabel("00.00")
	Male_3_Quarter.TextStyle = fyne.TextStyle{Italic: true}
	Male_3_Quarter.Move(fyne.NewPos(500, 260))

	Male_3_Half := widget.NewLabel("00.00")
	Male_3_Half.TextStyle = fyne.TextStyle{Italic: true}
	Male_3_Half.Move(fyne.NewPos(650, 260))

	Male_3_Year := widget.NewLabel("00.00")
	Male_3_Year.TextStyle = fyne.TextStyle{Italic: true}
	Male_3_Year.Move(fyne.NewPos(800, 260))

	Male_3_Total := widget.NewLabel("00.00")
	Male_3_Total.TextStyle = fyne.TextStyle{Italic: true}
	Male_3_Total.Move(fyne.NewPos(950, 260))

	cont := container.NewWithoutLayout(
		labelSub,
		labelMonth,
		labelQuarter,
		labeHalf,
		labelYear,
		labelTotal,
		labelOneTime,
		OneTimeMonth,
		OneTimeQuarter,
		OneTimeHalf,
		OneTimeYear,
		OneTimeTotal,
		labelMale,
		MaleMonth,
		MaleQuarter,
		MaleHalf,
		MaleYear,
		MaleTotal,
		labelFemale,
		FemaleMonth,
		FemaleQuarter,
		FemaleHalf,
		FemaleYear,
		FemaleTotal,
		labelSchool,
		SchoolMonth,
		SchoolQuarter,
		SchoolHalf,
		SchoolYear,
		SchoolTotal,
		labelStudent,
		StudentMonth,
		StudentQuarter,
		StudentHalf,
		StudentYear,
		StudentTotal,
		labelMale_3,
		Male_3_Month,
		Male_3_Quarter,
		Male_3_Half,
		Male_3_Year,
		Male_3_Total,
	)

	return cont
}
