package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	name           int
	items          []int64
	operation      []string
	test           int64
	trueAction     int
	falseAction    int
	inspectedItems int
}

var modulo int64 = 1

func (monkey *Monkey) inspectAndThrow(gang map[int]Monkey, item int64) {
	var targetMonkey Monkey

	// operation
	switch monkey.operation[0] {
	case "+":
		if monkey.operation[1] != "old" {
			n, err := strconv.Atoi(monkey.operation[1])
			if err != nil {
				panic(err)
			}
			item = item + int64(n)
			fmt.Printf("Worry level is added by %v to %v.\n", n, item)

		} else if monkey.operation[1] == "old" {
			fmt.Printf("Worry level is added by  by %v to %v.\n", item, item+item)
			item = item + item
		}
	case "*":
		if monkey.operation[1] != "old" {
			n, err := strconv.Atoi(monkey.operation[1])
			if err != nil {
				panic(err)
			}
			item = item * int64(n)
			fmt.Printf("Worry level is multiplied by %v to %v.\n", n, item)
		} else if monkey.operation[1] == "old" {
			fmt.Printf("Worry level is multiplied by %v to %v.\n", item, item*item)
			item = item * item

		}
	}

	// we only compare via division, so any number divisible by product of tests will work (see modular arithmetic)
	item = item % modulo

	monkey.inspectedItems += 1

	// test & throw
	if item%monkey.test == 0 {
		fmt.Printf("Current worry level is divisible by %v.\n", monkey.test)
		targetMonkey = gang[monkey.trueAction]

	} else {
		fmt.Printf("Current worry level is not divisible by %v.\n", monkey.test)
		targetMonkey = gang[monkey.falseAction]
	}
	targetMonkey.items = append(targetMonkey.items, int64(item))
	monkey.items = monkey.items[:len(monkey.items)-1]
	gang[monkey.name] = *monkey
	gang[targetMonkey.name] = targetMonkey
	fmt.Printf("Item with worry level %v is thrown to monkey %v.\n", item, targetMonkey.name)
}

func parseConfig(scanner *bufio.Scanner, gang map[int]Monkey) {
	monkey := 0
	items := make([]int64, 0)
	operation := make([]string, 0)
	test := 0
	trueAction := 0
	falseAction := 0

	for scanner.Scan() {
		s := strings.Split(strings.TrimSpace(scanner.Text()), " ")

		switch s[0] {
		case "Monkey":
			fmt.Sscanf(s[1], "%v:", &monkey)
		case "Starting":
			for i := 2; i < len(s); i++ {
				var n int
				fmt.Sscanf(s[i], "%v", &n)
				items = append(items, int64(n))
			}
		case "Operation:":
			var operator string
			var n string
			fmt.Sscanf(strings.TrimSpace(strings.Join(s[1:], "")), "new=old%1s%s", &operator, &n)
			operation = append(operation, operator, n)
		case "Test:":
			test, _ = strconv.Atoi(s[3])
		case "If":
			if s[1] == "true:" {
				trueAction, _ = strconv.Atoi(s[5])
			} else if s[1] == "false:" {
				falseAction, _ = strconv.Atoi(s[5])
				gang[monkey] = Monkey{name: monkey, items: items, operation: operation, test: int64(test), trueAction: trueAction, falseAction: falseAction, inspectedItems: 0}
			}
		default: // when we encounter a blank, we reset all variables and start the next monkey
			monkey = 0
			items = make([]int64, 0)
			operation = make([]string, 0)
			test = 0
			trueAction = 0
			falseAction = 0
		}
	}
}

func main() {
	gang := make(map[int]Monkey)
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	parseConfig(bufio.NewScanner(file), gang)

	// generate product of all tests
	for i := 0; i < len(gang); i++ {
		modulo *= gang[i].test
	}

	for round := 0; round < 10000; round++ {
		for i := 0; i < len(gang); i++ {
			monkey := gang[i]
			fmt.Printf("Monkey %v:\n", monkey.name)
			for _, item := range monkey.items {
				fmt.Printf("Monkey inspects an item with a worry level of %v.\n", item)
				monkey.inspectAndThrow(gang, int64(item))
			}
		}
	}

	best := 0
	second := 0
	for i := 0; i < len(gang); i++ {
		fmt.Printf("Monkey %v inspected items %v times.\n", gang[i].name, gang[i].inspectedItems)

		if gang[i].inspectedItems > best {
			second = best
			best = gang[i].inspectedItems
		} else if gang[i].inspectedItems > second && gang[i].inspectedItems < best {
			second = gang[i].inspectedItems
		}
	}
	fmt.Println("Monkeybusiness: ", best*second)

}
