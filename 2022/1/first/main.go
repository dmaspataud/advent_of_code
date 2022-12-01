package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	maxSum := 0
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			num, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			sum += num
		} else {
			if sum > maxSum {
				maxSum = sum
			}
			sum = 0
		}
	}
	fmt.Println(maxSum)
}
