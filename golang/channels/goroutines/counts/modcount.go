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


func P(c *counter) {
	for {
		c.increment()
		time.Sleep(time.Millisecond * time.Duration(10+rand.Intn(100)))
	}
}

func Q(c *counter) {
	var last int
	var next int = c.read()
	for {
		fmt.Printf("count of %d\n",next)
		time.Sleep(time.Millisecond*100)
		last = next
		next = c.read()
		for next == last {
			last = next
			next = c.read()
			time.Sleep(time.Millisecond*100)
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	var ctr counter = counter{value:0}
	fmt.Printf("Starting...\n")
	go P(&ctr)
	go Q(&ctr)
	for {
		time.Sleep(time.Second)
	}
}
