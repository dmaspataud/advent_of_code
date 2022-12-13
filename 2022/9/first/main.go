package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Part struct {
	x int
	y int
}

type Cell struct {
	x int
	y int
}

func (part *Part) move(direction string) {
	switch direction {
	case "U":
		part.y++
	case "D":
		part.y--
	case "L":
		part.x--
	case "R":
		part.x++
	}
}

func (tail *Part) follow(head Part) {
	diffX := head.x - tail.x
	diffY := head.y - tail.y

	if diffX == 2 {
		if diffY == 1 {
			tail.move("U")
		} else if diffY == -1 {
			tail.move("D")
		}
		tail.move("R")
	} else if diffX == -2 {
		if diffY == 1 {
			tail.move("U")
		} else if diffY == -1 {
			tail.move("D")
		}
		tail.move("L")
	} else if diffY == 2 {
		if diffX >= 1 {
			tail.move("R")
		} else if diffX <= -1 {
			tail.move("L")
		}
		tail.move("U")
	} else if diffY == -2 {
		if diffX >= 1 {
			tail.move("R")
		} else if diffX <= -1 {
			tail.move("L")
		}
		tail.move("D")
	}

}

func drawGrid(hashmap map[string]Cell, width int, height int) {
	header := ""
	for i := 0; i <= width; i++ {
		if i < 10 {
			header += "0" + strconv.Itoa(i) + " "
		} else {
			header += strconv.Itoa(i) + " "
		}
	}
	fmt.Println("  ", header)
	for y := width; y >= 0; y-- {
		if y < 10 {
			fmt.Printf("0%v", y)
		} else {
			fmt.Printf("%v", y)
		}
		for x := 0; x <= height; x++ {
			if _, exist := hashmap[strconv.Itoa(x)+":"+strconv.Itoa(y)]; exist {
				fmt.Printf(" # ")
			} else {
				fmt.Printf(" . ")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hashmap := make(map[string]Cell)
	head := Part{0, 0}
	tail := Part{0, 0}
	count := 0

	for scanner.Scan() {
		s := scanner.Text()
		count++
		var direction string
		var steps int
		_, err := fmt.Sscanf(s, "%s %d", &direction, &steps)
		if err != nil {
			panic(err)
		}
		for i := 0; i < steps; i++ {
			head.move(direction)
			tail.follow(head)
			hashmap[strconv.Itoa(tail.x)+":"+strconv.Itoa(tail.y)] = Cell{x: tail.x, y: tail.y}
		}
	}
	fmt.Println(len(hashmap))
}
