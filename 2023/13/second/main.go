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
	var horizontalSeek []string
	var res int
	var patternCount int
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
			log.Printf("*** Pattern %v ***\n", patternCount)
			patternCount++
			// we reached the end of the part, lets generate the horizontalSeek
			res += computeSolution(horizontalSeek)
			horizontalSeek = []string{}
		}
	}
	return res
}

func findMirror(pattern []string) int {
	var potentialMirrors []int

	// edge case : first two elements of pattern, as we only have one mirrored line, it absolutely needs to contain smudge
	if findSmudge(string(pattern[0]), string(pattern[1])) {
		potentialMirrors = append(potentialMirrors, 0)
	}

	l2, l1, r1, r2 := 0, 1, 2, 3

	for r2 <= len(pattern)-1 {
		if pattern[l2] == pattern[r2] && pattern[l1] == pattern[r1] {
			potentialMirrors = append(potentialMirrors, l1)
		} else if pattern[l2] == pattern[r2] && findSmudge(string(pattern[l1]), string(pattern[r1])) {
			potentialMirrors = append(potentialMirrors, l1)
		} else if pattern[l1] == pattern[r1] && findSmudge(string(pattern[l2]), string(pattern[r2])) {
			potentialMirrors = append(potentialMirrors, l1)
		} // this is not a mirror line, slide window
		l2, l1, r1, r2 = l2+1, l1+1, r1+1, r2+1
	}

	// edge case last two lines of pattern, as we only have one mirrored line, it absolutely needs to contain smudge
	if findSmudge(string(pattern[len(pattern)-1]), string(pattern[len(pattern)-2])) {
		potentialMirrors = append(potentialMirrors, len(pattern)-1)
	}

	log.Println(potentialMirrors)
	// check if the mirror is indeed a real one
	for _, potentialMirror := range potentialMirrors {
		foundSmudge := false
		left, right := potentialMirror, potentialMirror+1

		for left >= 0 && right < len(pattern) {
			if pattern[left] == pattern[right] {
				left--
				right++
			} else if pattern[left] != pattern[right] && findSmudge(pattern[left], pattern[right]) && !foundSmudge {
				// if left and right are not equal and we haven't found the smudge, and there is a smudge
				foundSmudge = true
				left--
				right++
			} else {
				// if left and right are not equal and its not a smudge, or we have already found the smudge, this is not a mirror
				potentialMirror = -1
				break
			}
		}
		if potentialMirror != -1 && foundSmudge {
			return potentialMirror + 1
		}
	}
	return -1
}

func computeSolution(horizontalSeek []string) int {
	var verticalSeek []string
	var res int

	// make the pattern vertical to easily go through it
	for col := 0; col < len(horizontalSeek[0]); col++ {
		var word string
		for _, line := range horizontalSeek {
			word += string(line[col])
		}
		verticalSeek = append(verticalSeek, word)
	}

	horizontalMirror, verticalMirror := findMirror(horizontalSeek), findMirror(verticalSeek)

	if verticalMirror != -1 && horizontalMirror == -1 { // if the mirror is vertical
		log.Printf("vertical: %v\n", verticalMirror)
		res += verticalMirror
	} else if horizontalMirror != -1 && verticalMirror == -1 { // if the mirror is horizontal
		log.Printf("horizontal: %v\n", horizontalMirror)
		res += horizontalMirror * 100
	} else {
		log.Println("!!!!!!!!! MISSING MIRROR !!!!!!!!!!!")
		for i, v := range horizontalSeek {
			if i >= 10 {
				log.Printf("%v %v\n", i, v)
			} else if i < 10 {
				log.Printf("0%v %v\n", i, v)
			}
		}
	}

	return res
}

func findSmudge(string1 string, string2 string) bool {
	smudgeCount := 0

	if len(string1) != len(string2) || len(string1) == 0 {
		return false
	}

	for i := 0; i < len(string1); i++ {
		if string1[i] != string2[i] {
			if (string(string1[i]) == "#" && string(string2[i]) == ".") ||
				(string(string1[i]) == "." && string(string2[i]) == "#") {
				if smudgeCount >= 1 {
					return false
				} else {
					smudgeCount++
				}
			}
		}
	}

	return true
}
