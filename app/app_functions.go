package app

import (
	"fmt"
	"hangman/terminal"
)

func printOptions(menuOptions []string) {
	for index, option := range menuOptions {
		fmt.Printf("%d. %s\n", index+1, option)
	}
}

func (app *App) printMenu() {
	terminal.CallClear()

	fmt.Println("Hangman Game")
	fmt.Println("------------")

	if menu := app.menu; menu == "main" {
		printOptions(mainMenuOptions[:])
	} else if menu == "game" {
		printOptions(gameMenuOptions[:])
	} else if menu == "login" {
		printOptions(loginMenuOptions[:])
	} else if menu == "leaderboard" {
		printOptions(leaderboardMenuOptions[:])
	} else if menu == "exit" {
		printOptions(exitMenuOptions[:])
	}
}

func (app *App) chooseOption() {

	if menu := app.menu; menu == "main" {
		app.chooseMainMenu()
	} else if menu == "game" {
		app.chooseGameMenu()
	} else if menu == "login" {
		app.chooseLoginMenu()
	} else if menu == "leaderboard" {
		app.chooseLeaderboardMenu()
	} else if menu == "exit" {
		app.chooseExitMenu()
	}
}

func (app *App) chooseMainMenu() {

	switch app.option {
	case "1":
		app.menu = "game"
	case "2":
		app.menu = "login"
	case "3":
		app.menu = "leaderboard"
	case "4":
		app.menu = "exit"
	default:

	}
}

func (app *App) chooseGameMenu() {
	switch app.option {
	case "1":
		app.menu = "game"
	case "2":
		app.menu = "main"
	default:

	}
}

func (app *App) chooseLoginMenu() {
	switch app.option {
	case "1":
		app.menu = "login"
	case "2":
		app.menu = "login"
	case "3":
		app.menu = "main"
	default:

	}
}

func (app *App) chooseLeaderboardMenu() {
	switch app.option {
	case "1":
		app.menu = "main"
	default:

	}
}

func (app *App) chooseExitMenu() {
	switch app.option {
	case "1":
		app.menu = "main"
	case "2":
		app.finished = true
	default:

	}
}

func (app *App) Init() {
	app.menu = "main"
	for !app.finished {
		app.printMenu()
		fmt.Print("Ingrese una opcion: ")
		fmt.Scanln(&app.option)
		app.chooseOption()
	}
}
