package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type GameState struct {
	PlayerName string
	Score      int
	Questions  []Question
}

type Question struct {
	Text    string
	Options []string
	Answer  int
}

func (g *GameState) Init() {
	fmt.Println("Initializing game state...")
	fmt.Println("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n') // read until newline (enter)

	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}

	g.PlayerName = name[:len(name)-1] // remove newline character

	fmt.Printf("Welcome, %s! Let's start the quiz.\n", g.PlayerName)
}

func (g *GameState) LoadGameData() {
	file, err := os.Open("quiz.csv")
	if err != nil {
		fmt.Println("Error opening quiz file:", err)
		return
	}
	defer file.Close() // needs to close the file when we're done

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	for index, record := range records {
		if index > 0 {
			question := Question{
				Text:    record[0],
				Options: record[1:5],
				Answer:  toInt(record[5]),
			}

			g.Questions = append(g.Questions, question)
		}
	}
}

func main() {
	game := GameState{}
	go game.LoadGameData()
	game.Init()
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return 0
	}
	return i
}
