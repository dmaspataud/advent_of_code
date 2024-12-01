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

	leftList, rightList := parseLists(text)

	sort.Ints(leftList)
	sort.Ints(rightList)

	for i := 0; i < len(leftList); i++ {
		res += abs(leftList[i] - rightList[i])
	}

	return res
}

func solveSecond(text string) int {
	res := 0
	similarityScore := make(map[int]int)

	leftList, rightList := parseLists(text)

	for i := 0; i < len(rightList); i++ {
		similarityScore[rightList[i]] += 1
	}

	for i := 0; i < len(leftList); i++ {
		res += leftList[i] * similarityScore[leftList[i]]
	}

	return res
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func parseLists(text string) ([]int, []int) {
	leftList, rightList := []int{}, []int{}
	s := bufio.NewScanner(strings.NewReader(text))

	for s.Scan() {
		str := strings.Split(s.Text(), "   ")
		l, err := strconv.Atoi(str[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(str[1])
		if err != nil {
			panic(err)
		}
		leftList = append(leftList, l)
		rightList = append(rightList, r)
	}

	return leftList, rightList
}
