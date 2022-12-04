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
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		a, err := strconv.Atoi(strings.Split(string(s[0]), "-")[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(strings.Split(string(s[0]), "-")[1])
		if err != nil {
			panic(err)
		}
		c, err := strconv.Atoi(strings.Split(string(s[1]), "-")[0])
		if err != nil {
			panic(err)
		}
		d, err := strconv.Atoi(strings.Split(string(s[1]), "-")[1])
		if err != nil {
			panic(err)
		}

		if (a >= c && a <= d && b <= d && b >= c) || (c >= a && c <= b && d >= a && d <= b) {
			total += 1
		}

	}
	fmt.Println(total)
}
