package main

import (
	"fmt"
	"strconv"
)

func fibit(n int) int {
	var i int = 0
	var f0 int = 0
	var f1 int = 1
	for i < n {
		f0, f1 = f1, f0+f1
		i = i + 1
	}
	return f0
}

func fibtrh(n int, f0 int, f1 int) int {
	if n == 0 {
		return f0
	} else {
		return fibtrh(n-1, f1, f0+f1)
	}
}

func fibtr(n int) int {
	return fibtrh(n, 0, 1)
}

func fibrec(n int) int {
	if n < 2 {
		return n
	} else {
		return fibrec(n-2) + fibrec(n-1)
	}
}

func main() {
	var s string = ""
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&s)
	var n int = 0
	n, _ = strconv.Atoi(s)
	fmt.Printf("fib(%d) = %d\n", n, fibrec(n))
	fmt.Printf("fib(%d) = %d\n", n, fibtr(n))
	fmt.Printf("fib(%d) = %d\n", n, fibit(n))
}
