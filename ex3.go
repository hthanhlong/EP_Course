/// Viết chương trình nhập 1 slice số, in ra tổng, số lớn nhất, số nhỏ nhất, trung bình cộng, slice đã được sắp xếp (tăng/giảm)

package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	var s []int
	var n int
	var x int
	fmt.Print("Enter the number of elements in the slice: ")
	fmt.Scan(&n)
	fmt.Println("Enter elements of the slice: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		s = append(s, x)
	}
	fmt.Printf("Sum: %.3f\n", sliceSum(&s))
	fmt.Printf("Maximum: %d\n", slices.Max(s))
	fmt.Printf("Minimum: %d\n", slices.Min(s))
	fmt.Printf("Average: %.3f\n", sliceSum(&s)/float64(len(s)))
	fmt.Println("Sorted in ascending order:", sortSlice("de5sc", &s))

}

func sliceSum(s *[]int) float64 {
	var total float64
	for _, v := range *s {
		total += float64(v)
	}
	return total
}

func sortSlice(order string, s *[]int) []int {
	if order != "asc" && order != "desc" {
		fmt.Println("Invalid order")
		return []int{}
	}
	sort.Slice(*s, func(i, j int) bool {
		if order == "desc" {
			return (*s)[i] > (*s)[j]
		}
		return (*s)[i] < (*s)[j]
	})
	return *s
}
