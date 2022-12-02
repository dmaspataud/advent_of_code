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
		"X": 1, // Rock
		"B": 2, // Paper
		"Y": 2, // Paper
		"C": 3, // Scissors
		"Z": 3, // Scissors
	}
	score := 0
	roundScore := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		round := strings.Split(scanner.Text(), " ")
		roundScore = symbols[round[1]]
		if round[1] == "X" && round[0] == "C" || round[1] == "Y" && round[0] == "A" || round[1] == "Z" && round[0] == "B" { // win
			roundScore += 6
		} else if symbols[round[1]] == symbols[round[0]] { // draw
			roundScore += 3
		} // else it's a loss, we skip adding 0
		score += roundScore
	}
	fmt.Println(score)
}
