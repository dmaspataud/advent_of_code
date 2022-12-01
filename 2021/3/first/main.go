package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	

	file, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	data := strings.Split(string(file), "\n")
	inputCount := len(data)

	for i := 0; i < 12; i++ { // input is 12 chars wide
		total := 0
		for _, sbit := range data {
			bit, err := strconv.Atoi(sbit)
			if err != nil {
				panic(err)
			}
			total += bit
		}
		if total < inputCount/2 {


		}
	}

}
