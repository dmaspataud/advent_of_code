package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Race struct {
	time int
	best int
}

func main() {
	f, err := os.ReadFile("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve(string(f)))
}

func solve(text string) int {
	res := 0
	races := parseRaces(text)

	for _, race := range races {
		if res == 0 {
			res = race.Calculate()
		} else {
			res *= race.Calculate()
		}
	}

	log.Println(races)
	return res
}

func parseRaces(text string) []Race {
	races := []Race{}
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`.*:`)

		if re.FindString(line) == "Time:" {
			num := convertNumbers(strings.Split(strings.Split(line, ":")[1], " "))
			for _, time := range num {
				races = append(races, Race{time: time})
			}
		} else if re.FindString(line) == "Distance:" {
			num := convertNumbers(strings.Split(strings.Split(line, ":")[1], " "))
			for i, distance := range num {
				races[i] = Race{time: races[i].time, best: distance}
			}
		}
	}
	return races
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

		res = append(res, num)
	}

	return res
}

func (race Race) Calculate() int {
	res := 0
	for pushButton := 0; pushButton < race.time; pushButton++ {
		thisDistance := pushButton * (race.time - pushButton)
		if thisDistance > race.best {
			res++
		}
	}
	return res
}
