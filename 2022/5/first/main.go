package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cargoParsed := false
	cargoState := make(map[string][]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" && line != " 1   2   3   4   5   6   7   8   9 " && cargoParsed == false {
			parseCargo(line, cargoState)
		} else if line == "" && cargoParsed == false {
			cargoParsed = true
		} else if cargoParsed == true {
			move := strings.Split(line, " ")
			moveCargo(move[1], move[3], move[5], cargoState)

		}
	}
	fmt.Println(computeResult(cargoState))
}

func parseCargo(line string, cargoState map[string][]string) {
	for i := 1; i <= 9; i++ {
		n := (i * 4)
		if i == 1 && line[1] <= 90 && line[1] >= 65 { // if it's the first column, it's composed of 3 chars instead of 4. Check if it's a letter, if it, append
			if len(cargoState[strconv.Itoa(i)]) == 0 {
				cargoState[strconv.Itoa(i)] = append(cargoState[strconv.Itoa(i)], line[1:2])
			} else {
				cargoState[strconv.Itoa(i)] = append(cargoState[strconv.Itoa(i)][:1], cargoState[strconv.Itoa(i)][0:]...)
				cargoState[strconv.Itoa(i)][0] = line[1:2]
			}
		} else if line[n-3] <= 90 && line[n-3] >= 65 {
			if len(cargoState[strconv.Itoa(i)]) == 0 {
				cargoState[strconv.Itoa(i)] = append(cargoState[strconv.Itoa(i)], line[n-3:n-2])
			} else {
				cargoState[strconv.Itoa(i)] = append(cargoState[strconv.Itoa(i)][:1], cargoState[strconv.Itoa(i)][0:]...)
				cargoState[strconv.Itoa(i)][0] = line[n-3 : n-2]
			}
		}
		fmt.Println(cargoState)
	}
}

func moveCargo(n string, src string, dst string, cargoState map[string][]string) {
	qty, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}

	for i := 0; i < qty; i++ {
		cargoState[dst] = append(cargoState[dst], cargoState[src][len(cargoState[src])-1]) // add item from source stack to destination stack
		cargoState[src] = cargoState[src][:len(cargoState[src])-1]                         // remove item from source stack
	}
}

func computeResult(cargoState map[string][]string) string {
	result := ""
	for i := 1; i <= 9; i++ {
		result += cargoState[strconv.Itoa(i)][len(cargoState[strconv.Itoa(i)])-1]
	}
	return result
}
