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


func P(p int, c *counter) {
	for {
		var i int = c.read()
		fmt.Printf("%d: count of %d\n",p,i)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	var ctr counter = counter{value:0}
	fmt.Printf("Starting...\n")
	go P(1,&ctr)
	go P(2,&ctr)
	for {
		time.Sleep(time.Millisecond * 200)
		ctr.increment()
	}
}
