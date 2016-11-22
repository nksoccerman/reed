package main

import (
	"fmt"
	"math/rand"
	"time"
)

type labelint struct {
	label string
	value int
}

func shout(name string, c chan labelint) {
	var count int = 0
	for {
		c <- labelint{name,count}
		count = count + 1
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func echo(c chan labelint) {
	for {
		var msg labelint = <-c
		fmt.Println(msg.label, ":", msg.value)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func balance(i1, i2 chan labelint, o1,o2 chan labelint) {
	var msg labelint
	var bool status = false 
	for {
		var msg labelint
		select {
		case msg = (<- i1)
			if !status {
				o1 <- i1
			} else {
				o2 <- i1
			}
		case msg = (<- i2)
			if !status 
				o1 <- i2
			} else {
				o2 <- i2
			}
		status = !status
	}
}

func main() {
	var c chan labelint = make(chan labelint)
	go shout("joe",c)
	go shout("bob",c)
	go echo(c)
	go echo(c) 
	for {
		time.Sleep(time.Second)
	}
}

