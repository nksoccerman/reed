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
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	}
}

func multiplex(c1 chan labelint, c2 chan labelint) {
	for {
		var which int
		var msg labelint

		select {
		case msg = <-c1:
			which = 1
		case msg = <-c2:
			which = 2
		default:
			which = 0
		}

		if which > 0 {
			fmt.Println("Channel", which, "received", msg.value, "from", msg.label)
		} else {
			fmt.Println("Idle.")
			time.Sleep(time.Duration(time.Second))
		}
	}
}

func main() {
	var c1 chan labelint = make(chan labelint)
	var c2 chan labelint = make(chan labelint)
	go shout("joe.",c1)
	go shout("bob.",c2)
	go multiplex(c1,c2)
	for {
		time.Sleep(time.Second)
	}
}

