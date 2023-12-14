package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type node struct {
	left  string
	right string
}

func main() {
	f, err := os.ReadFile("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve(string(f)))
}

func solve(text string) int {
	steps := 0
	directions, network := parse(text)

	thisStep := "AAA" // we always start at AAA

	for i := 0; thisStep != "ZZZ"; i++ {
		if i == len(directions) {
			i = 0
		}
		steps += 1

		if directions[i] == "L" {
			thisStep = network[thisStep].left
		} else if directions[i] == "R" {
			thisStep = network[thisStep].right
		} else {
			log.Fatal("Error: Next direction unknown.")
		}
	}

	return steps
}

func parse(text string) ([]string, map[string]node) {

	directions := []string{}
	network := make(map[string]node)
	scanner := bufio.NewScanner(strings.NewReader(text))

	for scanner.Scan() {
		line := scanner.Text()

		ok, err := regexp.MatchString(".*=.*", line) // node description

		if err != nil {
			log.Fatal(err)
		}

		if !ok {
			ok, err = regexp.MatchString("L|R+", line) // directions
			if err != nil {
				log.Fatal(err)
			}

			if ok {
				directions = strings.Split(line, "")
			}
		} else {
			n, ln, rn := "", "", ""

			fmt.Sscanf(line, "%v = (%s %s)", &n, &ln, &rn)
			ln = strings.Replace(ln, ",", "", -1) // remove the comma because sscanf is misbehaving
			rn = strings.Replace(rn, ")", "", -1) // remove the parenthesis because sscanf is misbehaving
			network[n] = node{left: ln, right: rn}
		}

	}

	return directions, network
}
