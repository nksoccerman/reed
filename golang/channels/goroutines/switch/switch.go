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

func shout(name string, c chan<- labelint) {
	var count int = 0
	for {
		c <- labelint{name,count}
		count = count + 1
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func selector(sw <- chan int, c1 <-chan labelint, c2 <-chan labelint) {
	var which int = <-sw
	var cs []<-chan labelint = make([]<-chan labelint,3)
	cs[1] = c1
	cs[2] = c2
	for {
		var msg labelint
		select {
		case which = <-sw:
		case msg = <-cs[which]:
			fmt.Println("Channel", which, "received", msg.value, "from", msg.label)
		}
	}
}

func main() {
	var sw chan int = make(chan int)
	var c1 chan labelint = make(chan labelint)
	var c2 chan labelint = make(chan labelint)
	go shout("joe.",c1)
	go shout("bob.",c2)
	go selector(sw,c1,c2)
	var which = 1
	for {
		sw <- which
		which = 3 - which
		time.Sleep(time.Second)
	}
}

