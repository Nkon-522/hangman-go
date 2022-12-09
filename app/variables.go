package app

type App struct {
	title    string
	menu     string
	option   string
	finished bool
	user     string

	menus   map[string][]string
	options map[string]func(app *App)
}
