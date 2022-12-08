package app

var mainMenuOptions = [...]string{"Play", "Login", "Leaderboard", "Exit"}

var gameMenuOptions = [...]string{"Do Something", "Return"}
var loginMenuOptions = [...]string{"Login", "Sign up", "Return"}
var leaderboardMenuOptions = [...]string{"Return"}
var exitMenuOptions = [...]string{"Return", "Exit"}

var menus = map[string][]string{
	"main":        mainMenuOptions[:],
	"game":        gameMenuOptions[:],
	"login":       loginMenuOptions[:],
	"leaderboard": leaderboardMenuOptions[:],
	"exit":        exitMenuOptions[:],
}

var options = map[string]func(app *App){
	"main":        (*App).chooseMainMenu,
	"game":        (*App).chooseGameMenu,
	"login":       (*App).chooseLoginMenu,
	"leaderboard": (*App).chooseLeaderboardMenu,
	"exit":        (*App).chooseExitMenu,
}

type App struct {
	menu     string
	option   string
	finished bool
}
