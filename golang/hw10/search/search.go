package main

import (
	"fmt"
	"strconv"
)

func contains(a []int, k int) bool {
	var l int = 0
	var r int = len(a) - 1
	for l <= r {
		var m int = (l + r) / 2
		if a[m] == k {
			return true
		} else if a[m] < k {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false
}

func output(a []int) string {
	var c bool = false
	var s string = "["
	for _, x := range a {
		if c {
			s = s + (", ")
		}
		c = true
		s = s + strconv.Itoa(x)
	}
	s = s + "]"
	return s
}

func sort(a []int) {
	var n int = len(a)
	var i int = 1
	for i < n {
		var j int = i
		for j > 0 {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
			j = j - 1
		}
		i = i + 1
	}
	return
}

func sum(a []int) int {
	var s int = 0
	var i int = 0
	var n int = len(a)
	for i < n {
		s = s + a[i]
		i = i + 1
	}
	return s
}

func input(prompt string) int {
	var s string = ""
	fmt.Printf("%s", prompt)
	fmt.Scanln(&s)
	var n int = 0
	n, _ = strconv.Atoi(s)
	return n
}

func main() {
	var n int = input("Enter an array size: ")
	var a []int = make([]int, n)
	var i int = 0
	for i < n {
		var s string = "Enter integer #" + strconv.Itoa(i+1) + ": "
		a[i] = input(s)
		i = i + 1
	}

	fmt.Printf("That array is %s.\n", output(a))
	fmt.Printf("The sum of that array is %d.\n", sum(a))
	sort(a)
	fmt.Printf("That array sorted is %s.\n", output(a))
	i = 0
	for i < 10 {
		if contains(a, i) {
			fmt.Printf("That array contains %d.\n", i)
		} else {
			fmt.Printf("That array does not contain %d.\n", i)
		}
		i = i + 1
	}
}
