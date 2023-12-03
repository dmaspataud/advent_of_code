package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	ID       int
	maxBlue  int
	maxRed   int
	maxGreen int
}

func main() {
	f, err := os.ReadFile("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve(string(f)))
}

func solve(text string) int {
	games := parseHands(text)
	res := 0

	for _, game := range games {
		power := game.maxBlue * game.maxGreen * game.maxRed
		res += power
	}

	return res
}

func parseHands(text string) []Game {
	games := []Game{}
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		thisGame := Game{}
		line := scanner.Text()

		// populate game ID
		_, err := fmt.Sscanf(line, "Game %v:", &thisGame.ID)
		if err != nil {
			log.Fatal(err)
		}

		_, hands, _ := strings.Cut(line, ":")
		handList := strings.Split(hands, ";")

		for _, hand := range handList {

			for _, color := range strings.Split(hand, ",") {

				color = color[1:] // get rid of the first space

				n, c, _ := strings.Cut(color, " ")

				num, err := strconv.Atoi(n)
				if err != nil {
					log.Fatal(err)
				}
				if c == "blue" {
					thisGame.maxBlue = max(thisGame.maxBlue, num)
				} else if c == "red" {
					thisGame.maxRed = max(thisGame.maxRed, num)
				} else if c == "green" {
					thisGame.maxGreen = max(thisGame.maxGreen, num)
				}
			}
		}

		games = append(games, thisGame)
	}
	return games
}
