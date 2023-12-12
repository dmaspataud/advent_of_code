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
	count          int
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

	// go through cards and see how many matches we have.
	for c := 1; c < len(cards)+1; c++ {
		card := cards[c]
		log.Printf("card %v, count %v\n", card.ID, card.count)
		for i := 0; i < card.count; i++ {
			res += 1
			matches := 0
			for cardNumber := range card.cardNumbers {
				if _, ok := card.winningNumbers[cardNumber]; ok {
					matches += 1
				}
			}

			// add the new cards
			if matches > 0 {
				for i := 1; i <= matches; i++ {
					log.Printf("Card ID %v - adding 1 card for card ID: %v", card.ID, cards[c+i].ID)
					currentCard := cards[c+i]
					currentCard.count += 1
					cards[c+i] = currentCard
				}
			}
		}
	}
	return res
}

func parseCards(text string) map[int]card {
	cards := make(map[int]card)

	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		line := scanner.Text()
		thisCard := NewCard()
		var err error

		thisCard.ID, err = strconv.Atoi(strings.TrimSpace(strings.Split(strings.Split(line, ":")[0], "d")[1]))
		if err != nil {
			log.Fatal("ID", err)
		}
		winningNumbers := convertNumbers(strings.Split(strings.Split(strings.Split(line, ":")[1], "|")[0], " "))
		cardNumbers := convertNumbers(strings.Split(strings.Split(strings.Split(line, ":")[1], "|")[1], " "))

		cards[thisCard.ID] = card{ID: thisCard.ID, winningNumbers: winningNumbers, cardNumbers: cardNumbers, count: thisCard.count}
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

func NewCard() card {
	return card{
		winningNumbers: make(map[int]bool),
		cardNumbers:    make(map[int]bool),
		count:          1,
	}
}
