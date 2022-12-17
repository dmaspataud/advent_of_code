package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Operation struct {
	cycles int
	value  int
}

func processQueue(queue *[]Operation, register int) int {
	var op Operation
	if len(*queue) > 0 {
		op = (*queue)[0]
	} else {
		return register
	}

	op.cycles--
	if op.cycles == 0 {
		register += op.value
		(*queue) = append((*queue)[:0], (*queue)[1:]...) // if operation is finished, remove it from queue
	} else {
		(*queue)[0] = op
	}

	return register
}

func main() {

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cycle := 1 // start at cycle 1

	queue := make([]Operation, 0) // map[cycle]value
	register := 1

	crt := ""
	result := 0
	for scanner.Scan() || len(queue) > 0 {

		s := scanner.Text()

		// Start Cycle
		var instruction string
		var value int
		if len(s) > 0 {
			_, err := fmt.Sscanf(s, "%v %v", &instruction, &value)
			if err != nil {
				_, err = fmt.Sscanf(s, "%v", &instruction)
				if err != nil {
					panic(err)
				}
			}
		}

		if (cycle%40)-2 == register || cycle%40-1 == register || (cycle%40) == register {
			crt += "#"
		} else {
			crt += "."
		}
		if cycle%40 == 0 {
			crt += "\n"
		}

		if instruction == "addx" {
			op := Operation{cycles: 2, value: value}
			queue = append(queue, op)
		} else if instruction == "noop" {
			op := Operation{cycles: 1, value: 0} // noop
			queue = append(queue, op)
		}

		// compute signal
		if cycle == 20 || (cycle-20)%40 == 0 {
			result += register * cycle
		}

		register = processQueue(&queue, register)

		cycle++
	}
	fmt.Println(crt)
}
