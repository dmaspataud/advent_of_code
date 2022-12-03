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

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		hashmap := make(map[string]int)
		for i, l := range s[:len(s)/2] {
			hashmap[string(l)] = i
		}

		for _, l := range s[len(s)/2:] {
			if _, present := hashmap[string(l)]; present {
				if unicode.IsUpper(l) {
					total += int(l) - (64 - 26)
				} else {
					total += int(l) - (64 + 32)
				}
				break
			}
		}
	}
	fmt.Println(total)
}
