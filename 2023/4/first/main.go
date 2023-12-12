package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type card struct {
	winningNumbers map[int]bool
	cardNumbers    map[int]bool
	ID             int
}

func main() {
	f, err := os.ReadFile("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve(string(f)))
}

func solve(text string) int {
	res := 0
	cards := parseCards(text)

	for _, card := range cards {
		cardRes := 0
		for cardNumber := range card.cardNumbers {
			if _, ok := card.winningNumbers[cardNumber]; ok {
				if cardRes == 0 {
					cardRes += 1
				} else {
					cardRes *= 2
				}
			}
		}
		res += cardRes
	}

	return res
}

func parseCards(text string) []card {
	cards := []card{}

	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		line := scanner.Text()
		thisCard := card{}
		var err error

		thisCard.ID, err = strconv.Atoi(strings.TrimSpace(strings.Split(strings.Split(line, ":")[0], "d")[1]))
		if err != nil {
			log.Fatal("ID", err)
		}
		winningNumbers := convertNumbers(strings.Split(strings.Split(strings.Split(line, ":")[1], "|")[0], " "))
		cardNumbers := convertNumbers(strings.Split(strings.Split(strings.Split(line, ":")[1], "|")[1], " "))

		cards = append(cards, card{ID: thisCard.ID, winningNumbers: winningNumbers, cardNumbers: cardNumbers})
	}
	return cards
}

func convertNumbers(numbers []string) map[int]bool {
	res := make(map[int]bool)

	for _, numStr := range numbers {
		cleanNumStr := strings.TrimSpace(numStr)

		if cleanNumStr == "" {
			continue
		}

		num, err := strconv.Atoi(cleanNumStr)
		if err != nil {
			log.Printf("Error converting %s to int: %v", cleanNumStr, err)
			continue
		}

		res[num] = true
	}

	return res
}
