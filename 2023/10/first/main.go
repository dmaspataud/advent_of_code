package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type location struct {
	x       int
	y       int
	char    string
	visited bool
}

func main() {
	f, err := os.ReadFile("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve(string(f)))
}

func solve(text string) int {
	res := 0
	graph, nextNode := parse(text)

	for {

		n := graph[nextNode]
		n.visited = true
		graph[nextNode] = n

		nextNode = findNext(graph, nextNode)
		res += 1

		if nextNode == "finished" {
			break
		}
	}

	return res / 2
}

func parse(text string) (map[string]location, string) {

	scanner := bufio.NewScanner(strings.NewReader(text))
	start := ""

	y := -1
	graph := make(map[string]location)

	for scanner.Scan() {
		line := scanner.Text()
		y += 1
		for x := 0; x <= len(line)-1; x++ {
			thisPos := fmt.Sprintf("%v:%v", x, y)
			graph[thisPos] = location{char: string(line[x]), x: x, y: y}
			if string(line[x]) == "S" {
				start = fmt.Sprintf("%v:%v", x, y)
			}
		}
	}
	return graph, start
}

func findNext(graph map[string]location, thisPos string) string {
	pos := graph[thisPos]
	down := fmt.Sprintf("%v:%v", pos.x, pos.y+1)
	up := fmt.Sprintf("%v:%v", pos.x, pos.y-1)
	left := fmt.Sprintf("%v:%v", pos.x-1, pos.y)
	right := fmt.Sprintf("%v:%v", pos.x+1, pos.y)

	// the first time we run, we need to replace start with its real char
	if pos.char == "S" {
		pos.char = findSChar(graph, thisPos)
	}

	switch pos.char {
	case "|":
		if !graph[down].visited {
			return down
		} else if !graph[up].visited {
			return up
		}
	case "-":
		if !graph[left].visited {
			return left
		} else if !graph[right].visited {
			return right
		}
	case "L":
		if !graph[up].visited {
			return up
		} else if !graph[right].visited {
			return right
		}
	case "J":
		if !graph[left].visited {
			return left
		} else if !graph[up].visited {
			return up
		}
	case "7":
		if !graph[left].visited {
			return left
		} else if !graph[down].visited {
			return down
		}
	case "F":
		if !graph[right].visited {
			return right
		} else if !graph[down].visited {
			return down
		}
	case ".":
		log.Fatal("you should not be here")
	}

	// if we reach this, it means we there is no more unvisited positions, and we have reached the start point
	return "finished"
}

func findSChar(graph map[string]location, thisPos string) string {
	pos := graph[thisPos]
	down := fmt.Sprintf("%v:%v", pos.x, pos.y+1)
	up := fmt.Sprintf("%v:%v", pos.x, pos.y-1)
	left := fmt.Sprintf("%v:%v", pos.x-1, pos.y)
	right := fmt.Sprintf("%v:%v", pos.x+1, pos.y)
	guess := make(map[string]bool)

	// check up
	if _, ok := graph[up]; ok {
		if graph[up].char == "|" || graph[up].char == "7" || graph[up].char == "F" {

			guess[up] = true
		}
	}
	// check down
	if _, ok := graph[down]; ok {
		if graph[down].char == "|" || graph[down].char == "L" || graph[down].char == "J" {

			guess[down] = true
		}
	}
	// check left
	if _, ok := graph[left]; ok {
		if graph[left].char == "-" || graph[left].char == "L" || graph[left].char == "F" {

			guess[left] = true
		}
	}
	// check right
	if _, ok := graph[right]; ok {
		if graph[right].char == "-" || graph[right].char == "7" || graph[right].char == "J" {

			guess[right] = true
		}
	}

	if guess[up] && guess[down] {
		return "|"
	} else if guess[up] && guess[right] {
		return "L"
	} else if guess[up] && guess[left] {
		return "J"
	} else if guess[down] && guess[right] {
		return "F"
	} else if guess[down] && guess[left] {
		return "7"
	}

	log.Fatal("could not determine S char")
	return ""
}
