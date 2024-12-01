package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type springLine struct {
	springs string
	record  []int
}

var memo = make(map[string]int)

func main() {
	f, err := os.ReadFile("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve(string(f)))
}

func solve(text string) int {

	res := 0

	parsed := (parse(text))

	for i := 0; i < len(parsed); i++ {
		parsed[i] = unfold(parsed[i])
	}

	for i := 0; i < len(parsed); i++ {
		res += count(parsed[i].springs, parsed[i].record)
	}

	return res
}

func parse(text string) []springLine {
	scanner := bufio.NewScanner(strings.NewReader(text))
	var springField []springLine

	for scanner.Scan() {
		line := scanner.Text()
		records := convertNumbers(strings.Split(strings.Split(line, " ")[1], ","))
		springs := strings.Split(line, " ")[0]
		springField = append(springField, springLine{springs: springs, record: records})
	}
	return springField
}

func convertNumbers(numbers []string) []int {
	var res []int

	for _, numStr := range numbers {
		cleanNumStr := strings.TrimSpace(numStr)

		if cleanNumStr == "" {
			continue
		}

		num, err := strconv.Atoi(cleanNumStr)
		if err != nil {
			log.Printf("Error converting %s to int: %v", cleanNumStr, err)
			continue
		}

		res = append(res, int(num))
	}

	return res
}

func unfold(line springLine) springLine {
	var unfoldedLine string
	var unfoldedRecords []int

	for j := 0; j < 5; j++ {
		if j != 4 {
			unfoldedLine += line.springs + "?"
		} else {
			unfoldedLine += line.springs
		}

		unfoldedRecords = append(unfoldedRecords, line.record...)
	}

	return springLine{springs: unfoldedLine, record: unfoldedRecords}
}

func count(springs string, counts []int) int {
	res := 0
	key := springs + strings.Join(strings.Fields(fmt.Sprint(counts)), ",")
	n, hit := memo[key]
	if hit {
		return n
	}

	if len(counts) == 0 {
		if strings.Contains(springs, "#") {
			return 0
		}
		return 1
	}

	minsprings := len(counts) + sum(counts) - 1

	for i := 0; i <= len(springs)-minsprings; i++ {
		if i > 0 && springs[i-1] == '#' {
			break
		}
		if !strings.Contains(springs[i:i+counts[0]], ".") {
			if i+counts[0] == len(springs) && len(counts) == 1 {
				res++
			} else if springs[i+counts[0]] != '#' {
				res += count(springs[i+counts[0]+1:], counts[1:])
			}
		}
	}
	memo[key] = res
	if memo[key] != 0 {
		log.Printf("%v : %v", key, memo[key])
	}
	return res
}

func sum(numbers []int) int {
	res := 0
	for _, number := range numbers {
		res += number
	}

	return res
}
