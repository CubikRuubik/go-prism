package main

import "fmt"

type Greeter struct {
	Name string
}

func NewGreeter(name string) *Greeter {
	return &Greeter{Name: name}
}

func (g *Greeter) Greet() {
	fmt.Printf("Hello, %s!\n", g.Name)
}
