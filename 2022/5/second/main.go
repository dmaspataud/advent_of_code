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

	total := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

	}
	fmt.Println(total)
}
