package app

var mainMenuOptions = [...]string{"Play", "Login", "Leaderboard", "Exit"}

var gameMenuOptions = [...]string{"Do Something", "Return"}
var loginMenuOptions = [...]string{"Login", "Sign up", "Return"}
var leaderboardMenuOptions = [...]string{"Return"}
var exitMenuOptions = [...]string{"Return", "Exit"}

type App struct {
	menu     string
	option   string
	finished bool
}

type Menus interface {
	printMenu()
	chooseOption()

	chooseMainMenu()

	chooseGameMenu()
	chooseLoginMenu()
	chooseLeaderboardMenu()
	chooseExitMenu()
}
