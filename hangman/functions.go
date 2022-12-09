package hangman

import (
	"bufio"
	"fmt"
	"hangman/helpers"
	"hangman/terminal"
	"os"
	"strings"
)

func init() {
	vocabulary = helpers.ReadCsvFile("hangman/nounlist.csv")
}

func Start_game(user string) {
	hangman := NewHangman()

	for hangman.livesLeft > 0 && !hangman.win {
		terminal.CallClear()
		hangman.printHangman()
		hangman.printScore()
		hangman.printWord()
		hangman.printLettersChose()
		hangman.choose()
	}

	terminal.CallClear()
	hangman.handleGameOver(user)
}

func NewHangman() *hangman {
	hangman := new(hangman)

	hangman.livesLeft = 6
	hangman.word = helpers.PickRandom(&vocabulary)
	hangman.score = 0
	hangman.win = false
	if hangman.match == nil {
		hangman.match = make(map[string]bool)
	}
	for _, letter := range hangman.word {
		hangman.match[string(letter)] = false
	}

	return hangman
}

func (hangman *hangman) printHangman() {
	fmt.Println(lives[hangman.livesLeft])
}

func (hangman *hangman) printScore() {
	fmt.Println("Score: ", hangman.score)
}

func (hangman *hangman) printWord() {
	for _, letter := range hangman.word {
		if hangman.match[string(letter)] {
			fmt.Print(string(letter), " ")
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Print("\n")
}

func (hangman *hangman) printLettersChose() {
	fmt.Print("Letters chose: ")
	for letter, chose := range hangman.match {
		if chose {
			fmt.Print(letter)
		}
	}
	fmt.Print("\n")
}

func (hangman *hangman) choose() {
	var letter string
	for {
		fmt.Print("Choose a letter: ")
		fmt.Scan(&letter)
		fmt.Print("\n")
		if picked, exists := hangman.match[letter]; !exists || (exists && !picked) {
			hangman.match[letter] = true
			break
		} else {
			fmt.Println("You have already picked this letter!")
		}
	}
	if !strings.Contains(hangman.word, letter) {
		hangman.livesLeft--
	} else {
		hangman.updateScore(letter)
	}
}

func (hangman *hangman) updateScore(letter string) {
	if helpers.StringInSlice(letter, vocals) {
		hangman.score = hangman.score + 1
	} else {
		hangman.score = hangman.score + 1.5
	}

}

func (hangman *hangman) printGameOver() {
	fmt.Println("The word was:", hangman.word)
	if hangman.win {
		fmt.Println("You won!!")
	} else {
		fmt.Println("You lost!")
	}
	fmt.Println("Your score:", hangman.score)
	fmt.Print("Press ENTER to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func (hangman *hangman) handleGameOver(user string) {
	hangman.printGameOver()

	var save string

	for {
		fmt.Println("Do you wish to save your results", user, "?")
		fmt.Scan(&save)
		if helpers.StringInSlice(save, affirmative) {
			hangman.saveResults(user)
			break
		} else if helpers.StringInSlice(save, negative) {
			break
		} else {
			fmt.Println("Please choose one of the two options!!")
		}
	}

}

func (hangman *hangman) saveResults(user string) {
	fmt.Println("Now I should be saving but it is not implemented yet...", user)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
