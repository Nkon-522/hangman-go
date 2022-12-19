package menu

import (
	"fmt"
	"hangman/terminal"
	"log"
)

func (menu *Menu) printMenuOptions() {
	for index, option := range menus[*menu.menu] {
		fmt.Printf("%d. %s\n", index+1, option)
	}
}

func (menu *Menu) printMenu() {
	terminal.CallClear()

	fmt.Println(*menu.title)
	fmt.Println("------------")

	menu.printMenuOptions()
}

func (menu *Menu) chooseOption() {
	option, ok := options[*menu.menu]
	if !ok {
		log.Fatal("Something went wrong!")
	}
	option(menu)
}
