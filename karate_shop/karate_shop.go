package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(chop(5, []int{9, 6, 3, 5, 7, 1, 2, 4, 10}))
}

func chop(target int, collection []int) int {
	sort.Ints(collection)
	lowerbound := 1
	upperbound := len(collection)

	for {

		if upperbound < lowerbound {
			return -1
		}

		midValue := lowerbound + ((upperbound - lowerbound) / 2)

		if collection[midValue] < target {
			lowerbound = midValue + 1
		}

		if collection[midValue] > target {
			upperbound = midValue - 1
		}

		if collection[midValue] == target {
			return midValue
		}
	}
}
