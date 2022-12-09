package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		hashmap := make(map[string]int)

		for i := 0; i < len(s); i++ {
			if _, exist := hashmap[s[i:i+1]]; exist { // if the letter is already in the map, we remove every item until this one reset the map and add that letter
				offset := hashmap[s[i:i+1]]
				// remove items in order until we get to the duplicate char
				for k, v := range hashmap {
					if v <= offset {
						delete(hashmap, k)
					} else {
					}
				}
				hashmap[s[i:i+1]] = i
			} else {
				hashmap[s[i:i+1]] = i
				if len(hashmap) == 4 {
					fmt.Println(i + 1)
					os.Exit(0)
				}
			}

		}
	}
}
