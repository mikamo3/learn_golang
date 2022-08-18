package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 43}
	z := Person{"Zephod Beeblebrox", 9002}
	fmt.Println(a, z)
}
