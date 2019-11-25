package main

import "fmt"

func main() {
	var arr1 [5]int
	//arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	var grid [4][5]int
	//fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	//
	//for i := 0; i < len(arr3); i++ {
	//	fmt.Println(arr3[i])
	//}
	//
	//for i, V := range arr3 {
	//	fmt.Println(i, V)
	//}

	printArray(arr1)
	printArray(arr3)
	//printArray(arr2)

	s := arr3[2:3]
	fmt.Println(s)
}

func printArray(arr [5]int) {
	for i, V := range arr {
		fmt.Println(i, V)
	}
}
