package main

import (
	"bufio"
	"fmt"
	"os"
)

func init() {
	fmt.Println("init")
}

// main 方法需要在 main 包下面
//func main() {
//	fmt.Println("Hello World")
//	fmt.Println(os.Hostname())
//
//	var s, sep string
//	for i := 0; i < 5+len(os.Args); i++ {
//		s += sep + os.Args[0]
//		sep = " "
//	}
//
//	fmt.Println(s)
//}
//
//func main() {
//	s, sep := "", ""
//	for index, arg := range os.Args[0:] {
//		s += sep + arg
//		sep = " "
//		fmt.Println(index, arg)
//	}
//}

//
//func main() {
//	//fmt.Println(strings.Join(os.Args[0:], " "))
//	fmt.Println(os.Args[0:])
//}
//
//func main() {
//	// 返回一个map
//	counts := make(map[string]int)
//	input := bufio.NewScanner(os.Stdin)
//	for input.Scan() {
//		counts[input.Text()]++
//	}
//	for line, n := range counts {
//		if n > 1 {
//			fmt.Println(n, line)
//		}
//	}
//}

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Println(os.Stderr, "dup2:%v\n", err)
			}
			countLines(f, counts)
			f.Close()
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
