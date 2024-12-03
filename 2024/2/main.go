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
	f, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("first: ", solveFirst(string(f)))
	fmt.Println("second: ", solveSecond(string(f)))
}

func solveFirst(text string) int {
	res := 0

	reportList := parseList(text)

	for _, report := range reportList {
		if isSafeReport(report) {
			res++
		}
	}
	return res
}

func solveSecond(text string) int {
	res := 0
	reportList := parseList(text)

	for _, report := range reportList {
		if isSafeReport(report) {
			res++
		} else {
			for i := 0; i < len(report); i++ {
				dampenedReport := removeIdxFromSlice(report, i)
				if isSafeReport(dampenedReport) {
					res++
					break
				}
			}
		}
	}
	return res
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func removeIdxFromSlice(slice []int, i int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return append(newSlice[:i], newSlice[i+1:]...)
}

func parseList(text string) [][]int {
	reportList := [][]int{}
	s := bufio.NewScanner(strings.NewReader(text))

	for s.Scan() {
		str := strings.Fields(s.Text())
		line := []int{}
		for _, char := range str {
			num, err := strconv.Atoi(char)
			if err != nil {
				panic(err)
			}
			line = append(line, num)
		}
		reportList = append(reportList, line)
	}

	return reportList
}

func isSafeReport(report []int) bool {
	var isAscendant bool

	if report[0] < report[1] {
		isAscendant = true
	}

	for i := 0; i < len(report)-1; i++ { // len(report) - 1 to stop after the penultimate
		if isAscendant && report[i] > report[i+1] ||
			!isAscendant && report[i] < report[i+1] ||
			abs(report[i]-report[i+1]) > 3 ||
			abs(report[i]-report[i+1]) == 0 {
			return false
		}

		if i == len(report)-2 {
			return true
		}
	}
	return false
}
