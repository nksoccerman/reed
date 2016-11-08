package main

import (
	"fmt"
	"strconv"
)

func factit(n int) int{
	var i int = 1
	var sum int = 0
	for i <= n {
		sum = sum + i
		i = i + 1
	}
	return sum
}

func facttrh(total int, n int) int {
	if n != 0 {
		return facttrh(total+n, n-1)
	} else {
		return total
	}
}

func facttr(n int) int {
	return facttrh(0, n)
}


func factrec(n int) int {
	if n != 0 {
		return n + factrec(n-1)
	} else {
		return 0
	}
}


func main() {
        var s string = ""
        var x int = 0

        fmt.Println("Enter an integer: ")
        fmt.Scanln(&s)
        x, _ = strconv.Atoi(s)
	fmt.Println("The result of three factorial functions: ")
	fmt.Printf("Loop-based = %d\n", factit(x))
	fmt.Printf("Tail-recursive = %d\n", facttr(x))
	fmt.Printf("Single-recursive = %d\n", factrec(x))
}
