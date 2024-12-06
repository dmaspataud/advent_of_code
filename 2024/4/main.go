package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("first: ", solveFirst(string(f)))
	fmt.Println("second:", solveSecond(string(f)))
}

func solveFirst(text string) int {
	var grid []string
	res := 0

	// grab the strings and make a grid
	s := bufio.NewScanner(strings.NewReader(text))
	for s.Scan() {
		grid = append(grid, s.Text())
	}

	rows := len(grid)
	cols := len(grid[0])

	dx := []int{0, 1, 1, 1, 0, -1, -1, -1}
	dy := []int{1, 1, 0, -1, -1, -1, 0, 1}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'X' {
				for dir := 0; dir < 8; dir++ {
					if checkXMAS(grid, i, j, dx[dir], dy[dir]) {
						res++
					}
				}
			}
		}
	}

	return res
}

func checkXMAS(grid []string, startX, startY, dx, dy int) bool {
	word := "XMAS"
	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < len(word); i++ {
		newX := startX + i*dx
		newY := startY + i*dy

		if newX < 0 || newX >= rows || newY < 0 || newY >= cols {
			return false
		}

		if grid[newX][newY] != word[i] {
			return false
		}
	}
	return true
}

func solveSecond(text string) int {
	var grid []string
	res := 0

	// grab the strings and make a grid
	s := bufio.NewScanner(strings.NewReader(text))
	for s.Scan() {
		grid = append(grid, s.Text())
	}

	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'A' && i > 0 && i < rows-1 && j > 0 && j < cols-1 {
				// Check top-left to bottom-right diagonal
				topLeft := grid[i-1][j-1]
				bottomRight := grid[i+1][j+1]
				// Check top-right to bottom-left diagonal
				topRight := grid[i-1][j+1]
				bottomLeft := grid[i+1][j-1]

				if (topLeft == 'M' && bottomRight == 'S' && topRight == 'M' && bottomLeft == 'S') ||
					(topLeft == 'M' && bottomRight == 'S' && topRight == 'S' && bottomLeft == 'M') ||
					(topLeft == 'S' && bottomRight == 'M' && topRight == 'M' && bottomLeft == 'S') ||
					(topLeft == 'S' && bottomRight == 'M' && topRight == 'S' && bottomLeft == 'M') {
					res++
				}
			}
		}
	}

	return res
}
