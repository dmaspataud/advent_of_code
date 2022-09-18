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

	data := strings.Fields(string(file))

	count := 0

	for i, n := range data {
		depth, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}

		if i+1 < len(data) {
			nextDepth, err := strconv.Atoi(data[i+1])
			if err != nil {
				panic(err)
			}

			if depth < nextDepth {
				count++
			}
		}
	}
	fmt.Println(count)
}
