package app

import (
	"hangman/connection"
	"hangman/hangman"
	"hangman/menu"
)

func (app *App) menuMode() {
	menuConnection := connection.MenuConnection{
		Title:    &app.title,
		Menu:     &app.menu,
		Mode:     &app.mode,
		Option:   &app.option,
		Finished: &app.finished,
	}
	menu.InitMenu(&menuConnection)
}

func (app *App) hangmanMode() {
	hangman.Start_game(app.user)
	app.mode = "menu"
	app.menu = "hangman"
}

func (app *App) processMode() {
	switch app.mode {
	case "menu":
		app.menuMode()
	case "hangman":
		app.hangmanMode()
	}
}
