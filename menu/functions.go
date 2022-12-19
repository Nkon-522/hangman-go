package menu

import (
	"fmt"
	"hangman/connection"
)

func (menu *Menu) chooseMainMenu() {

	switch *menu.option {
	case "1":
		*menu.menu = "game"
	case "2":
		*menu.menu = "login"
	case "3":
		*menu.menu = "leaderboard"
	case "4":
		*menu.menu = "exit"
	default:

	}
}

func (menu *Menu) chooseGameMenu() {
	switch *menu.option {
	case "1":
		*menu.mode = "hangman"
	case "2":
		*menu.menu = "main"
	default:

	}
}

func (menu *Menu) chooseHangmanMenu() {
	switch *menu.option {
	case "1":
		*menu.mode = "hangman"
	case "2":
		*menu.menu = "main"
	case "3":
		*menu.finished = true
	default:

	}
}

func (menu *Menu) chooseLoginMenu() {
	switch *menu.option {
	case "1":
		*menu.menu = "login"
	case "2":
		*menu.menu = "login"
	case "3":
		*menu.menu = "main"
	default:

	}
}

func (menu *Menu) chooseLeaderboardMenu() {
	switch *menu.option {
	case "1":
		*menu.menu = "main"
	default:

	}
}

func (menu *Menu) chooseExitMenu() {
	switch *menu.option {
	case "1":
		*menu.menu = "main"
	case "2":
		*menu.finished = true
	default:

	}
}

func InitMenu(menuConnection *connection.MenuConnection) {
	menu := Menu{
		title:    menuConnection.Title,
		menu:     menuConnection.Menu,
		mode:     menuConnection.Mode,
		option:   menuConnection.Option,
		finished: menuConnection.Finished,
	}
	menu.printMenu()
	fmt.Print("Ingrese una opcion: ")
	fmt.Scanln(menu.option)
	menu.chooseOption()
}
