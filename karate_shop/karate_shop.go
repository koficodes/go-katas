package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(binarySearch([]int{9, 6, 3, 5, 7, 1, 2, 4, 10}))
}

func binarySearch(collection []int) int {
	sort.Ints(collection)
	lowerbound := 1
	upperbound := len(collection)

	for index, value := range collection {

		if upperbound < lowerbound {
			return -1
		}

		midValue := getMidValue(lowerbound, upperbound)

	}
	return collection
}

func getMidValue(lowerbound int, upperbound int) int {
	return lowerbound + (upperbound-lowerbound)/2
}
