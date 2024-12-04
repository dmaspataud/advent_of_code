package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("first: ", solveFirst(string(f)))
	fmt.Println("second:", solveSecond(string(f)))
}

func solveFirst(text string) int {
	res := 0

	s := bufio.NewScanner(strings.NewReader(text))

	for s.Scan() {
		var total int
		nums := parseFirstLine(s.Text())

		for _, num := range nums {
			total += num
		}
		res += total
	}
	return res
}

func solveSecond(text string) int {
	var nums []int
	res := 0
	do := true
	s := bufio.NewScanner(strings.NewReader(text))

	for s.Scan() {
		var total int
		nums, do = parseSecondLine(s.Text(), do)

		for _, num := range nums {
			total += num
		}
		res += total
	}
	return res
}

func parseFirstLine(text string) []int {
	var numbers []int
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := pattern.FindAllStringSubmatch(text, -1)

	for _, match := range matches {
		first, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}

		second, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, first*second)
	}

	return numbers
}

func parseSecondLine(text string, do bool) ([]int, bool) {
	var numbers []int
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)|don't\(\)|do\(\)`)
	matches := pattern.FindAllStringSubmatch(text, -1)

	for _, match := range matches {
		if match[0] == "do()" {
			do = true
		} else if match[0] == "don't()" {
			do = false
		} else if do && strings.HasPrefix(match[0], "mul") {
			first, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatalf("encountered error while converting %v to int: %v", match[1], err)
			}

			second, err := strconv.Atoi(match[2])
			if err != nil {
				log.Fatalf("encountered error while converting %v to int: %v", match[2], err)
			}

			numbers = append(numbers, first*second)
		}
	}

	return numbers, do
}
