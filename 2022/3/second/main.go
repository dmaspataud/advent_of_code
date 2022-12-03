package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	var group []string
	counter := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		group = append(group, scanner.Text())
		counter += 1

		if counter == 3 {
			hashmap := make(map[string]int)
			for _, l := range group[0] { // fill hashmap with items from first bag
				hashmap[string(l)] = 1 // we assign value 1 as its found in bag 1
			}

			for _, l := range group[1] {
				if _, present := hashmap[string(l)]; present {
					hashmap[string(l)] = 2 // we assign value 2 if also found in bag 2
				}
			}
			for _, l := range group[2] {
				if _, present := hashmap[string(l)]; present { // if we find the char that was also found in bag 2
					if hashmap[string(l)] == 2 {
						hashmap[string(l)] = 3 // we tag the token as 3 not to count it several times
						if unicode.IsUpper(l) {
							total += int(l) - (64 - 26)
						} else {
							total += int(l) - (64 + 32)
						}
					}
				}
			}
			group = nil // clean group and counter
			counter = 0
		}
	}
	fmt.Println(total)
}
