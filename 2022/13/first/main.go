package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	index := 1
	var results []int
	var left string
	var right string

	for scanner.Scan() {
		s := scanner.Text()

		if s == "" {
			left = ""
			right = ""
		} else {
			if index%2 == 1 {
				left = s
				index++
			} else if index%2 == 0 {
				right = s
				index++
				// we have a pair, process them
				if isValid(left, right) {
					results = append(results, index)
				}
			}
		}
	}
	fmt.Println(results)
}

func isValid(left string, right string) bool {
	lArray := strings.Split(strings.ReplaceAll(strings.ReplaceAll(left, "[", ""), "]", ""), ",")
	rArray := strings.Split(strings.ReplaceAll(strings.ReplaceAll(right, "[", ""), "]", ""), ",")

}
