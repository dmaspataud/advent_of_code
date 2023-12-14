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
	directions, network, entrypoints := parse(text)

	solutions := []int{}

	for j := 0; j < len(entrypoints); j++ {
		log.Printf("Evaluating %v", entrypoints[j])
		step := entrypoints[j]
		stepsToEnd := 0

		for i := 0; !foundSolution(step); i++ {

			if i == len(directions) {
				i = 0
			}

			stepsToEnd += 1

			if directions[i] == "L" {
				step = network[step].left
			} else if directions[i] == "R" {
				step = network[step].right
			} else {
				log.Fatal("Error: Next direction unknown.")
			}
		}

		solutions = append(solutions, stepsToEnd)

	}

	log.Println(solutions)
	return LCM(solutions[0], solutions[1], solutions...)
}

func parse(text string) ([]string, map[string]node, []string) {

	entrypoints := []string{}
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

			ok, _ = regexp.MatchString(".*A", n)
			if ok {
				entrypoints = append(entrypoints, n)
			}
		}
	}

	return directions, network, entrypoints
}

func foundSolution(step string) bool {
	ok, _ := regexp.MatchString(".*Z", step)
	if !ok {
		return false
	}
	return true
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
