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
	var verticalSeek []string
	var horizontalSeek []string
	var res int
	/*
	  we had an extra newline at the end of the text as we rely on empty lines
	  to separate patterns.
	*/
	text += "\n\n"
	scanner := bufio.NewScanner(strings.NewReader(text))

	for scanner.Scan() {

		part := scanner.Text()

		if part != "" {
			horizontalSeek = append(horizontalSeek, part)
		} else if part == "" && len(horizontalSeek) != 0 {
			// we reached the end of the part, lets generate the horizontalSeek
			for col := 0; col < len(horizontalSeek[0]); col++ {
				var word string
				for _, line := range horizontalSeek {
					word += string(line[col])
				}
				verticalSeek = append(verticalSeek, word)
			}

			horizontalMirror, verticalMirror := findMirror(horizontalSeek), findMirror(verticalSeek)

			if verticalMirror != -1 && horizontalMirror == -1 { // if the mirror is vertical
				res += verticalMirror
			} else if horizontalMirror != -1 && verticalMirror == -1 { // if the mirror is horizontal
				res += horizontalMirror * 100
			}

			verticalSeek, horizontalSeek = []string{}, []string{}
		}
	}
	return res
}

func findMirror(pattern []string) int {
	var potentialMirrors []int
	var previous string

	for i, line := range pattern {
		if line == previous {
			potentialMirrors = append(potentialMirrors, i-1)
		} else {
			previous = line
		}
	}

	for _, potentialMirror := range potentialMirrors {

		left, right := potentialMirror, potentialMirror+1

		for left >= 0 && right < len(pattern) {

			if pattern[left] == pattern[right] {
				left--
				right++
			} else {
				// if left and right are not eequal, this is not a mirror
				potentialMirror = -1
				break
			}
		}
		if potentialMirror != -1 {
			return potentialMirror + 1
		}
	}
	return -1
}
