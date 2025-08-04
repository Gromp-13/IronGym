package screens

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Gromp-13/IronGym/internal/db"
	"github.com/Gromp-13/IronGym/internal/models"
)

type SubscriptionOption struct {
	Label    string
	Duration int
}

func NewClientScreen(a fyne.App) {

	var subscriptionOption = []SubscriptionOption{
		{"Мужской", 30},
		{"Женский", 30},
		{"Школьный", 30},
		{"Студенчиский", 30},
		{"Мужской 3 месяца", 90},
		{"Мужской пол года", 1800},
		{"Мужской на год", 365},
		{"Женский 3 месяца", 90},
		{"Женский пол года", 1800},
		{"Женский на год", 365},
	}
	var selectedSubscription SubscriptionOption

	options := make([]string, len(subscriptionOption))
	for i, o := range subscriptionOption{
		options[i] = o.Label
	}

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

	selectSub := widget.NewSelect(options,
		func(s string) {
			for _, o := range subscriptionOption{
				if o.Label == s {
					selectedSubscription = o
					break
				}
			}
		})
	selectSub.Resize(fyne.NewSize(300, 40))
	selectSub.Move(fyne.NewPos(50, 350))
	selectSub.PlaceHolder = "Выберити абонемент"

	save := widget.NewButton("Сохранить", 
	func() {
		if userlname.Text == "" || userfname.Text == ""{
			fmt.Println("Фамилия и имя обязательны")
			return
		}

		parsedDate, err := time.Parse("2006-01-02", birthdate.Text)
		if err != nil {
			fmt.Println("Используйте ГГГГ-ММ-ДД")
			return
		}

		if selectedSubscription.Label == ""{
			fmt.Println("Выберите абонемент")
			return
		}
		client := models.Client{
			LastName: userlname.Text,
			FirstName: userfname.Text,
			MiddleName: patronymic.Text,
			PhoneNumber: phone.Text,
			BirthDate: parsedDate,
			CardBarcode: birthdate.Text,
			GenderID: 0,
		}

		err = db.Repo.NewClients(client)
		if err != nil{
			fmt.Println("Ошибка при добавлении клиента", err)
			return
		}

		createdClients, err, := db.Repo.GetLastClientByBarcode(client.CardBarcode)
		if err != nil{
			fmt.Println("Ошибка при поиске клиента:", err)
			return
		}

		now := time.Now()
		sub := models.Subscriptions{
			ClientID: createdClients.ID,
			StartDate: now,
			DurationDays: selectedSubscription.Duration,
			EndDate: now.AddDate(0, 0, selectedSubscription.Duration),
		}

		err = db.Repo.NewSubscription(sub)
		if err != nil{
			fmt.Println("Ошибка при создании абонимента:", err)
			return
		}

		fmt.Println("Клиент и абонимент успешно добавленны")
		windowNCS.Close()
	})
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
