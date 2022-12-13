package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Tree struct {
	size int
	x    int
	y    int
}

// read the whole file, it will make it easier on parsing
// create function that check if tree is visible from up, down, left, right instead of re-computing everything
// border should straight away be assigned a visible flag (+ total)

func (tree *Tree) computeScore() int {
	tree.visibleFromNorth()
	tree.visibleFromWest()
	tree.visibleFromSouth()
	tree.visibleFromEast()
	return scenicScore[0] * scenicScore[1] * scenicScore[2] * scenicScore[3]
}

func (tree *Tree) visibleFromNorth() bool {
	treeVisible := true
	score := 0

	if tree.y != 0 {
		for direction := tree.y - 1; direction >= 0; direction-- {
			comparedTree := hashmap["x"+strconv.Itoa(tree.x)+"y"+strconv.Itoa(direction)]
			if comparedTree.size >= tree.size {
				treeVisible = false
				score++
				break
			} else {
				score++
			}
		}
	}

	scenicScore = append(scenicScore, score)
	return treeVisible
}

func (tree *Tree) visibleFromSouth() bool {
	treeVisible := true
	score := 0

	if tree.y != yLimit {
		for direction := tree.y + 1; direction <= yLimit; direction++ {
			comparedTree := hashmap["x"+strconv.Itoa(tree.x)+"y"+strconv.Itoa(direction)]
			if comparedTree.size >= tree.size {
				treeVisible = false
				score++
				break
			} else {
				score++
			}
		}
	}

	scenicScore = append(scenicScore, score)
	return treeVisible
}

func (tree *Tree) visibleFromWest() bool {
	treeVisible := true
	score := 0

	if tree.x != 0 {
		for direction := tree.x - 1; direction >= 0; direction-- {
			comparedTree := hashmap["x"+strconv.Itoa(direction)+"y"+strconv.Itoa(tree.y)]
			if comparedTree.size >= tree.size {
				treeVisible = false
				score++
				break
			} else {
				score++
			}
		}
	}

	scenicScore = append(scenicScore, score)
	return treeVisible
}

func (tree *Tree) visibleFromEast() bool {
	treeVisible := true
	score := 0

	if tree.y != xLimit {
		for direction := tree.x + 1; direction <= xLimit; direction++ {
			comparedTree := hashmap["x"+strconv.Itoa(direction)+"y"+strconv.Itoa(tree.y)]
			if comparedTree.size >= tree.size {
				treeVisible = false
				score++
				break
			} else {
				score++
			}
		}
	}

	scenicScore = append(scenicScore, score)
	return treeVisible
}

var hashmap map[string]Tree
var xLimit int
var yLimit int
var scenicScore []int
var maxScenicScore []int

func main() {
	hashmap = make(map[string]Tree)
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// x == column number (left/right)
	// y == line number (up/down)

	scanner := bufio.NewScanner(file)
	curLine := 0
	for scanner.Scan() {
		s := scanner.Text()
		xLimit = len(s) - 1
		for curColumn, rune := range s {
			size, err := strconv.Atoi(string(rune))
			if err != nil {
				panic(err)
			}
			hashmap["x"+strconv.Itoa(curColumn)+"y"+strconv.Itoa(curLine)] = Tree{size: size, x: curColumn, y: curLine} // every tree is visible by default, we will mark them as hidden afterward
		}
		curLine++
	}

	yLimit = curLine - 1

	maxScore := 0
	for _, tree := range hashmap {
		currentScore := tree.computeScore()
		if currentScore > maxScore {
			maxScore = currentScore
			maxScenicScore = scenicScore
		}
		currentScore = 0
		scenicScore = nil
	}
	fmt.Println(maxScore)
}
