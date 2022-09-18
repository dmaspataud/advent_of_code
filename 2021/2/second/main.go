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

	depth := 0
	horizontal := 0
	aim := 0

	data := strings.Split(string(file), "\n")
	for _, n := range data {
		direction := strings.Split(n, " ")[0]
		val, err := strconv.Atoi(strings.Split(n, " ")[1])
		if err != nil {
			panic(err)
		}

		if direction == "up" {
			aim = aim - val
		} else if direction == "down" {
			aim = aim + val
		} else if direction == "forward" {
			horizontal = horizontal + val
			depth = depth + (aim * val)
		} else {
			panic(direction)
		}
	}
	fmt.Println(depth * horizontal)
}
