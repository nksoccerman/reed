package main

import (
	"fmt"
	"strconv"
)

type bstnode struct {
	key int
	left *bstnode
	right *bstnode
}

type bstree struct {
	root *bstnode
}

func (t *bstree) contains(x int) bool {
	var n *bstnode = t.root
	for n != nil {
		if x == n.key {
			return true
		} else if x < n.key {
			n = n.left
		} else {
			n = n.right
		}
	}
	return false
}

func (t *bstree) insert(x int) {

	var p *bstnode = nil
	var n *bstnode = t.root
	for n != nil {
		p = n
		if n.key == x { 
			return
		} else if x < n.key {
			n = n.left
		} else {
			n = n.right
		}
	}

	var b *bstnode = &bstnode{key: x, left: nil, right: nil}
	if p == nil {
		t.root = b
	} else if x < p.key {
		p.left = b
	} else {
		p.right = b
	}
}


func makeBstree(xs []int) *bstree {
	var t *bstree = &bstree{root:nil}
	for _,x := range xs {
		t.insert(x)
	}
	return t
}

func(n *bstnode) String() string {
	if n == nil {
		return "Lf"
	} else {
		var ks string = strconv.Itoa(n.key)
		var ls string = n.left.String()
		var rs string = n.right.String()
		return "Br"+"("+ks+","+ls+","+rs+")"
	}
}	
	
func(t *bstree) String() string {
	return t.root.String()
}

func flatten(n *bstnode) []int {
	dat := []int{}
	if n != nil {
		if (n.right != nil && n.left != nil) {
			l := flatten(n.left) 
			r := flatten(n.right)
			dat = append(l, n.key)
			dat = append(dat, r...)
			return dat
		} else if (n.right == nil && n.left != nil) {
			l := flatten(n.left)
			dat = append(l, n.key)
			return dat
				
		} else if (n.right != nil && n.left == nil) {
			r := flatten(n.right)
			dat = append(dat, n.key)
			dat = append(dat, r...)
			return dat
		} else {
			dat = []int{n.key}
			return dat
		}
	} else {
		return dat
	}
}

func(t *bstree) keysOf() []int {
        dat := flatten(t.root)
	return dat
}

func main() {
	var xs []int = []int{5,3,1,2,6,9,8}
	var t *bstree = makeBstree(xs)
	fmt.Println(t)
	var i int = 0
	for i < 10 {
		if t.contains(i) {
			fmt.Printf("The tree above contains %d.\n",i)
		} else {
			fmt.Printf("The tree above does not contain %d.\n",i)
		}
		i++
	}
	fmt.Println(t.keysOf())
}



