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

func echo(name string, c chan labelint) {
	for {
		var msg labelint = <-c
		fmt.Println(name, "<-", msg.label, ":", msg.value)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func balance(i1, i2 chan labelint, o1,o2 chan labelint) {
        var status bool = false
        for {
                var msg labelint
                select {
                case msg = (<- i1):
                        if !status {
                                o1 <- msg
                        } else {
                                o2 <- msg
                        }
                case msg = (<- i2):
                        if !status {
                                o1 <- msg
                        } else {
                                o2 <- msg
                        }
                status = !status
        	}
	}	
}

func main() {
	var c chan labelint = make(chan labelint)
	var x chan labelint = make(chan labelint)
	go shout("bela",c)
	go shout("bob",x)
	var c2 chan labelint = make(chan labelint)
	var x2 chan labelint = make(chan labelint)
	go balance(c, x, c2, x2)
	go echo("stu",c2)
	go echo("todd",x2)
	for {
		time.Sleep(time.Second)
	}
}

