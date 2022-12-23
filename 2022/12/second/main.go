package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Node struct {
	letter  string
	name    string
	x, y    int
	parent  string
	visited bool
}

var origin Node
var solution Node
var frontier []Node
var nodeMap map[string]Node
var bestPath int
var rightBoundary int
var upBoundary int

func (node *Node) Expand() {
	parentNode := nodeMap[node.parent]
	if node.y != upBoundary {
		thisNode := nodeMap["x"+strconv.Itoa(node.x)+"y"+strconv.Itoa(node.y+1)]
		if thisNode != parentNode && !thisNode.visited && isClimbable(thisNode, *node) && !thisNode.isInFrontier() {
			thisNode.addToFrontier(*node)
		} else {
		}
	}
	if node.x != rightBoundary {
		thisNode := nodeMap["x"+strconv.Itoa(node.x+1)+"y"+strconv.Itoa(node.y)]
		if thisNode != parentNode && !thisNode.visited && isClimbable(thisNode, *node) && !thisNode.isInFrontier() {
			thisNode.addToFrontier(*node)
		} else {
		}
	}
	if node.y != 0 {
		thisNode := nodeMap["x"+strconv.Itoa(node.x)+"y"+strconv.Itoa(node.y-1)]
		if thisNode != parentNode && !thisNode.visited && isClimbable(thisNode, *node) && !thisNode.isInFrontier() {
			thisNode.addToFrontier(*node)
		} else {
		}
	}
	if node.x != 0 {
		thisNode := nodeMap["x"+strconv.Itoa(node.x-1)+"y"+strconv.Itoa(node.y)]
		if thisNode != parentNode && !thisNode.visited && isClimbable(thisNode, *node) && !thisNode.isInFrontier() {
			thisNode.addToFrontier(*node)
		} else {
		}
	}
}

func (node *Node) removeFromFrontier() {
	nodeMap["x"+strconv.Itoa(node.x)+"y"+strconv.Itoa(node.y)] = *node
	frontier = frontier[1:]
}

func (node *Node) addToFrontier(parent Node) {
	node.parent = parent.name
	frontier = append(frontier, *node)
}

func (node *Node) isInFrontier() bool {
	for i := 0; i < len(frontier); i++ {
		if frontier[i].name == node.name {
			return true
		}
	}
	return false
}

func isClimbable(n1 Node, n2 Node) bool {
	var l1, l2 int

	if n1.letter == "S" {
		l1 = int("a"[0])
	} else if n1.letter == "E" {
		l1 = int("z"[0])
	} else {
		l1 = int(n1.letter[0])
	}

	if n2.letter == "S" {
		l2 = int("a"[0])
	} else if n2.letter == "E" {
		l2 = int("z"[0])
	} else {
		l2 = int(n2.letter[0])
	}

	d := l1 - l2
	if d < 2 {
		return true
	} else {
		return false
	}
}

// Parse map, locate origin and solution
func parseMap() {
	nodeMap = make(map[string]Node)
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	y := 0

	for scanner.Scan() {
		s := scanner.Text()
		for x, l := range s {
			rightBoundary = len(s) - 1
			nodeMap["x"+strconv.Itoa(x)+"y"+strconv.Itoa(y)] = Node{x: x, y: y, letter: string(l), name: "x" + strconv.Itoa(x) + "y" + strconv.Itoa(y), visited: false}
			if nodeMap["x"+strconv.Itoa(x)+"y"+strconv.Itoa(y)].letter == "E" {
				solution = nodeMap["x"+strconv.Itoa(x)+"y"+strconv.Itoa(y)]
			}
		}
		y++
	}
	upBoundary = y - 1
}

func main() {
	parseMap()
	bestPath = 999999
	secondMap := nodeMap // make a copy of the map and iterate over it
	// loop over map, each time you find a a letter, start from there
	for _, node := range secondMap {
		if node.letter == "a" {
			parseMap() // for some reason, reseting the flag didn't do the trick, but re-parsing worked
			origin = node

			// reset frontier, then add origin
			frontier = make([]Node, 0)
			frontier = append(frontier, origin)
			for len(frontier) > 0 {
				// iterate over each node in the frontier, add viable childs to frontier
				node := frontier[0]
				node.visited = true

				// if we found the solution, compute the way back
				if node.letter == solution.letter {
					var path []Node

					for node.name != origin.name {
						parent := nodeMap[node.parent]
						path = append(path, parent)
						node = parent
					}
					//drawGrid(nodeMap, 143, 40)
					if bestPath > len(path) {
						bestPath = len(path)
					}
				} else {
					node.Expand()
				}
				node.removeFromFrontier()
			}
		}
	}
	fmt.Println(bestPath)
}
