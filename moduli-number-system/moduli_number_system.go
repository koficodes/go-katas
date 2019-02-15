package main

import (
	"fmt"
	"strconv"
	"strings"
)

//FromNbToStr converts number to string of moduli number system
func FromNbToStr(n int64, sys []int64) string {
	var residue []int64
	for _, value := range sys {
		residue = append(residue, n%value)
	}

	var multiSum int64 = 1

	for _, num := range sys {
		multiSum *= num
	}

	if multiSum < n {
		return "Not applicable"
	}

	//fmt.Println(residue)
	for _, res1 := range residue {
		for _, res2 := range residue {
			fmt.Println(res1, res2, residue)
			if res1 == res2 {
				continue
			}
			if !areCoPrimes(res1, res2) {
				return "Not applicable"
			}
		}

	}

	return "-" + strings.Join(Map(residue, strconv.Itoa), "--") + "-"
}

func makeRange(start int64, end int64) []int64 {
	var rangeValues []int64
	var i int64
	for i <= end {
		rangeValues = append(rangeValues, start+i)
		i++
	}
	return rangeValues
}

//Map applies a function a slice
func Map(sliceOfTypes []int64, applyFunc func(item int) string) []string {
	var bucket []string
	for _, item := range sliceOfTypes {
		bucket = append(bucket, applyFunc(int(item)))
	}
	return bucket
}

func getFactors(num int64) []int64 {
	var factors []int64
	for _, divider := range makeRange(1, num+1) {
		if num%divider == 0 {
			factors = append(factors, divider)
		}
	}

	return factors
}

func areCoPrimes(res1 int64, res2 int64) bool {
	var commonFactors []int64
	for _, factor := range getFactors(res1) {

		if factorExists(factor, getFactors(res2)) {
			commonFactors = append(commonFactors, factor)
		}
	}

	if len(commonFactors) > 1 || commonFactors[0] != 1 {
		return false
	}

	return true
}

func factorExists(num int64, factors []int64) bool {

	for _, factor := range factors {
		if num == factor {
			return true
		}
	}
	return false
}

func main() {

	fmt.Println(FromNbToStr(779, []int64{8, 7, 5, 3}))

	fmt.Println(FromNbToStr(187, []int64{8, 7, 5, 3}))

	fmt.Println(FromNbToStr(15, []int64{8, 6, 5}))

	fmt.Println(FromNbToStr(3450, []int64{17, 5, 3}))

}
