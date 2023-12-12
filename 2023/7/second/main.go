package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var ruleSet = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 1,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

type hand struct {
	cards    string
	bid      int
	handType int
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

	t := parseHands(text)
	for i := range t {
		t[i].Evaluate()
	}

	Rank(t)

	for k, v := range t {
		res += (v.bid * (k + 1))
	}

	return res
}

func parseHands(text string) []hand {
	scanner := bufio.NewScanner(strings.NewReader(text))

	allHands := []hand{}
	for scanner.Scan() {
		line := scanner.Text()
		cards := strings.Split(line, " ")[0]
		bid, err := strconv.Atoi(strings.Split(line, " ")[1])

		if err != nil {
			log.Fatal(err)
		}

		allHands = append(allHands, hand{cards: cards, bid: bid})
	}

	return allHands
}

func (h *hand) Evaluate() {
	double, triple, quadruple, quintuple := 0, 0, 0, 0

	// count cards to determine which type of hand we have, fill cardType
	cardMap := make(map[string]int)
	for _, card := range h.cards {
		cardMap[string(card)] += 1
	}

	for _, v := range cardMap {
		switch v {
		case 2:
			double++
		case 3:
			triple++
		case 4:
			quadruple++
		case 5:
			quintuple++
		}
	}

	switch {
	case quintuple == 1:
		h.handType = 7 // five of a kind
	case quadruple == 1:
		if cardMap["J"] == 1 || cardMap["J"] == 4 {
			h.handType = 7 // five of a kind
		} else {
			h.handType = 6 // four of a kind
		}
	case triple == 1 && double == 1:
		if cardMap["J"] == 2 || cardMap["J"] == 3 {
			h.handType = 7 // five of a kind
		} else if cardMap["J"] == 1 {
			h.handType = 6 // four of a kind
		} else {
			h.handType = 5 // full house
		}
	case triple == 1 && double == 0:
		if cardMap["J"] == 1 || cardMap["J"] == 3 {
			h.handType = 6 // four of a kind
		} else {
			h.handType = 4 // three of a kind
		}
	case double == 2:
		if cardMap["J"] == 1 {
			h.handType = 5 // full house
		} else if cardMap["J"] == 2 {
			h.handType = 6 // four of a kind
		} else {
			h.handType = 3 // two pair
		}
	case double == 1:
		if cardMap["J"] == 1 || cardMap["J"] == 2 {
			h.handType = 4 // three of a kind
		} else {
			h.handType = 2 // one pair
		}
	default:
		if cardMap["J"] == 1 {
			h.handType = 2 // one pair
		} else {
			h.handType = 1 // high card
		}
	}
}

func Rank(hands []hand) []hand {

	//log.Println("before: ", hands)
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType == hands[j].handType {
			for k := 0; k < len(hands[i].cards); k++ {
				if ruleSet[string(hands[i].cards[k])] != ruleSet[string(hands[j].cards[k])] {
					return ruleSet[string(hands[i].cards[k])] < ruleSet[string(hands[j].cards[k])]
				}
			}
			log.Println("WARNING: are both hands equal ? ", hands[i].cards, hands[j].cards)
		}

		return hands[i].handType < hands[j].handType

	})

	log.Println("after: ", hands)

	return hands
}
