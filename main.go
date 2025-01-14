package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	game := GameState{}
	game.Init()
}
