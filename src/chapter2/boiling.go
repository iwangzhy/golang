package main

import (
	"fmt"
)

const boilingF = 212.0

func main() {
	//var f = boilingF
	//var c = (f - 32) * 5 / 9
	//fmt.Printf("boiling point = %g F or %g C\n", f, c)
	//fmt.Printf("%g F = %g C\n", f, fToC(boilingF))
	defineVariable()
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func defineVariable() {
	var s string
	fmt.Printf("@%s@\n", s)

	var a, b, c int
	fmt.Printf("a:%d b:%d c:%d \n", a, b, c)

	var d, e, f = true, 0, "def"
	fmt.Printf("%t %d %s\n", d, e, f)
	//
	//var file, err = os.Open("a")
	//if err != nil {
	//	fmt.Printf("%s", file.Name())
	//}

	var z, x int
	fmt.Println(&x == &x, &z == &x, &x == nil, x == 0)
	var v = 1
	incr(&v)
	fmt.Println(incr(&v))

}

func incr(p *int) int {
	*p++
	return *p
}
