package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type springLine struct {
	springs string
	record  []int
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

	parsed := (parse(text))

	for _, v := range parsed {
		variants := generateVariants(strings.Count(v.springs, "?"), "")

		for _, variant := range variants {
			if isValidArrangement(generateArrangement(variant, v.springs), v.record) {
				res += 1
			}
		}
	}

	return res
}

func parse(text string) []springLine {
	scanner := bufio.NewScanner(strings.NewReader(text))
	var springField []springLine

	for scanner.Scan() {
		line := scanner.Text()
		records := convertNumbers(strings.Split(strings.Split(line, " ")[1], ","))
		springs := strings.Split(line, " ")[0]
		springField = append(springField, springLine{springs: springs, record: records})
	}
	return springField
}

func convertNumbers(numbers []string) []int {
	var res []int

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

		res = append(res, int(num))
	}

	return res
}

func isValidArrangement(springs string, records []int) bool {
	currentRecord := 0
	currentSetLength := 0
	springs = fmt.Sprintf("%s.", springs) // we add "." to create artificial boundary at the end of the string

	for i := 0; i < len(springs); i++ {
		switch string(springs[i]) {
		case "#":
			currentSetLength += 1
		case ".":
			if currentSetLength != 0 {
				if currentRecord >= len(records) || records[currentRecord] != currentSetLength {
					return false
				}
				currentSetLength = 0
				currentRecord += 1
			}
		}
	}

	return currentRecord == len(records)
}

func generateVariants(length int, current string) []string {
	if length == 0 {
		return []string{current}
	}

	var variants []string
	variants = append(variants, generateVariants(length-1, current+".")...)
	variants = append(variants, generateVariants(length-1, current+"#")...)

	return variants
}

func generateArrangement(variant string, springs string) string {
	cur := 0
	arrangement := ""

	for i := 0; i < len(springs); i++ {
		if string(springs[i]) == "?" {
			arrangement += string(variant[cur])
			cur++
		} else {
			arrangement += string(springs[i])
		}
	}

	return arrangement
}
