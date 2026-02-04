package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bb2-b/word-search-factory/board/src/board"
)

func main() {
	// Get filename from standard input.
	if len(os.Args) != 2 {
		fmt.Println("one file argument containing a list words separated by a newline must be provided")
		return
	}
	fileName := os.Args[1]

	// Extract words from the provided file.
	words, err := getWordsFromFile(fileName)
	if err != nil {
		fmt.Printf("error getting words from file: %s\n", err)
		return
	}
	var wordList []string
	for word := range words {
		wordList = append(wordList, word)
	}

	// Create the game.
	game, err := board.NewGameBoard(wordList, strings.Split(fileName, "-")[0])
	if err != nil {
		fmt.Printf("error generating new game board: %s\n", err)
		return
	}

	game.PrettyPrintGameBoard(game.Grid(), false)
	game.PrettyPrintGameBoard(game.AnswerKey(), false)
}

func getWordsFromFile(file string) (map[string]bool, error) {
	f, err := os.Open("example-word-lists/" + file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error during scanning: %e", err)
	}

	if len(lines) == 0 {
		return nil, fmt.Errorf("provided empty file")
	}
	words := make(map[string]bool, len(lines))
	for _, word := range lines {
		words[word] = true
	}

	return words, nil
}
