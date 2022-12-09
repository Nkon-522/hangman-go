package main

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
		hangman.Init(&app.user, &app.option, &app.finished, &app.user)
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
