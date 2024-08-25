package main

import (
	"fmt"
)

func twoSum(nums *[]int, target *int) []int {
	dict := make(map[int]int)
	result := make([]int, 2)
	for i, num := range *nums {
		if val, ok := dict[*target-num]; ok {
			result[0] = val
			result[1] = i
			break
		}
		dict[num] = i
	}
	return result
}

func main() {
	testCase1()
	testCase2()
	testCase3()
}

func testCase1() {
	var nums = []int{2, 7, 11, 15}
	var target = 9
	fmt.Println(twoSum(&nums, &target))
	// Output: [0 1]
}

func testCase2() {
	var nums = []int{3, 2, 4}
	var target = 6
	fmt.Println(twoSum(&nums, &target))
	// Output: [1 2]
}

func testCase3() {
	var nums = []int{1, 2, 3, 4, 5}
	var target = 5
	fmt.Println(twoSum(&nums, &target))
	// Output: [1 2]
}
