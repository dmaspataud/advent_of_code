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

	galaxyMap := parse(text)

	horizontalExpansions, verticalExpansions := applyUniverseExpansion(galaxyMap)

	locations := getGalaxiesLocations(galaxyMap)

	for first := 0; first < len(locations); first++ {
		for second := 0; second < len(locations); second++ {
			res += calculateDistance(locations[first], locations[second], horizontalExpansions, verticalExpansions)
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

	return galaxyMap
}

func applyUniverseExpansion(galaxyMap [][]string) ([]int, []int) {
	horizontalExpansions := []int{}
	verticalExpansions := []int{}

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
			verticalExpansions = append(verticalExpansions, y)
		}
	}

	// vertical expansion
	for x := 0; x < len(galaxyMap[0]); x++ {
		isGalaxy := false
		for y := 0; y < len(galaxyMap); y++ {
			if galaxyMap[y][x] == "#" {
				isGalaxy = true
				break
			}
		}
		if !isGalaxy {
			horizontalExpansions = append(horizontalExpansions, x)
		}
	}
	return horizontalExpansions, verticalExpansions
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

func calculateDistance(firstGalaxy []int, secondGalaxy []int, horizontalExpansions []int, verticalExpansions []int) int {
	//horizontal

	horizontalExpansionFields := 0

	for _, eachExpansion := range horizontalExpansions {
		if eachExpansion < max(firstGalaxy[0], secondGalaxy[0]) && eachExpansion > min(firstGalaxy[0], secondGalaxy[0]) {
			horizontalExpansionFields += 1
		}
	}

	horizontalDifference := max(firstGalaxy[0], secondGalaxy[0]) - min(firstGalaxy[0], secondGalaxy[0]) + (horizontalExpansionFields * 999999)

	// vertical

	verticalExpansionFields := 0

	for _, eachExpansion := range verticalExpansions {
		if eachExpansion < max(firstGalaxy[1], secondGalaxy[1]) && eachExpansion > min(firstGalaxy[1], secondGalaxy[1]) {
			verticalExpansionFields += 1
		}
	}

	verticalDifference := max(firstGalaxy[1], secondGalaxy[1]) - min(firstGalaxy[1], secondGalaxy[1]) + (verticalExpansionFields * 999999)

	return horizontalDifference + verticalDifference
}
