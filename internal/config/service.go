package config

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Gromp-13/IronGym/internal/db"
	"github.com/Gromp-13/IronGym/internal/models"
)

func Service() fyne.CanvasObject {
	searchLastname := widget.NewEntry()
	searchLastname.Resize(fyne.NewSize(150, 36))
	searchLastname.Move(fyne.NewPos(10, 5))
	searchLastname.SetPlaceHolder("Поиск по фамилии")

	searchBarcode := widget.NewEntry()
	searchBarcode.Resize(fyne.NewSize(150, 36))
	searchBarcode.Move(fyne.NewPos(170, 5))
	searchBarcode.SetPlaceHolder("Поиск по карте")

	// Данные
	allClients, _ := db.Repo.GetClients()
	clients := append([]models.Client(nil), allClients...)
	selectedClients := []models.Client{}

	leftSel := -1
	rightSel := -1

	// Левый список (из базы)
	ClientList := widget.NewList(
		func() int { return len(clients) },
		func() fyne.CanvasObject { return widget.NewLabel("template") },
		func(i widget.ListItemID, o fyne.CanvasObject) {
			client := clients[i]
			o.(*widget.Label).SetText(client.LastName + " " + client.FirstName + " " + client.MiddleName)
		},
	)
	ClientList.Resize(fyne.NewSize(500, 400))
	ClientList.Move(fyne.NewPos(10, 50))
	ClientList.OnSelected = func(id widget.ListItemID) { leftSel = int(id) }
	ClientList.OnUnselected = func(id widget.ListItemID) {
		if leftSel == int(id) {
			leftSel = -1
		}
	}

	// Правый список (выбранные клиенты)
	SelectedList := widget.NewList(
		func() int { return len(selectedClients) },
		func() fyne.CanvasObject { return widget.NewLabel("template") },
		func(i widget.ListItemID, o fyne.CanvasObject) {
			client := selectedClients[i]
			o.(*widget.Label).SetText(client.LastName + " " + client.FirstName + " " + client.MiddleName)
		},
	)
	SelectedList.Resize(fyne.NewSize(500, 400))
	SelectedList.Move(fyne.NewPos(580, 50))
	SelectedList.OnSelected = func(id widget.ListItemID) { rightSel = int(id) }
	SelectedList.OnUnselected = func(id widget.ListItemID) {
		if rightSel == int(id) {
			rightSel = -1
		}
	}

	// Кнопка >
	btn1 := widget.NewButton(">", func() {
		if leftSel >= 0 && leftSel < len(clients) {
			c := clients[leftSel]
			already := false
			for _, s := range selectedClients {
				if s.ID == c.ID {
					already = true
					break
				}
			}
			if !already {
				selectedClients = append(selectedClients, c)
				SelectedList.Refresh()
			}
			ClientList.UnselectAll()
			leftSel = -1
		}
	})
	btn1.Resize(fyne.NewSize(20, 20))
	btn1.Move(fyne.NewPos(540, 170))

	// Кнопка <
	btn2 := widget.NewButton("<", func() {
		if rightSel >= 0 && rightSel < len(selectedClients) {
			selectedClients = append(selectedClients[:rightSel], selectedClients[rightSel+1:]...)
			SelectedList.Refresh()
			SelectedList.UnselectAll()
			rightSel = -1
		}
	})
	btn2.Resize(fyne.NewSize(20, 20))
	btn2.Move(fyne.NewPos(540, 210))

	// Поиск
	applyFilters := func() {
		ln := strings.ToLower(strings.TrimSpace(searchLastname.Text))
		bc := strings.ToLower(strings.TrimSpace(searchBarcode.Text))

		clients = clients[:0]
		for _, c := range allClients {
			okLast := ln == "" || strings.Contains(strings.ToLower(c.LastName), ln)
			okBar := bc == "" || strings.Contains(strings.ToLower(c.CardBarcode), bc)
			if okLast && okBar {
				clients = append(clients, c)
			}
		}
		ClientList.UnselectAll()
		leftSel = -1
		ClientList.Refresh()
	}
	searchLastname.OnChanged = func(string) { applyFilters() }
	searchBarcode.OnChanged = func(string) { applyFilters() }

	// Кнопка "разовый"
	onetimesub := widget.NewButton("Разовый абонемент", func() { fmt.Println("Разовый") })
	onetimesub.Resize(fyne.NewSize(175, 40))
	onetimesub.Move(fyne.NewPos(875, 475))

	cont := container.NewWithoutLayout(
		searchBarcode,
		searchLastname,
		ClientList,
		btn1,
		btn2,
		SelectedList,
		onetimesub,
	)

	return cont
}
