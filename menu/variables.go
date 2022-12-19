package menu

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

var options = map[string]func(menu *Menu){
	"main":        (*Menu).chooseMainMenu,
	"game":        (*Menu).chooseGameMenu,
	"hangman":     (*Menu).chooseHangmanMenu,
	"login":       (*Menu).chooseLoginMenu,
	"leaderboard": (*Menu).chooseLeaderboardMenu,
	"exit":        (*Menu).chooseExitMenu,
}

type Menu struct {
	title    *string
	menu     *string
	mode     *string
	option   *string
	finished *bool
}
