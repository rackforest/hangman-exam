package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var initialStatus = GameStatus{
	SelectedWord: "Example",
	Letters: map[Letter]LetterState{
		'E': Unrevealed,
		'X': Unrevealed,
		'A': Unrevealed,
		'M': Unrevealed,
		'P': Unrevealed,
		'L': Unrevealed,
	}, WrongGuesses: 0,
	MaximumWrongGuesses: 3,
	UniqueLetters:       6,
	RevealedLetters:     0,
}

var inGameStatus = GameStatus{
	SelectedWord:        "Data",
	Letters:             map[Letter]LetterState{'A': Revealed},
	WrongGuesses:        1,
	MaximumWrongGuesses: 0,
	UniqueLetters:       3,
	RevealedLetters:     1,
}

var endGameStatus = GameStatus{
	SelectedWord:        "Elephant",
	WrongGuesses:        1,
	MaximumWrongGuesses: 20,
	UniqueLetters:       7,
	RevealedLetters:     7,
	Letters: map[Letter]LetterState{
		'E': Revealed,
		'L': Revealed,
		'P': Revealed,
		'H': Revealed,
		'A': Revealed,
		'N': Revealed,
		'T': Revealed,
	},
}

func TestNewStatus(t *testing.T) {
	status := NewGame("Example", 3)

	assert.Equal(t, initialStatus, status)
}

func TestInitialPrintStatus(t *testing.T) {
	printed := initialStatus.PrintWord()

	assert.Equal(t, "_ _ _ _ _ _ _", printed)
}
func TestInGamePrintStatus(t *testing.T) {
	printed := inGameStatus.PrintWord()

	assert.Equal(t, "_ a _ a", printed)
}
func TestEndGamePrintStatus(t *testing.T) {
	printed := endGameStatus.PrintWord()

	assert.Equal(t, "E l e p h a n t", printed)
}
