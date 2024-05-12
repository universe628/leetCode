package main

import (
	"fmt"

	practice "github.com/unvierse628/leetCode/practice"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	slice = practice.SliceInsert(slice, 5, 8)
	fmt.Print(slice)
	fmt.Print(len(slice))
}
