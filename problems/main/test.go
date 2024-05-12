package main

import (
	"fmt"

	problem "github.com/unvierse628/leetCode/problems"
)

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	unMatched := problem.RemoveElement(nums, val)
	fmt.Print(unMatched, nums)
}
