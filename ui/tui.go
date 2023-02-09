package ui

import (
	"strconv"

	"github.com/Dee-6777/jokesforall/cmd"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Datas struct {
	firstName string
	number    string
}

var data = make([]Datas, 0)

// Tview
var pages = tview.NewPages()
var jokeText = tview.NewTextView()
var app = tview.NewApplication()
var form = tview.NewForm()
var jokesList = tview.NewList().ShowSecondaryText(false)
var flex = tview.NewFlex()
var text = tview.NewTextView().
	SetTextColor(tcell.ColorPaleVioletRed).
	SetText("(a) get random joke \n(b) get n no of jokes \n(q) quit")

func Greet() {

	jokesList.SetSelectedFunc(func(index int, name string, number string, shortcut rune) {
		setConcatText(&data[index])
	})
	flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().
			AddItem(jokesList, 0, 1, true).
			AddItem(jokeText, 0, 4, false), 0, 6, false).
		AddItem(text, 0, 1, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 97 {
			form.Clear(true)
			addContactForm_()
			pages.SwitchToPage("Joke of the day")
		} else if event.Rune() == 98 {
			form.Clear(true)
			addContactForm()
			pages.SwitchToPage("Your Jokes")
		} else if event.Rune() == 113 {
			app.Stop()
		}
		return event
	})
	pages.AddPage("Menu", flex, true, true)
	pages.AddPage("Your Jokes", form, true, false)
	pages.AddPage("Joke of the day", form, true, false)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func addJokeList() {
	jokesList.Clear()
	for index, jokes := range data {
		jokesList.AddItem(jokes.firstName+" "+jokes.number, " ", rune(49+index), nil)
	}
}

func addContactForm() *tview.Form {

	jokes_ := Datas{}

	form.AddInputField("First Name", "", 20, nil, func(firstName string) {
		jokes_.firstName = firstName
	})

	form.AddInputField("Number", "", 20, nil, func(number string) {
		jokes_.number = number
	})

	form.AddButton("Save", func() {
		data = append(data, jokes_)
		addJokeList()
		pages.SwitchToPage("Menu")
	})

	return form
}

func addContactForm_() *tview.Form {

	jokes_ := Datas{}

	form.AddInputField("First Name", "", 20, nil, func(firstName string) {
		jokes_.firstName = firstName
	})

	jokes_.number = ""

	form.AddButton("Save", func() {
		data = append(data, jokes_)
		addJokeList()
		pages.SwitchToPage("Menu")
	})

	return form
}

func setConcatText(contact *Datas) {
	jokeText.Clear()
	text := "Hey " + contact.firstName + " ready to laugh?\n"
	if contact.number != "" {
		t, _ := strconv.Atoi(contact.number)
		for i := 0; i < t; i++ {
			text += cmd.GetRandomJoke() + "\n"
		}
		if t < 1 {
			text += "You forget to add a number. Enter a number, please."
		}
	} else {
		text += cmd.GetRandomJoke()
	}
	jokeText.SetText(text)
}
