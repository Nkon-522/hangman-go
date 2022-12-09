package app

import (
	"fmt"
	"hangman/terminal"
	"log"
)

func (app *App) printMenuOptions() {
	for index, option := range app.menus[app.menu] {
		fmt.Printf("%d. %s\n", index+1, option)
	}
}

func (app *App) printMenu() {
	terminal.CallClear()

	fmt.Println(app.title)
	fmt.Println("------------")

	app.printMenuOptions()
}

func (app *App) chooseOption() {
	option, ok := app.options[app.menu]
	if !ok {
		log.Fatal("Something went wrong!")
	}
	option(app)
}
