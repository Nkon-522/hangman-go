package main

import "hangman/app"

var mainMenuOptions = [...]string{"Play", "Login", "Leaderboard", "Exit"}
var gameMenuOptions = [...]string{"Start Game", "Return"}
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

var options = map[string]func(app *app.App){
	"main":        (*App).chooseMainMenu,
	"game":        (*App).chooseGameMenu,
	"login":       (*App).chooseLoginMenu,
	"leaderboard": (*App).chooseLeaderboardMenu,
	"exit":        (*App).chooseExitMenu,
}
