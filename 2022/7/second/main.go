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

	scanner := bufio.NewScanner(file)
	currentPwd := make([]string, 0)
	hashmap := make(map[string]int)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		if s[0] == "$" { // input of a command
			if s[1] == "cd" {
				if s[2] == ".." {
					currentPwd = currentPwd[:len(currentPwd)-1] //remove last currentPwd
				} else {
					currentPwd = append(currentPwd, s[2])
				}
			}
		} else { // output of a command
			if s[0] == "dir" {
				// Pretty sure this is useless information, we'll cd in it at some point anyway
			} else {
				var pwd string
				size, err := strconv.Atoi(s[0])

				if err != nil {
					panic(err)
				}

				if len(currentPwd) > 1 {
					pwd = strings.Join(currentPwd[:1], "/")
				} else if len(currentPwd) == 1 {
					pwd = currentPwd[0]
				}

				if _, exist := hashmap[pwd]; exist {
					for i := len(currentPwd); i >= 1; i-- {
						hashmap[strings.Join(currentPwd[0:i], "/")] += size

					}
				} else {
					hashmap[pwd] = size
				}
			}
		}
	}
	result := 0
	for _, v := range hashmap {
		if v < 100000 {
			result += v
		}
	}
	fmt.Println("part 1:", result)

	diskSize := 70000000
	reqSpace := 30000000
	need := reqSpace - (diskSize - hashmap["/"])
	result = 0

	for _, v := range hashmap {
		if v >= need {
			if result > v || result == 0 {
				result = v
			}
		}
	}

	fmt.Println("part 2:", result)
}
