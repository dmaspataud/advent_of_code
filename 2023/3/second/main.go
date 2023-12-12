package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type pos struct {
	isNum    bool
	isSymbol bool
	isStar   bool
	value    int
}

func main() {
	f, err := os.ReadFile("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve(string(f)))
}

func solve(text string) int {
	graph, xBoundary, yBoundary := parseMap(text)

	return findParts(graph, xBoundary, yBoundary)
}

func parseMap(text string) (map[string]pos, int, int) {
	graph := make(map[string]pos)
	y := 0
	maxX := 0
	wasNumber := false
	continueNumberList := []string{}

	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		line := scanner.Text()
		maxX = len(line)
		parsedNumber := ""

		for x, char := range line {
			thisPos := fmt.Sprintf("%vx%v", x, y)

			if string(char) == "." {
				graph[thisPos] = pos{isNum: false, isSymbol: false}
				wasNumber = false
			} else {
				if unicode.IsNumber(char) {
					graph[thisPos] = pos{isNum: true, isSymbol: false}
					wasNumber = true
					parsedNumber = fmt.Sprint(parsedNumber + string(char))
					continueNumberList = append(continueNumberList, thisPos)
				} else if string(char) == "*" {
					graph[thisPos] = pos{isNum: false, isSymbol: true, isStar: true}
					wasNumber = false
				} else {
					graph[thisPos] = pos{isNum: false, isSymbol: true}
					wasNumber = false
				}
			}
			if x == maxX-1 && len(continueNumberList) > 0 || (!wasNumber && len(continueNumberList) > 0) {
				val, err := strconv.Atoi(parsedNumber)
				if err != nil {
					log.Fatal(err)
				}
				for _, p := range continueNumberList {
					graph[p] = pos{isNum: graph[p].isNum, isSymbol: graph[p].isSymbol, value: val}
				}

				// clear the buffers
				wasNumber = false
				continueNumberList = []string{}
				parsedNumber = ""
			}
		}
		y += 1

	}
	return graph, maxX, y
}

func findParts(graph map[string]pos, xBoundary int, yBoundary int) int { // iterate over every pos, if it's a *, look for numbers around it
	total := 0
	res := make(map[int]bool)

	for y := 0; y < yBoundary; y++ {
		for x := 0; x < xBoundary; x++ {
			thisPos := fmt.Sprintf("%vx%v", x, y)
			if graph[thisPos].isStar {
				gears := []int{}

				// check left
				if x > 0 {
					//log.Print("left: ", fmt.Sprintf("%vx%v", x-1, y))
					if graph[fmt.Sprintf("%vx%v", x-1, y)].isNum {
						gears = append(gears, graph[fmt.Sprintf("%vx%v", x-1, y)].value)
						res[graph[fmt.Sprintf("%vx%v", x-1, y)].value] = true
					}
				}
				// check right
				if x < xBoundary {
					if graph[fmt.Sprintf("%vx%v", x+1, y)].isNum {
						res[graph[fmt.Sprintf("%vx%v", x+1, y)].value] = true
					}
				}
				// check top
				if y > 0 {
					if graph[fmt.Sprintf("%vx%v", x, y-1)].isNum {
						res[graph[fmt.Sprintf("%vx%v", x, y-1)].value] = true
					}
				}
				// check bottom
				if y < yBoundary {
					if graph[fmt.Sprintf("%vx%v", x, y+1)].isNum {
						res[graph[fmt.Sprintf("%vx%v", x, y+1)].value] = true
					}
				}
				// check top left
				if y > 0 && x > 0 {
					if graph[fmt.Sprintf("%vx%v", x-1, y-1)].isNum {
						res[graph[fmt.Sprintf("%vx%v", x-1, y-1)].value] = true
					}
				}
				// check top right
				if y > 0 && x < xBoundary {
					if graph[fmt.Sprintf("%vx%v", x+1, y-1)].isNum {
						res[graph[fmt.Sprintf("%vx%v", x+1, y-1)].value] = true
					}
				}

				// check bottom left
				if y < yBoundary && x > 0 {
					if graph[fmt.Sprintf("%vx%v", x-1, y+1)].isNum {
						res[graph[fmt.Sprintf("%vx%v", x-1, y+1)].value] = true
					}
				}

				// check bottom right
				if y < yBoundary && x < xBoundary {
					if graph[fmt.Sprintf("%vx%v", x+1, y+1)].isNum {
						res[graph[fmt.Sprintf("%vx%v", x+1, y+1)].value] = true
					}
				}
			} else {
				if len(res) == 2 {
					theseGears := []int{}
					for value := range res {
						theseGears = append(theseGears, value)
					}
					total += (theseGears[0] * theseGears[1])
					log.Printf("Added %v, total: %v", theseGears, total)
				}
				for k := range res {
					delete(res, k)
				}

			}
		}
	}
	return total
}
