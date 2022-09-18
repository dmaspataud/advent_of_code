package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	data := stringsToInts(strings.Fields(string(file)))

	count := 0

	for i := 0; i+3 < len(data); i++ {
		if data[i]+data[i+1]+data[i+2] < data[i+1]+data[i+2]+data[i+3] {
			count++
		}
	}
	fmt.Println(count)
}

func stringsToInts(stringsSlice []string) []int {
	intSlice := []int{}
	for i := range stringsSlice {
		s := stringsSlice[i]
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, n)
	}
	return intSlice
}
