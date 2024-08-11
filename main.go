package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

const WORDS_FILE_NAME = "words.txt"
const MAXIMUM_NUMBER_OF_GUESSES = 10
const CTRL_C rune = '\003'

func Print(status GameStatus, result GameResult) {
	statusLine := status.PrintWord()

	var resultText string
	switch result {
	case WonGame:
		resultText = Green("You Won!")
	case LostGame:
		resultText = Yellow("You Lost")
	default:
		resultText = "Take a guess: "
	}

	text := []string{
		"",
		fmt.Sprintf("Word: %s", statusLine),
		fmt.Sprintf("Wrong guesses: %d", status.WrongGuesses),
		resultText,
	}
	fmt.Print(strings.Join(text, "\r\n"))
}

func main() {

	fileName := WORDS_FILE_NAME
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	defer fmt.Printf("\r\n")

	wordFile := ReadFile(fileName)
	category, word := SelectWord((wordFile))

	game := NewGame(word, MAXIMUM_NUMBER_OF_GUESSES)

	PrintMessage(fmt.Sprintf("Category: %s", category), Yellow)
	PrintMessage("Start quessing!", Red)
	for {
		Print(game, OnGoingGame)

		b := make([]byte, 1)
		_, err = os.Stdin.Read(b)
		if err != nil {
			fmt.Println(err)
			return
		}
		char := rune(b[0])

		if char == CTRL_C {
			os.Exit(-1)
		}

		PrintMessage(string(char), Blue)

		letter, err := ToLetter(char)
		if err != nil {
			PrintMessage("Not a letter", Red)
			continue
		}

		result, err := game.Guess(letter)
		if err != nil {
			PrintMessage(err.Error(), Red)
			continue
		}

		if result != OnGoingGame {
			Print(game, result)
			break
		}
	}
}
