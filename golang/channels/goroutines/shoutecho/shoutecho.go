package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv" 
)

func shout(c chan int) {
	var count int = 0
	for {
		c <- count
		count = count + 1
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func echo(c chan int) {
	for {
		var count int = <-c
		fmt.Println(count)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func forward(p int, c_in chan int, c_out chan int) {
	var rec int = 0
	for {
		rec = <-c_in	
		fmt.Printf("#%d is fowarding message %d\n", p, rec)
		c_out <- rec
	}
}
func main() {
	var n int = 0 
	var i int = 0
	var s string = ""
	fmt.Println("Enter the number of forwarders: ")
	fmt.Scanln(&s)	 
	n, _ = strconv.Atoi(s) 

	var t chan int = make(chan int)
	var x chan int = make(chan int)
	go shout(t)
	var z bool = true
	for i < n {
		if z {
			x = make(chan int)
			go forward((i+1), t, x)
		} else {
			t = make(chan int)
			go forward((i+1), x, t) 
		}
		i = i + 1
		z = !z
	}
	if !z {
		go echo(x)
	} else {
		go echo(t)
	}
	for {
		time.Sleep(time.Second)
	}
}

