package app

func (app *App) Init() {
	app.title = "Hangman Game"
	app.menu = "main"
	app.user = "Nicolas"
	app.mode = "menu"
	for !app.finished {
		app.processMode()
	}
}
