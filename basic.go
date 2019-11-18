package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//var aa = 3
//var ss = "kkk"
var (
	aa = 3
	ss = "kkk"
)

func variableZeroValue() {
	var a int
	var s string

	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, false, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println("Hello world")
	fmt.Printf("packege variable %d %s\n", aa, ss)

	euler()
	triangle()
	conses()
	enums()
}

func euler() {
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)

	fmt.Printf("%.3f \n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a) + float64(b*b)))
	fmt.Println(c)
}

func conses() {
	const (
		filename string = "abc.txt"
		a, b            = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {

	const (
		cpp = iota
		java
		python
		golang
		_
		javaScript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, python, golang, javaScript)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func flow() {

}
