package config

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Gromp-13/IronGym/internal/db"
)

func Subscriptions() fyne.CanvasObject {
	cont := container.NewWithoutLayout()

	// Заголовки
	labelSub := widget.NewLabel("Абонементы")
	labelSub.TextStyle = fyne.TextStyle{Bold: true}
	labelSub.Move(fyne.NewPos(5, 3))
	cont.Add(labelSub)

	labelNewPrice := widget.NewLabel("Новая цена")
	labelNewPrice.TextStyle = fyne.TextStyle{Bold: true}
	labelNewPrice.Move(fyne.NewPos(250, 3))
	cont.Add(labelNewPrice)

	labelOldPrice := widget.NewLabel("Старая цена")
	labelOldPrice.TextStyle = fyne.TextStyle{Bold: true}
	labelOldPrice.Move(fyne.NewPos(450, 3))
	cont.Add(labelOldPrice)

	// Получаем все абонементы из базы
	memberships, _ := db.Repo.GetMemberships()

	for i, m := range memberships {
		y := float32(70 + i*50)

		// Название абонемента
		lbl := widget.NewLabel(m.Name)
		lbl.TextStyle = fyne.TextStyle{Italic: true}
		lbl.Move(fyne.NewPos(5, y))
		cont.Add(lbl)

		// Поле для новой цены
		newPrice := widget.NewEntry()
		newPrice.Resize(fyne.NewSize(110, 40))
		newPrice.Move(fyne.NewPos(245, y))
		newPrice.SetPlaceHolder("Введите цену")
		cont.Add(newPrice)

		// Старая цена
		oldPrice := widget.NewLabel(strconv.Itoa(int(m.Price)))
		oldPrice.TextStyle = fyne.TextStyle{Italic: true}
		oldPrice.Move(fyne.NewPos(480, y))
		cont.Add(oldPrice)

		// Кнопка "Сохранить"
		saveBtn := widget.NewButton("Сохранить", func(id int32, np *widget.Entry, op *widget.Label) func() {
			return func() {
				price, _ := strconv.Atoi(np.Text)
				db.Repo.UpdateMembershipPrice(id, int32(price))
				op.SetText(np.Text)
			}
		}(m.ID, newPrice, oldPrice))
		saveBtn.Resize(fyne.NewSize(110, 40))
		saveBtn.Move(fyne.NewPos(600, y))
		cont.Add(saveBtn)
	}

	return cont
}
