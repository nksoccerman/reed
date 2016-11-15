package main

import (
	"fmt"
	"strconv"
)

type turtle struct {
	name string
	x int
	y int
	facing string
}

func makeTurtle(s string) *turtle {
	return &turtle{name: s, x: 0, y: 0, facing: "E"}
}

func (t *turtle) right() {
	if t.facing == "E" {
		t.facing = "S"
	} else if t.facing == "S" {
		t.facing = "W"
	} else if t.facing == "W" {
		t.facing = "N"
	} else {
		t.facing = "E"
	}
}

func (t *turtle) forward() {
	if t.facing == "E" {
		t.x += 1
	} else if t.facing == "S"{
		t.y -= 1
	} else if t.facing == "W" {
		t.x -= 1
	} else {
		t.y += 1
	}
}

func (t *turtle) left() {
        if t.facing == "E" {
                t.facing = "N"
        } else if t.facing == "S" {
                t.facing = "E"
        } else if t.facing == "W" {
                t.facing = "S"
        } else {
                t.facing = "W"
        }
}

func (t turtle) String() string {
	return "{turtle " + t.name + " at (" + strconv.Itoa(t.x) + "," + strconv.Itoa(t.y) + ") heading " + t.facing + "}" 
}

func main() {
  var p *turtle = makeTurtle("pat")
  fmt.Println(p)
  p.left()
  p.forward()
  p.forward()
  fmt.Println(p)
  p.right()
  p.forward()
  fmt.Println(p)
}
