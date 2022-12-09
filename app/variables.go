package app

var mainMenuOptions = [...]string{"Play", "Login", "Leaderboard", "Exit"}
var gameMenuOptions = [...]string{"Start Game", "Return"}
var hangmanMenuOptions = [...]string{"Play again", "Go back to main menu", "Exit"}
var loginMenuOptions = [...]string{"Login", "Sign up", "Return"}
var leaderboardMenuOptions = [...]string{"Return"}
var exitMenuOptions = [...]string{"Return", "Exit"}

var menus = map[string][]string{
	"main":        mainMenuOptions[:],
	"game":        gameMenuOptions[:],
	"hangman":     hangmanMenuOptions[:],
	"login":       loginMenuOptions[:],
	"leaderboard": leaderboardMenuOptions[:],
	"exit":        exitMenuOptions[:],
}

var options = map[string]func(app *App){
	"main":        (*App).chooseMainMenu,
	"game":        (*App).chooseGameMenu,
	"hangman":     (*App).chooseHangmanMenu,
	"login":       (*App).chooseLoginMenu,
	"leaderboard": (*App).chooseLeaderboardMenu,
	"exit":        (*App).chooseExitMenu,
}

type App struct {
	title    string
	menu     string
	option   string
	finished bool
	user     string
}
