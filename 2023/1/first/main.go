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

	for scanner.Scan() {
		line := scanner.Text()
		first, last := "", ""
		for i := 0; i < len(line); i++ {
			if isNumeral(line[i]) {
				if first == "" {
					first = string(line[i])
					last = first
				} else {
					last = string(line[i])
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

func isNumeral(bytes byte) bool {
	_, err := strconv.Atoi(string(bytes))

	return err == nil
}
