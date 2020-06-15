package main

import "fmt"

type greeter interface {
	greet(string) string
}

type russian struct{}
type american struct{}

func (r *russian) greet(name string) string {
	return fmt.Sprintf("Привет, %s", name)
}

func (a *american) greet(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func sayHello(g greeter, name string) {
	fmt.Println(g.greet(name))
}

func greet() {
	sayHello(&russian{}, "Петя")
	sayHello(&american{}, "Bill")
}
