package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// get data
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// hashmap with symbols and values
	symbols := map[string]int{
		"A": 1, // Rock
		"B": 2, // Paper
		"C": 3, // Scissors
	}

	loseMap := map[string]int{
		"A": 3, // to lose, play scissors
		"B": 1, // to lose, play rock
		"C": 2, // to lose, play paper
	}

	winMap := map[string]int{
		"A": 2, // to win play Paper
		"B": 3, // to win play Scissors
		"C": 1, // to win play Rock
	}

	score := 0
	roundScore := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		round := strings.Split(scanner.Text(), " ")
		switch {
		case round[1] == "X": // lose
			roundScore = loseMap[round[0]]
		case round[1] == "Y": // draw
			roundScore = symbols[round[0]]
			roundScore += 3
		case round[1] == "Z": // win
			roundScore = winMap[round[0]]
			roundScore += 6
		}
		score += roundScore
	}
	fmt.Println(score)
}
