package main

import "fmt"

func main() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	var i int8 = 127
	fmt.Println(i, i+1, i*i)

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\b\n", x)
	fmt.Printf("%08b\b\n", y)
}
