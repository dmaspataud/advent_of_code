package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.ReadFile("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve(string(f)))
}

func solve(text string) int {
	res := 0

	locations := getGalaxiesLocations(parse(text))

	for first := 0; first < len(locations); first++ {
		for second := 0; second < len(locations); second++ {
			res += calculateDistance(locations[first], locations[second])
		}
	}

	return res / 2
}

func parse(text string) [][]string {
	scanner := bufio.NewScanner(strings.NewReader(text))
	var galaxyMap [][]string

	for scanner.Scan() {
		line := scanner.Text()
		galaxyMap = append(galaxyMap, strings.Split(line, ""))
	}

	galaxyMap = applyUniverseExpansion(galaxyMap)

	return galaxyMap
}

func applyUniverseExpansion(galaxyMap [][]string) [][]string {
	// horizontal expansion
	for y := 0; y < len(galaxyMap); y++ {

		isGalaxy := false
		for x := 0; x < len(galaxyMap[y]); x++ {
			if galaxyMap[y][x] == "#" {
				isGalaxy = true
				break
			}
		}
		if !isGalaxy {
			galaxyMap = append(galaxyMap[:y], append([][]string{galaxyMap[y]}, galaxyMap[y:]...)...)
			y++
		}
	}

	// vertical expansion
	expandList := []int{}
	for x := 0; x < len(galaxyMap[0]); x++ {
		isGalaxy := false
		for y := 0; y < len(galaxyMap); y++ {
			if galaxyMap[y][x] == "#" {
				isGalaxy = true
				break
			}
		}
		if !isGalaxy {
			expandList = append(expandList, x)
		}
	}

	shift := 0
	for _, x := range expandList {
		// we need to shift by one each time we expand
		x += shift
		for y := 0; y < len(galaxyMap); y++ {
			galaxyMap[y] = append(galaxyMap[y][:x], append([]string{"."}, galaxyMap[y][x:]...)...)
		}
		shift += 1
	}
	return galaxyMap
}

func getGalaxiesLocations(galaxyMap [][]string) [][]int {
	var locations [][]int

	for y := 0; y < len(galaxyMap); y++ {
		for x := 0; x < len(galaxyMap[y]); x++ {
			if galaxyMap[y][x] == "#" {
				locations = append(locations, []int{x, y})
			}
		}
	}

	return locations
}

func calculateDistance(firstGalaxy []int, secondGalaxy []int) int {

	horizontalDifference := max(firstGalaxy[0], secondGalaxy[0]) - min(firstGalaxy[0], secondGalaxy[0])
	verticalDifference := max(firstGalaxy[1], secondGalaxy[1]) - min(firstGalaxy[1], secondGalaxy[1])

	return horizontalDifference + verticalDifference
}
