package main

import (
    "fmt"
    "time"
    "math/rand"
)

type counter struct {
    value int
}

func (c *counter) increment() {
    c.value = c.value + 1
}

func (c *counter) read() int {
    return c.value
}


func P(c chan int) {
	var ctr counter = counter{value:0}
	for {
		c <- ctr.read()
		ctr.increment()
		time.Sleep(time.Millisecond * time.Duration(10+rand.Intn(100)))
	}
}

func Q(c chan int) {
	for {
		fmt.Printf("count of %d\n",<-c)
		time.Sleep(time.Millisecond*100)
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Printf("Starting...\n")
	var c chan int = make(chan int)
	go P(c)
	go Q(c)
	for {
		time.Sleep(time.Second)
	}
}
