package main

import (
	"fmt"
	"strconv"
)

func mul(x int, y int) int {
	var s int = 0
	for y > 0 {
		if y%2 == 1 {
			s = s + x
		}
		x = x * 2
		y = y / 2
	}
	return s
}

func powerit2(x int, n int) int {
	var z int = x
	var p int = 1
	for n > 0 {
		if n%2 == 1 {
			p = p * z
		}
		z = z * z
		n = n / 2
	}
	return p
}

func powerec2(x int, n int) int {
	if n == 0 {
		return 1
	} else if n%2 == 0 {
		var z int = powerec2(x, n/2)
		return z * z
	} else {
		return x * powerec2(x, n-1)
	}
}

func powerec(x int, n int) int {
	if n == 0 {
		return 1
	} else {
		return x * powerec(x, n-1)
	}
}

func powerit(x int, n int) int {
	var p int = 1
	var i int = 0
	for i < n {
		p = p * x
		i = i + 1
	}
	return p
}

func main() {
	var s string = ""
	var x int = 0
	var n int = 0

	fmt.Println("Enter an integer: ")
	fmt.Scanln(&s)
	x, _ = strconv.Atoi(s)
	fmt.Println("Enter another integer: ")
	fmt.Scanln(&s)
	n, _ = strconv.Atoi(s)

	fmt.Printf("power(%d,%d) = %d\n", x, n, powerit(x, n))
	fmt.Printf("power(%d,%d) = %d\n", x, n, powerec(x, n))
	fmt.Printf("power(%d,%d) = %d\n", x, n, powerec2(x, n))
	fmt.Printf("power(%d,%d) = %d\n", x, n, powerit2(x, n))
	fmt.Printf("%d * %d = %d\n", x, x, mul(x, x))
}
