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
	g.PlayerName = getUserEntry("Enter your name: ")

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

func (g *GameState) PlayQuiz() {
	for i, question := range g.Questions {
		fmt.Printf("\033[33m %d. %s \033[0m\n", i+1, question.Text)

		for j, option := range question.Options {
			fmt.Printf("[%d]. %s\n", j+1, option)
		}

		for {
			answer := toInt(getUserEntry("Type your answer:  "))

			if answer < 1 || answer > len(question.Options) {
				fmt.Println("Invalid answer. Please try again.")
				continue
			}

			if answer == question.Answer {
				g.Score += 10
				fmt.Printf("\033[32mCorrect!\033[0m\n")
			} else {
				fmt.Printf("\033[31mWrong!\033[0m\n")
			}

			fmt.Printf("------------------------- \n")
			break
		}
	}
}

func main() {
	game := GameState{}
	game.LoadGameData()
	game.Init()
	game.PlayQuiz()

	fmt.Println("Game Over!!!")
	fmt.Printf("Your final score: %d\n", game.Score)
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return 0
	}
	return i
}

func getUserEntry(title string) string {
	fmt.Println(title)
	reader := bufio.NewReader(os.Stdin)

	value, err := reader.ReadString('\n') // read until newline (enter)
	if err != nil {
		panic("Error reading user entry")
	}

	return value[:len(value)-1] // remove newline character
}
