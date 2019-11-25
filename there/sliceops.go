package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("slice=%v len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, 2*i+1)
		printSlice(s)
	}
	fmt.Println(s)

	s1 := []int{3, 4, 5, 6}

	s2 := make([]int, 16)

	s3 := make([]int, 10, 30)

	printSlice(s1)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copy Slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("delete elements from slice: \n delect index at 3")
	s4 := append(s2[:3], s2[4:]...)
	printSlice(s4)

	fmt.Printf("Poping font")
	//front := s2[0]

}
