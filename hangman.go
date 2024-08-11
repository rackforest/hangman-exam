package main

import (
	"errors"
	"strings"
	"unicode"
)

type Word string

type Letter rune

type LetterState int

const (
	Unrevealed LetterState = iota
	Revealed
	Missed
)

type LetterStateMap map[Letter]LetterState

type GameStatus struct {
	SelectedWord        Word
	Letters             LetterStateMap
	WrongGuesses        int
	MaximumWrongGuesses int
	UniqueLetters       int
	RevealedLetters     int
}

type GameResult int

const (
	OnGoingGame GameResult = iota
	WonGame
	LostGame
)

func ToLetter(r rune) (Letter, error) {
	var e error = nil
	if !unicode.IsLetter(r) {
		e = errors.New("not a letter")
	}
	return Letter(unicode.ToUpper(r)), e
}

func MakeLetterStateMap(word Word) (LetterStateMap, int) {
	letters := make(LetterStateMap)
	nrOfUniqueLetters := 0

	for _, c := range word {
		letter, _ := ToLetter(c)
		_, knownLetter := letters[letter]
		if !knownLetter {
			nrOfUniqueLetters++
			letters[letter] = Unrevealed
		}
	}

	return letters, nrOfUniqueLetters
}

func NewGame(word Word, maximumWrongGuesses int) GameStatus {
	letters, nrOfUniqueLetters := MakeLetterStateMap(word)

	return GameStatus{
		SelectedWord:        word,
		Letters:             letters,
		WrongGuesses:        0,
		MaximumWrongGuesses: maximumWrongGuesses,
		UniqueLetters:       nrOfUniqueLetters,
		RevealedLetters:     0,
	}
}

func (status *GameStatus) Evaluate() GameResult {
	if status.RevealedLetters >= status.UniqueLetters {
		return WonGame
	}

	if status.WrongGuesses >= status.MaximumWrongGuesses {
		return LostGame
	}

	return OnGoingGame
}

func (status *GameStatus) Guess(guess Letter) (GameResult, error) {
	letterStatus, found := status.Letters[guess]

	if !found {
		status.Letters[guess] = Missed
		status.WrongGuesses++
		return status.Evaluate(), nil
	}

	if letterStatus == Revealed || letterStatus == Missed {
		return status.Evaluate(), errors.New("repeated letter")
	}

	status.Letters[guess] = Revealed
	status.RevealedLetters++

	return status.Evaluate(), nil
}

func (status *GameStatus) PrintWord() string {
	letters := []string{}
	for _, c := range status.SelectedWord {
		var toPrint rune
		letter, _ := ToLetter(c)
		if status.Letters[Letter(letter)] == Revealed {
			toPrint = c
		} else {
			toPrint = '_'
		}
		letters = append(letters, string(toPrint))
	}

	return strings.Join(letters, " ")
}
