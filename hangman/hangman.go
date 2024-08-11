package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

// wrongGuessesAllowed is the number of wrong guesses allowed before the player loses
const wrongGuessesAllowed = 10

// alphabet contains the letters of the alphabet
const alphabet = "abcdefghijklmnopqrstuvwxyz"

// filePaths contains the paths to the files containing the words
var filePaths = map[string]string{
	"animals":   "animals.txt",
	"food":      "food.txt",
	"countries": "countries.txt",
}

// Start starts the hangman game
func Start() {
	// Seed the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println("Welcome to Hangman game!")
	fmt.Println("____________")
	fmt.Println()

	// Load words from files
	words := loadWords()
	// Select a category
	selectedCategory := selectCategory(words)
	// Select a random word from the chosen category
	selectedWord := words.selectWord(selectedCategory)

	guesses := make([]string, len(selectedWord))
	for i := range guesses {
		guesses[i] = "_"
	}
	fmt.Println(guesses)

	wrongGuesses := 0
	for {
		clearScreen()
		printHeader()

		drawHangman(wrongGuesses - 1)

		fmt.Println()
		fmt.Println("The selected category is the ", selectedCategory)
		fmt.Println()
		fmt.Printf("Word: %s\n", strings.Join(guesses, " "))
		fmt.Printf("Wrong guesses: %d\n", wrongGuesses)
		fmt.Print("Guess a letter: ")

		if wrongGuesses >= wrongGuessesAllowed {
			fmt.Println()
			fmt.Println("GAME OVER!")
			fmt.Printf("Sorry, you've lost! The word was: %s\n", selectedWord)
			break
		}

		reader := bufio.NewReader(os.Stdin)
		input, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		guess := strings.TrimSpace(strings.ToLower(string(input)))
		if len(guess) != 1 || !strings.Contains(alphabet, guess) {
			wrongGuesses++
			continue
		}

		if strings.Contains(selectedWord, guess) {
			for i, char := range selectedWord {
				if string(char) == guess {
					guesses[i] = guess
				}
			}

			if !strings.Contains(strings.Join(guesses, ""), "_") {
				fmt.Printf("\nThe word was: %s\n", selectedWord)
				fmt.Println("Congratulations, you've won!")
				break
			}
		} else {
			wrongGuesses++
		}

	}
}

// loadWords creates a map of categories to a list of words
func loadWords() (words Words) {
	words = make(Words)
	for category, filePath := range filePaths {
		wordList, err := readWordsFromFile(filePath)
		if err != nil {
			fmt.Printf("Error reading from %s: %v\n", filePath, err)
			return nil
		}
		words[category] = wordList
	}
	return words
}

type Words map[string][]string

// selectWord selects a random word from the chosen category
func (w Words) selectWord(category string) (selectedWord string) {
	if len(w[category]) > 0 {
		selectedWord = w[category][rand.Intn(len(w[category]))]
		selectedWord = strings.ToLower(selectedWord)
	} else {
		fmt.Println("The selected category does not contain any words.")
		return ""
	}
	return selectedWord
}

// selectCategory prompts the user to select a category
func selectCategory(words Words) (selectedCategory string) {
	var categoryNames []string

	for category := range words {
		categoryNames = append(categoryNames, category)
	}

	sort.SliceStable(categoryNames, func(i, j int) bool {
		return categoryNames[i] < categoryNames[j]
	})

	for i, category := range categoryNames {
		fmt.Printf("%d.) %s\n", i+1, category)
	}

	fmt.Println("Select a category by number!")
	fmt.Printf("Enter a number between 1 and %d: ", len(categoryNames))

	reader := bufio.NewReader(os.Stdin)
	categoryInput, _ := reader.ReadString('\n')
	categoryInput = strings.TrimSpace(categoryInput)
	inputNum, err := strconv.Atoi(categoryInput)
	if err != nil || inputNum < 1 || inputNum > len(categoryNames) {
		fmt.Println("Invalid input. Please enter a valid number.")
		return ""
	}

	// Adjust for zero-based indexing
	selectedCategory = categoryNames[inputNum-1]
	fmt.Printf("Selected category: %s\n", selectedCategory)

	return selectedCategory
}

// readWordsFromFile reads words from a file and returns a slice of strings
func readWordsFromFile(filePath string) (words []string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

// printHeader prints the header of the game
func printHeader() {
	fmt.Println("=============")
	fmt.Println("HANGMAN GAME")
	fmt.Println("=============")
}

// stages contains the ASCII art for the hangman stages
var stages = []string{
	`





=========`,
	`
      |
      |
      |
      |
      |
=========`, `
  +---+
      |
      |
      |
      |
      |
=========`,
	`
  +---+
  |   |
      |
      |
      |
      |
=========`, `
  +---+
  |   |
  O   |
      |
      |
      |
=========`, `
  +---+
  |   |
  O   |
  |   |
      |
      |
=========`, `
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========`, `
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========`, `
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========`, `
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========`,
}

// drawHangman prints the hangman stage corresponding to the number of wrong guesses
func drawHangman(wrongGuesses int) {
	if wrongGuesses > 0 {
		fmt.Println(stages[wrongGuesses])
	}
}

// clearScreen clears the terminal screen
func clearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error clearing screen:", err)
		return
	}
}
