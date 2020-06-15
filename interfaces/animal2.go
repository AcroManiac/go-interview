package main

import "fmt"

type animal2 interface {
	walker
	runner
}

type bird interface {
	walker
	flier
}

type walker interface {
	walk()
}

type runner interface {
	run()
}

type flier interface {
	fly()
}

type cat2 struct{}
type eagle struct{}

func (c *cat2) walk() {
	fmt.Println("cat is walking")
}

func (c *cat2) run() {
	fmt.Println("cat is running")
}

func (e *eagle) walk() {
	fmt.Println("eagle is walking")
}

func (e *eagle) fly() {
	fmt.Println("eagle is flying")
}

func walk(w walker) {
	w.walk()
}

func animal2Test() {
	var c animal2 = &cat2{}
	var e bird = &eagle{}
	c.walk()
	c.run()
	walk(c)
	walk(e)
}
