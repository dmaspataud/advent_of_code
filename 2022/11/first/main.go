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
	items          []int
	operation      []string
	test           int
	trueAction     int
	falseAction    int
	inspectedItems int
}

func (monkey *Monkey) inspectAndThrow(gang map[int]Monkey, item int) {
	var targetMonkey Monkey

	// operation
	switch monkey.operation[0] {
	case "+":
		if monkey.operation[1] != "old" {
			n, err := strconv.Atoi(monkey.operation[1])
			if err != nil {
				panic(err)
			}
			item = item + n
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
			item = item * n
			fmt.Printf("Worry level is multiplied by %v to %v.\n", n, item)
		} else if monkey.operation[1] == "old" {
			fmt.Printf("Worry level is multiplied by %v to %v.\n", item, item*item)
			item = item * item

		}
	}

	monkey.inspectedItems += 1

	// Monkey get bored of the item and worry goes down
	item = item / 3
	fmt.Printf("Monkey gets bored with item. Worry level is divided by 3 to %v.\n", item)

	// test & throw
	if item%monkey.test == 0 {
		fmt.Printf("Current worry level is divisible by %v.\n", monkey.test)
		targetMonkey = gang[monkey.trueAction]

	} else {
		fmt.Printf("Current worry level is not divisible by %v.\n", monkey.test)
		targetMonkey = gang[monkey.falseAction]
	}
	targetMonkey.items = append(targetMonkey.items, item)
	monkey.items = monkey.items[:len(monkey.items)-1]
	gang[monkey.name] = *monkey
	gang[targetMonkey.name] = targetMonkey
	fmt.Printf("Item with worry level %v is thrown to monkey %v.\n", item, targetMonkey.name)
}

func parseConfig(scanner *bufio.Scanner, gang map[int]Monkey) {
	monkey := 0
	items := make([]int, 0)
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
				items = append(items, n)
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
				gang[monkey] = Monkey{name: monkey, items: items, operation: operation, test: test, trueAction: trueAction, falseAction: falseAction, inspectedItems: 0}
			}
		default: // when we encounter a blank, we reset all variables and start the next monkey
			monkey = 0
			items = make([]int, 0)
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

	for round := 0; round < 20; round++ {
		for i := 0; i < len(gang); i++ {
			monkey := gang[i]
			fmt.Printf("Monkey %v:\n", monkey.name)
			for _, item := range monkey.items {
				fmt.Printf("Monkey inspects an item with a worry level of %v.\n", item)
				monkey.inspectAndThrow(gang, item)
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
		}
	}
	fmt.Println("Monkeybusiness: ", best*second)

}
