package main

import "fmt"

func updateArray(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	fmt.Println("arr[2:] =", arr[2:])
	fmt.Println("arr[:] =", arr[:])

	s1 := arr[2:6]
	s2 := arr[:]

	fmt.Println("before:")
	fmt.Println(s1) //[2,3,4,5]
	fmt.Println(s2) //[0, 1, 2, 3, 4, 5, 6, 7]

	updateArray(s1)

	fmt.Println("after:")
	fmt.Println(s1) //[100,3,4,5]
	fmt.Println(s2) //[0, 1, 100, 3, 4, 5, 6, 7]

	s3 := arr[3:7]
	fmt.Printf("s3=%v, len(s3)=%d, cap(s3)=%d\n", s3, len(s3), cap(s3))
	//s4 := s3[3:5]
	//fmt.Println(s3)
	//fmt.Println(s4)
	s4 := append(s3, 9)
	s5 := append(s4, 11)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(arr)
}
