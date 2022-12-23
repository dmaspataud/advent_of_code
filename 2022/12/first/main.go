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

var rightBoundary int
var upBoundary int

// DEBUG
var (
	visited = Green
	unknown = Red
)

var (
	Red   = Color("\033[1;31m%s\033[0m")
	Green = Color("\033[1;32m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

// DEBUG

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
			if nodeMap["x"+strconv.Itoa(x)+"y"+strconv.Itoa(y)].letter == "S" {
				origin = nodeMap["x"+strconv.Itoa(x)+"y"+strconv.Itoa(y)]
			} else if nodeMap["x"+strconv.Itoa(x)+"y"+strconv.Itoa(y)].letter == "E" {
				solution = nodeMap["x"+strconv.Itoa(x)+"y"+strconv.Itoa(y)]
			}
		}
		y++
	}
	upBoundary = y - 1
}

func drawGrid(hashmap map[string]Node, width int, height int) {
	header := ""
	for i := 0; i <= width; i++ {
		if i < 10 {
			header += "0" + strconv.Itoa(i) + " "
		} else {
			header += strconv.Itoa(i) + " "
		}
	}
	fmt.Println("  ", header)
	for y := 0; y <= height; y++ {
		if y < 10 {
			fmt.Printf("0%v", y)
		} else {
			fmt.Printf("%v", y)
		}
		for x := 0; x <= width; x++ {
			if n, exist := hashmap["x"+strconv.Itoa(x)+"y"+strconv.Itoa(y)]; exist {
				if n.visited {
					fmt.Printf(visited(" %v "), n.letter)
				} else {
					fmt.Printf(unknown(" %v "), n.letter)
				}

			} else {
				fmt.Printf("   ")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	parseMap()
	frontier = append(frontier, origin)

	for len(frontier) > 0 {
		// iterate over each node in the frontier, add viable childs to frontier
		node := frontier[0]
		node.visited = true
		if node.letter == solution.letter {
			var path []Node

			for node.name != origin.name {
				parent := nodeMap[node.parent]
				path = append(path, parent)
				node = parent
			}
			drawGrid(nodeMap, 143, 40)
			fmt.Println("Solution :", len(path))
			os.Exit(0)
		}
		node.Expand()
		node.removeFromFrontier()
	}
	drawGrid(nodeMap, 143, 40)
	fmt.Println("No solution.")
}
