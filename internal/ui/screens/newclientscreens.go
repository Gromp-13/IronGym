package screens

import (
	"fmt"
	"strings"
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
		{"Мужской", 31},
		{"Женский", 31},
		{"Школьный", 31},
		{"Студенческий", 31},
		{"Мужской 3 месяца", 90},
		{"Мужской пол года", 180},
		{"Мужской на год", 365},
		{"Женский 3 месяца", 90},
		{"Женский пол года", 180},
		{"Женский на год", 365},
	}
	var selectedSubscription SubscriptionOption
	var gender int32

	options := make([]string, len(subscriptionOption))
	for i, o := range subscriptionOption {
		options[i] = o.Label
	}

	windowNCS := a.NewWindow("Новый клиент")
	windowNCS.Resize(fyne.NewSize(400, 600))
	a.Settings().SetTheme(theme.DarkTheme())

	label := widget.NewLabel("РЕГИСТРАЦИЯ")
	label.TextStyle = fyne.TextStyle{Bold: true}
	label.Move(fyne.NewPos(135, 3))

	userlname := widget.NewEntry()
	userlname.Resize(fyne.NewSize(300, 40))
	userlname.Move(fyne.NewPos(50, 40))
	userlname.SetPlaceHolder("Фамилия")

	userfname := widget.NewEntry()
	userfname.Resize(fyne.NewSize(300, 40))
	userfname.Move(fyne.NewPos(50, 100))
	userfname.SetPlaceHolder("Имя")

	patronymic := widget.NewEntry()
	patronymic.Resize(fyne.NewSize(300, 40))
	patronymic.Move(fyne.NewPos(50, 160))
	patronymic.SetPlaceHolder("Отчество")

	phone := widget.NewEntry()
	phone.Resize(fyne.NewSize(300, 40))
	phone.Move(fyne.NewPos(50, 220))
	phone.SetPlaceHolder("Номер телефона")

	barcode := widget.NewEntry()
	barcode.Resize(fyne.NewSize(300, 40))
	barcode.Move(fyne.NewPos(50, 280))
	barcode.SetPlaceHolder("Штрихкод карты")

	birthDate := widget.NewEntry()
	birthDate.Resize(fyne.NewSize(300, 40))
	birthDate.Move(fyne.NewPos(50, 340))
	birthDate.SetPlaceHolder("ГГГГ-ММ-ДД")

	genderSel := widget.NewSelect(
		[]string{"Мужской", "Женский"},
		func(s string) {
			switch s {
			case "Мужской":
				gender = 1
			case "Женский":
				gender = 2
			default:
				gender = 0
			}
		},
	)
	genderSel.Resize(fyne.NewSize(300, 40))
	genderSel.Move(fyne.NewPos(50, 410))
	genderSel.PlaceHolder = "Выберите пол"

	selectSub := widget.NewSelect(options,
		func(s string) {
			for _, o := range subscriptionOption {
				if o.Label == s {
					selectedSubscription = o
					break
				}
			}
		})
	selectSub.Resize(fyne.NewSize(300, 40))
	selectSub.Move(fyne.NewPos(50, 470))
	selectSub.PlaceHolder = "Выберите абонемент"

	save := widget.NewButton("Сохранить", func() {
		if userlname.Text == "" || userfname.Text == "" {
			fmt.Println("Фамилия и имя обязательны")
			return
		}
		if birthDate.Text == "" {
			fmt.Println("Введите дату рождения")
			return
		}

		parsedDate, err := time.Parse("2006-01-02", strings.TrimSpace(birthDate.Text))
		if err != nil {
			fmt.Println("Неверный формат даты. Используйте ГГГГ-ММ-ДД")
			return
		}

		// Проверка на будущее и возраст
		if parsedDate.After(time.Now()) {
			fmt.Println("Дата рождения не может быть в будущем")
			return
		}
		age := int(time.Since(parsedDate).Hours() / 24 / 365)
		if age < 10 {
			fmt.Println("Клиенту должно быть не меньше 10 лет")
			return
		}

		if selectedSubscription.Label == "" {
			fmt.Println("Выберите абонемент")
			return
		}
		if gender == 0 {
			fmt.Println("Выберите пол")
			return
		}
		client := models.Client{
			LastName:    userlname.Text,
			FirstName:   userfname.Text,
			MiddleName:  patronymic.Text,
			PhoneNumber: phone.Text,
			BirthDate:   parsedDate,
			CardBarcode: barcode.Text,
			Gender:      gender,
		}

		err = db.Repo.NewClients(client)
		if err != nil {
			fmt.Println("Ошибка при добавлении клиента:", err)
			return
		}

		createdClient, err := db.Repo.GetLastClientByBarcode(client.CardBarcode)
		if err != nil {
			fmt.Println("Ошибка при поиске клиента:", err)
			return
		}

		now := time.Now()
		sub := models.Subscriptions{
			ClientID:     createdClient.ID,
			StartDate:    now,
			DurationDays: int32(selectedSubscription.Duration),
			EndDate:      now.AddDate(0, 0, selectedSubscription.Duration),
		}

		err = db.Repo.NewSubscription(sub)
		if err != nil {
			fmt.Println("Ошибка при создании абонемента:", err)
			return
		}

		fmt.Println("Клиент и абонемент успешно добавлены")
		windowNCS.Close()
	})
	save.Resize(fyne.NewSize(200, 40))
	save.Move(fyne.NewPos(100, 530))

	cont := container.NewWithoutLayout(
		label,
		userlname,
		userfname,
		patronymic,
		phone,
		barcode,
		birthDate,
		genderSel,
		selectSub,
		save,
	)
	windowNCS.SetContent(cont)
	windowNCS.Show()
}
