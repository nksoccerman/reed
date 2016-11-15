package main

import (
	"fmt"
	"strconv"
)

type listNode struct {
	value int
	next *listNode
}

type queue struct {
	first *listNode
	last *listNode
}

func makeQueue() *queue {
	return &queue{first: nil, last: nil}
}

func (Q *queue) enqueue(x int) {
	var n *listNode = &listNode{value: x, next: nil}
	if Q.isEmpty() {
		Q.first = n
		Q.last = n
	} else {
		Q.last.next = n
		Q.last = n
	}
}

func (Q *queue) head() int {
	return Q.first.value
}

func (Q *queue) dequeue() int {
	if Q.isEmpty() {
		return 666
	} else if Q.first == Q.last {
		var x int = Q.first.value
		Q.first = nil
		Q.last = nil
		return x
	} else {
		var x int = Q.first.value
		Q.first = Q.first.next
		return x
	}
}

func (Q *queue) isEmpty() bool {
	if (Q.first == nil && Q.last == nil) {
		return true
	} else {
		return false
	}
}

func (Q *queue) String() string {
	if Q.isEmpty() {
		return ""
	} else {	
		var str string = ""
		var ref *listNode = Q.first
		for ref != nil {
			var val = strconv.Itoa(ref.value)
			str = str + " " + val
			ref = ref.next
		}
		return str
	}
	
}

func main() {
	var q *queue = makeQueue()
	q.enqueue(1)
	q.enqueue(5)
	q.enqueue(26)
	fmt.Println(q)
	fmt.Println(q.dequeue())
	q.dequeue()
	fmt.Println(q.head())
	q.dequeue()
	fmt.Println(q.isEmpty())
}
