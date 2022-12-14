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
		if diffY >= 1 {
			tail.move("U")
		} else if diffY <= -1 {
			tail.move("D")
		}
		tail.move("R")
	} else if diffX == -2 {
		if diffY >= 1 {
			tail.move("U")
		} else if diffY <= -1 {
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

// func drawGrid(hashmap map[string]Cell, width int, height int) {
// 	fmt.Printf("   ")
// 	for i := -50; i <= width/2; i++ {
// 		if i%5 == 0 {
// 			fmt.Printf(" | ")
// 		} else {
// 			fmt.Printf(" . ")
// 		}

// 	}
// 	fmt.Printf(" \n")
// 	for y := height / 2; y >= -50; y-- {
// 		if y%5 == 0 {
// 			if y < 10 && y >= 0 {
// 				fmt.Printf("0%v ", y)
// 			} else if y <= -10 {
// 				fmt.Printf("%v", y)
// 			} else {
// 				fmt.Printf("%v ", y)
// 			}
// 		} else {
// 			fmt.Printf("-  ")
// 		}
// 		for x := -50; x <= width/2; x++ {
// 			if _, exist := hashmap[strconv.Itoa(x)+":"+strconv.Itoa(y)]; exist {
// 				fmt.Printf(" # ")
// 			} else {
// 				fmt.Printf(" . ")
// 			}
// 		}
// 		fmt.Printf("\n")
// 	}
// }

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hashmap := make(map[string]Cell)
	head := Part{0, 0}
	tails := make([]Part, 10)
	count := 0

	// creating rope with 10 tails
	for i := 0; i < 10; i++ {
		tails = append(tails, Part{x: 0, y: 0})
	}

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
			tails[0].follow(head)
			for i := 1; i < 9; i++ {
				tails[i].follow(tails[i-1])
			}
			hashmap[strconv.Itoa(tails[8].x)+":"+strconv.Itoa(tails[8].y)] = Cell{x: tails[8].x, y: tails[8].y}
		}
	}
	fmt.Println(len(hashmap))
}
