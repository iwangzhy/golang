package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func main() {
	fmt.Println("1 add 2 =", Add(1, 2))
}
