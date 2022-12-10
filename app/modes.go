package app

import (
	"fmt"
	"hangman/hangman"
)

func (app *App) menuMode() {
	app.printMenu()
	fmt.Print("Ingrese una opcion: ")
	fmt.Scanln(&app.option)
	app.chooseOption()
}

func (app *App) hangmanMode() {
	hangman.Start_game(app.user)
	app.mode = "menu"
	app.menu = "hangman"
}

func (app *App) selectMode() {
	switch app.mode {
	case "menu":
		app.menuMode()
	case "hangman":
		app.hangmanMode()
	}
}
