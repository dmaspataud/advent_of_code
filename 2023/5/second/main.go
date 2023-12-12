package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type mapping struct {
	dst    int64
	src    int64
	offset int64
}

func main() {
	f, err := os.ReadFile("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve(string(f)))
}

func solve(text string) int64 {
	res := int64(math.MaxInt64)
	var seed int64
	seeds, almanach := parseAlmanach(text)
	for i := 0; i < len(seeds); i += 2 {
		seedSrc, seedOffset := seeds[i], seeds[i+1]
		for j := seedSrc; j < (seedSrc + seedOffset); j++ {
			seed = j
			for section := 1; section <= len(almanach); section++ {
				for subsections := 0; subsections < len(almanach[section]); subsections++ {
					sub := almanach[section][subsections]
					if seed >= sub.src && seed < (sub.src+sub.offset) {
						seed = (seed - sub.src) + sub.dst
						break
					}
				}
			}
			res = min(res, seed)
		}
	}

	log.Println(res)
	return res
}

func parseAlmanach(text string) ([]int64, map[int][]mapping) {
	var seeds []int64
	scanner := bufio.NewScanner(strings.NewReader(text))
	currentSection := 0
	almanach := make(map[int][]mapping)
	section := []mapping{}

	for scanner.Scan() {
		line := scanner.Text()

		// Parse seeds
		ok, err := regexp.Match("seeds:", []byte(line))
		if err != nil {
			log.Fatal(err)
		}
		if ok {
			seeds = convertNumbers(strings.Split(strings.Split(line, ":")[1], " "))
			log.Println("seeds: ", seeds)
			continue
		}

		// Keep track of section
		ok, err = regexp.Match("map:", []byte(line))
		if err != nil {
			log.Fatal(err)
		}
		if ok {
			almanach[currentSection] = section
			currentSection += 1
			section = []mapping{}
		}

		// look for numbers
		ok, err = regexp.Match("\\d+ \\d+ \\d+", []byte(line))
		if err != nil {
			log.Fatal(err)
		}
		if ok {
			n := convertNumbers(strings.Split(line, " "))
			mapping := mapping{dst: n[0], src: n[1], offset: n[2]}
			section = append(section, mapping)
		}
	}
	return seeds, almanach
}

func convertNumbers(numbers []string) []int64 {
	var res []int64

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

		res = append(res, int64(num))
	}

	return res
}
