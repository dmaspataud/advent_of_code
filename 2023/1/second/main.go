package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve(string(f)))
}

func solve(text string) int {
	res := 0
	numbers := []int{}
	scanner := bufio.NewScanner(strings.NewReader(text))
	textNumbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for scanner.Scan() {
		first, last := "", ""
		line := scanner.Text()
		for i, char := range line {
			if char > '0' && char <= '9' {
				if isNumeral(fmt.Sprint(char)) {
					if first == "" {
						first = string(char)
						last = first
					} else {
						last = string(char)
					}
				}
			} else {
				for word, num := range textNumbers {
					if strings.HasPrefix(line[i:], word) {
						if first == "" {
							first = fmt.Sprint(num)
						}
						last = fmt.Sprint(num)
					}
				}
			}
		}
		this, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, this)
	}

	for _, num := range numbers {
		res += num
	}
	return res
}

func isNumeral(text string) bool {
	_, err := strconv.Atoi(text)

	return err == nil
}
