package app

import (
	"fmt"
)

func (app *App) Init() {
	app.title = "Hangman Game"
	app.menu = "main"
	app.user = "Nicolas"
	for !app.finished {
		app.printMenu()
		fmt.Print("Ingrese una opcion: ")
		fmt.Scanln(&app.option)
		app.chooseOption()
	}
}
