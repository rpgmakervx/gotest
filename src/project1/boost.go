package main

import (
	"fmt"
	"project1/sort"
)

func main() {
	array := make([]int, 0, 10)
	array = append(array, 4, 3, 2, 5, 7, 9, 8, 1, 0, 6)
	fmt.Println("array before sorted:", array)
	sort.QSort(array)
	fmt.Println("array sorted:", array)
}
