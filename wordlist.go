package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"math/rand/v2"
	"os"
	"strings"
)

type WordCategory struct {
	Name  string
	Words []Word
}

func ReadCategory(reader *bufio.Reader) (WordCategory, error) {
	name := ReadLine(reader)
	words := ReadLine(reader)
	ReadLine(reader)

	if name == "" || words == "" {
		return WordCategory{
			Name:  "",
			Words: []Word{},
		}, errors.New("could not read")
	}

	var wordList []Word = make([]Word, 0)
	for _, s := range strings.Split(words, ": ") {
		wordList = append(wordList, Word(strings.TrimSpace(s)))
	}

	return WordCategory{
		Name:  name,
		Words: wordList,
	}, nil
}

func ReadLine(reader *bufio.Reader) string {
	line, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			return ""
		}

		log.Fatalf("read file line error: %v", err)
	}
	return line
}

func ReadFile(fileName string) []WordCategory {
	f, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("Open word list file failed")
	}
	defer f.Close()

	rd := bufio.NewReader(f)

	var categories []WordCategory = make([]WordCategory, 0)

	for {
		category, err := ReadCategory(rd)
		if err != nil {
			return categories
		}
		categories = append(categories, category)
	}
}

func SelectWord(categories []WordCategory) (string, Word) {
	categoryCount := len(categories)
	categoryIndex := rand.IntN(categoryCount)

	category := categories[categoryIndex]
	wordCount := len(category.Words)
	selectedWordIndex := rand.IntN(wordCount)
	return category.Name, category.Words[selectedWordIndex]
}
