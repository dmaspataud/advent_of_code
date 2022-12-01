package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Getting input data
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// variables we'll need
	var elvesSums []int
	sum := 0
	result := 0

	// getting an array of sums
	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		if line != "" {
			num, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			sum += num
		} else {
			elvesSums = append(elvesSums, sum)
			sum = 0
		}
	}

	// sorting
	sort.Ints(elvesSums)

	// sum of the last 3
	for _, cal := range elvesSums[len(elvesSums)-3:] {
		result += cal
	}

	fmt.Println(result)
}
