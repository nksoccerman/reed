package main

import (
	"fmt"
	"strconv"
)

type listNode struct {
	value int
	next  *listNode
}

type stack struct {
	first *listNode
}

func makeStack() (*stack) {
	return &stack{first:nil}
}

func (S *stack) push(x int) {
	S.first = &listNode{value:x, next:S.first}
}

func (S stack) isEmpty() bool {
	return (S.first == nil)
}

func (S stack) top() int {
	return S.first.value
}

func (S *stack)pop() int {
	var t int = S.top()
	S.first = S.first.next
	return t
}

func (S stack) String() string {
	if S.isEmpty() {
		return "[]"
	} else {
		var s string = "[("
		s += strconv.Itoa(S.top())
		s += ")"
		var n *listNode = S.first.next
		for (n != nil) {
			s += " " + strconv.Itoa(n.value)
			n = n.next
		}
		
		s += "]"
		return s
	}
}

func main() {
	var S *stack = makeStack()
	fmt.Println(*S)
	fmt.Println("Pushing 1 ...")
	S.push(1)
	fmt.Println(S)
	fmt.Println("Pushing 2 ...")
	S.push(2)
	fmt.Println(S)
	fmt.Println("Pushing 3 ...")
	S.push(3)
	fmt.Println(S)
	fmt.Println("Pushing 4 ...")
	S.push(4)
	fmt.Println(S)
	fmt.Println("Popping off", S.pop(),"...")
	fmt.Println(S)
	fmt.Println("Popping off", S.pop(),"...")
	fmt.Println(S)
	fmt.Println("Pushing 5 ...")
	S.push(5)
	fmt.Println(S)
	fmt.Println("Pushing 6 ...")
	S.push(6)
	fmt.Println(S)
	fmt.Println("Popping off", S.pop(),"...")
	fmt.Println(S)
	fmt.Println("Popping off", S.pop(),"...")
	fmt.Println(S)
	fmt.Println("Popping off", S.pop(),"...")
	fmt.Println(S)
}	
			
	

