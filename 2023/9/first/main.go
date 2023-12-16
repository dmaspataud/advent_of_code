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

	dataset := parse(text)

	for i := 0; i < len(dataset); i++ {
		dataset[i] = calculateLayers(dataset[i])
		log.Println("test:", dataset[i])

	}
	log.Println("end: ", dataset)

	for i := 0; i < len(dataset); i++ {
		dataset[i] = completeLines(dataset[i])
		res += dataset[i][0][len(dataset[i][0])-1]
	}
	return res
}

func parse(text string) [][][]int {

	scanner := bufio.NewScanner(strings.NewReader(text))
	dataset := [][][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		dataline := convertNumbers(strings.Split(line, " "))
		dataset = append(dataset, [][]int{dataline})
	}

	return dataset
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

func isBottom(nums []int) bool {
	for _, n := range nums {
		if n != 0 {
			return false
		}
	}
	return true
}

func calculateLayers(dataline [][]int) [][]int {

	for {
		nums := dataline[len(dataline)-1]
		underLine := []int{}
		for i := 0; i < len(nums)-1; i++ {
			underLine = append(underLine, nums[i+1]-nums[i])
		}
		dataline = append(dataline, underLine)
		if isBottom(dataline[len(dataline)-1]) {
			return dataline
		}
	}
}

func completeLines(dataline [][]int) [][]int {
	lastline := []int{0}
	for i := len(dataline) - 1; i >= 0; i-- {
		log.Printf("before : i = %v : %v\n", i, dataline[i])
		dataline[i] = append(dataline[i], dataline[i][len(dataline[i])-1]+lastline[len(lastline)-1])
		lastline = dataline[i]
		log.Printf("after : i = %v : %v\n", i, dataline[i])

	}
	return dataline
}
