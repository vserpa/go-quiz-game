package main

import "fmt"

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

func main() {
	fmt.Println("Hello, World!")
}
