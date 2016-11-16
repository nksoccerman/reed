package main

import (
	"fmt"
	"strconv"
)

type expn interface {
	eval (map[string] int) int
}

type stmt interface {
	exec (map[string] int)
}

type cond interface {
	chck (map[string] int) bool
}

//
// blck type along with its run method.
//

type blck struct {
	stmts []stmt
}

func (b *blck) run(va map[string] int) {
	var s stmt
	for _,s = range b.stmts {
		s.exec(va)
	}
}

//
// cond types along with their chck method.
//

type lessCond struct {
	lhs expn
	rhs expn
}

func (c *lessCond) chck(va map[string] int) bool {
	return c.lhs.eval(va) < c.rhs.eval(va)
}

type equalCond struct {
	lhs expn
	rhs expn
}

func (c *equalCond) chck(va map[string] int) bool {
	return c.lhs.eval(va) == c.rhs.eval(va)
}

type orCond struct {
	lhs cond
	rhs cond 
}

func (c *orCond) chck(va map[string] int) bool {
	return (c.lhs.chck(va) || c.rhs.chck(va))
}

type andCond struct {
	lhs cond
	rhs cond
}

func (c *andCond) chck(va map[string] int) bool {
	return (c.lhs.chck(va) && c.rhs.chck(va))
}

type notCond struct {
	c cond
}

func (c *notCond) chck(va map[string] int) bool {
	return !(c.c.chck(va))
}

//
// expn types along with their eval method
//

type plusExpn struct {
	lhs expn
	rhs expn
}

func (e *plusExpn) eval(va map[string]int) int {
	return (e.lhs.eval(va) + e.rhs.eval(va))
}

type timesExpn struct {
	lhs expn
	rhs expn
}

func (e *timesExpn) eval(va map[string] int) int {
	return (e.lhs.eval(va) * e.rhs.eval(va))
}

type numberExpn struct {
	val int
}

func (e *numberExpn) eval(va map[string] int) int {
	return e.val
}

type lookupExpn struct {
	name string
}

func (e *lookupExpn) eval(va map[string] int) int {
	return va[e.name]
}

type inputExpn struct {
}

func (e *inputExpn) eval(va map[string] int) int {
	fmt.Print("? ")
	var i int
	fmt.Scanf("%d", &i)
	return i
}


//
// stmt types along with their exec method.
//

type ifThenElseStmt struct {
	cnd cond
	thn *blck
	els *blck
}

func (s ifThenElseStmt) exec(va map[string] int) {
	if s.cnd.chck(va) {
		s.thn.run(va)
	} else {
		s.els.run(va)
	}		
}

type whileStmt struct {
	cnd cond
	bdy *blck
}

func (s *whileStmt) exec(va map[string] int) {
	for s.cnd.chck(va) {
		s.bdy.run(va)
	}
}

type assignStmt struct {
	lhs string
	rhs expn
}

func (s *assignStmt) exec(va map[string] int) {
	var v int = s.rhs.eval(va)
	va[s.lhs] = v
}

type printStmt struct {
	arg expn
}

func (s *printStmt) exec(va map[string] int) {
	var v int = s.arg.eval(va)
	fmt.Println(strconv.Itoa(v))
}

//
// main
//

func main() {
	var pgm1 *blck = &blck{stmts:[]stmt{
		&assignStmt{lhs:"number",rhs:&inputExpn{}},
		&assignStmt{lhs:"product",rhs:&numberExpn{val:1}},
		&assignStmt{lhs:"count",rhs:&numberExpn{val:1}},
		&whileStmt{
			cnd:&orCond{
				lhs:&lessCond{
					lhs:&lookupExpn{name:"count"},
					rhs:&lookupExpn{name:"number"}},
				rhs:&equalCond{
					lhs:&lookupExpn{name:"count"},
					rhs:&lookupExpn{name:"number"}}},
			bdy:&blck{stmts:[]stmt{
				&assignStmt{
					lhs:"product",
					rhs:&timesExpn{
						lhs:&lookupExpn{name:"product"},
						rhs:&lookupExpn{name:"count"}}},
				&assignStmt{
					lhs:"count",
					rhs:&plusExpn{
						lhs:&lookupExpn{name:"count"},
						rhs:&numberExpn{val:1}}}}}},
		&printStmt{arg:&lookupExpn{name:"product"}}}}
	pgm1.run(make(map[string] int))

	var pgm2 *blck = &blck{stmts:[]stmt{
		&assignStmt{lhs:"number", rhs:&inputExpn{}},
		&whileStmt{
			cnd:&equalCond{
				lhs:&lookupExpn{name:"number"},
				rhs:&numberExpn{val:10}},
			bdy:&blck{stmts:[]stmt{
				&assignStmt{
					lhs:"number",
					rhs:&plusExpn{
						lhs:&lookupExpn{name:"number"},
						rhs:&numberExpn{val:1}}},
				&printStmt{arg:&lookupExpn{name:"number"}}}}}}}
	pgm2.run(make(map[string] int))

	var pgm3 *blck = &blck{stmts:[]stmt{
	&assignStmt{lhs:"number", rhs:&inputExpn{}},
	&ifThenElseStmt{
		cnd:&equalCond{
			lhs:&lookupExpn{name:"number"},
			rhs:&numberExpn{val:10}},
		thn:&blck{stmts:[]stmt{
			&printStmt{arg:&lookupExpn{name:"number"}}}},
		els:&blck{stmts:[]stmt{}}}}}
	pgm3.run(make(map[string] int))
			
}
