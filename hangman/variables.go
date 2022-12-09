package hangman

var hangmanMenuOptions = [...]string{"Go back to main menu", "Exit"}

var menus = map[string][]string{
	"hangman": hangmanMenuOptions[:],
}

var options = map[string]func(app *App){
	"hangman": (*App).chooseHangmanMenu,
}

var vocabulary [][]string

var vocals = []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

var affirmative = []string{"yes", "YES", "y", "Y"}
var negative = []string{"no", "NO", "n", "N"}

type hangman struct {
	livesLeft int
	score     float32
	word      string
	match     map[string]bool
	win       bool
}

var lives = map[int]string{
	1: " 0\n",
	2: " 0\n |",
	3: " 0\n/|",
	4: " 0\n/|\\\n",
	5: " 0\n/|\\\n/ ",
	6: " 0\n/|\\\n/ \\",
}
