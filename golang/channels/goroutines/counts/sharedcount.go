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
	var i int = c.value
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(50)))
    c.value = i + 1
}

func (c *counter) read() int {
    return c.value
}


func P(n int, c *counter) {
	var i int = 0
	for i < n {
		c.increment()
		time.Sleep(time.Microsecond * time.Duration(rand.Intn(50)))
		i = i+1
	}
	fmt.Printf("Done.\n")
}

func main() {
	rand.Seed(time.Now().Unix())
	var ctr counter = counter{value:0}
	fmt.Printf("Starting...\n")
	go P(10000, &ctr)
	go P(10000, &ctr)
	time.Sleep(time.Second)
	fmt.Printf("Counter is %d\n", ctr.read())
}
