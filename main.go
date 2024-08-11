package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

var config Config

func main() {
	fmt.Println("HANGMAN game 2024")

	err := readConfig()
	if err != nil {
		panic(err)
	}

	categories := extractCategories(config.Categories)

	var category string
	prompt := &survey.Select{
		Message: "Pick a category to receive a random challenge",
		Options: categories,
	}

	err = survey.AskOne(prompt, &category)
	if err != nil {
		fmt.Printf("Error during picking category. Error: %s\n", err.Error())
		panic(err)
	}

	fmt.Println("Getting random challenge from category:", category)

	challenge := generateChallenge(category)
	guessedLetters := make(map[rune]bool)
	wrongGuesses := 0

	// Game loop
	for wrongGuesses < config.GuessLimit {
		// Display current word state
		fmt.Printf("Word: ")
		for _, s := range challenge {
			if guessedLetters[s] {
				fmt.Printf("%c ", s)
			} else {
				fmt.Printf("_ ")
			}
		}
		fmt.Println()
		fmt.Println("Wrong guesses:", wrongGuesses, "Guesses left:", config.GuessLimit-wrongGuesses)

		var guessStr string
		prompt := &survey.Input{
			Message: "Guess a letter:",
		}

		err = survey.AskOne(prompt, &guessStr)
		if err != nil {
			fmt.Printf("Error during guess parse. Error: %s\n", err.Error())
			panic(err)
		}
		if _, err = strconv.ParseInt(guessStr, 10, 32); err == nil {
			fmt.Println("Error. Expected character. Got number.")
			continue
		}

		if len(guessStr) != 1 {
			fmt.Printf("Error. Expected single character. Got %s\n", guessStr)
			continue
		}

		guess := rune(guessStr[0])
		if guessedLetters[guess] {
			fmt.Printf("You already guessed the letter %c.\n", guess)
			continue
		}

		guessedLetters[guess] = true

		// Check if the guess is in the challenge
		if strings.ContainsRune(challenge, guess) {
			fmt.Println("Correct guess!")
		} else {
			fmt.Println("Wrong guess.")
			wrongGuesses++
		}

		// Check if the game is won
		guessedWord := ""
		for _, s := range challenge {
			if guessedLetters[s] {
				guessedWord += string(s)
			} else {
				guessedWord += "_"
			}
		}

		if guessedWord == challenge {
			fmt.Println("Congratulations! You guessed the word:", challenge)
			return
		}
	}

	fmt.Println("Game Over. The word was:", challenge)
}

func generateChallenge(category string) string {
	switch category {
	case "Animals":
		return config.Categories.Animals[rand.Intn(len(config.Categories.Animals))]
	case "Food":
		return config.Categories.Food[rand.Intn(len(config.Categories.Food))]
	case "Countries":
		return config.Categories.Countries[rand.Intn(len(config.Categories.Countries))]
	default:
		return ""
	}
}

func readConfig() error {
	b, err := os.ReadFile("settings.json")
	if err != nil {
		fmt.Printf("Error during opening config. Error: %s\n", err.Error())
		return err
	}

	err = json.Unmarshal(b, &config)
	if err != nil {
		fmt.Printf("Error during reading config. Error: %s\n", err.Error())
		return err
	}

	return nil
}

func extractCategories(s interface{}) []string {
	v := reflect.ValueOf(s)
	t := v.Type()
	var fieldNames []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldNames = append(fieldNames, field.Name)
	}

	return fieldNames
}

type Config struct {
	GuessLimit int             `json:"guessLimit"`
	Categories HangmanCategory `json:"categories"`
}

type HangmanCategory struct {
	Animals   []string `json:"animals"`
	Food      []string `json:"food"`
	Countries []string `json:"countries"`
}
