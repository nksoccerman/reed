package main

import (
	"fmt"
	"strconv"
)

type listNode struct {
	value int
	next *listNode
}

type pqueue struct {
	first *listNode
	last *listNode
}

func makeQueue() *pqueue {
	return &pqueue{first: nil, last: nil}
}

func (Q *pqueue) enqueue(x int) {
	var n *listNode = &listNode{value: x, next: nil}
	if Q.isEmpty() {
		Q.first = n
		Q.last = n
	} else {
		var next *listNode = Q.first.next
		var prev *listNode = Q.first
		for (next != nil && next.value > x) {
			prev = next
			next = next.next
		}
		if prev.value < x {
			Q.first = n
			n.next = prev
		} else{
			prev.next = n
			n.next = next
		}
	}
}

func (Q *pqueue) head() int {
	return Q.first.value
}

func (Q *pqueue) dequeue() int {
	if Q.isEmpty() {
		return 666
	} else if Q.first.next == nil {
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

func (Q *pqueue) isEmpty() bool {
	if Q.first == nil {
		return true
	} else {
		return false
	}
}

func (Q *pqueue) String() string {
	if Q.isEmpty() {
		return ""
	} else {	
		var str string = strconv.Itoa(Q.first.value)
		var ref *listNode = Q.first.next
		for ref != nil {
			var val = strconv.Itoa(ref.value)
			str = str + " " + val
			ref = ref.next
		}
		return str
	}
	
}

func main() {
	var q *pqueue = makeQueue()
	q.enqueue(10)
	q.enqueue(5)
	q.enqueue(26)
	fmt.Println(q)
	fmt.Println(q.dequeue())
	fmt.Println(q)
	fmt.Println(q.head())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.isEmpty())
}
