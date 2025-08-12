package screens

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	datepicker "github.com/tobbee/fyne-datepicker"

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
		{"Студенческий", 30},
		{"Мужской 3 месяца", 90},
		{"Мужской пол года", 180},
		{"Мужской на год", 365},
		{"Женский 3 месяца", 90},
		{"Женский пол года", 180},
		{"Женский на год", 365},
	}
	var selectedSubscription SubscriptionOption
	var genderID int32
	var selectedBirthDate time.Time

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

	// Дата рождения — аккуратный ввод
	birthLabel := widget.NewLabel("Дата рождения")
	birthLabel.Move(fyne.NewPos(50, 320))

	birthDateEntry := widget.NewEntry()
	birthDateEntry.Resize(fyne.NewSize(250, 40))
	birthDateEntry.Move(fyne.NewPos(50, 350))
	birthDateEntry.Disable()

	birthDateBtn := widget.NewButton("📅", func() {
		popup := a.NewWindow("Выберите дату")

		datePicker := datepicker.NewDatePicker(
			func() time.Time {
				if !selectedBirthDate.IsZero() {
					return selectedBirthDate
				}
				return time.Now()
			}(),
			time.Monday,
			nil, // обработчик тут нам не нужен
		)

		datePicker.Resize(fyne.NewSize(250, 40))

		saveBtn := widget.NewButton("Сохранить", func() {
			selectedBirthDate = datePicker.selectedDate()
			birthDateEntry.SetText(selectedBirthDate.Format("02.01.2006"))
			popup.Close()
		})

		closeBtn := widget.NewButton("Закрыть", func() {
			popup.Close()
		})

		popup.SetContent(container.NewVBox(
			datePicker,
			container.NewHBox(saveBtn, closeBtn),
		))
		popup.Resize(fyne.NewSize(300, 250))
		popup.Show()
	})
	birthDateBtn.Resize(fyne.NewSize(40, 40))
	birthDateBtn.Move(fyne.NewPos(310, 350))

	genderSel := widget.NewSelect(
		[]string{
			"Мужской",
			"Женский",
		},
		func(s string) {
			switch s {
			case "Мужской":
				genderID = 1
			case "Женский":
				genderID = 2
			default:
				genderID = 0
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

	save := widget.NewButton("Сохранить",
		func() {
			if userlname.Text == "" || userfname.Text == "" {
				fmt.Println("Фамилия и имя обязательны")
				return
			}
			if selectedBirthDate.IsZero() {
				fmt.Println("Выберите дату рождения")
				return
			}
			if selectedSubscription.Label == "" {
				fmt.Println("Выберите абонемент")
				return
			}
			if genderID == 0 {
				fmt.Println("Выберите пол")
				return
			}

			client := models.Client{
				LastName:    userlname.Text,
				FirstName:   userfname.Text,
				MiddleName:  patronymic.Text,
				PhoneNumber: phone.Text,
				BirthDate:   selectedBirthDate,
				CardBarcode: barcode.Text,
				GenderID:    genderID,
			}

			err := db.Repo.NewClients(client)
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
		birthLabel,
		birthDateEntry,
		birthDateBtn,
		genderSel,
		selectSub,
		save,
	)
	windowNCS.SetContent(cont)
	windowNCS.Show()
}
