package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}
type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I
	i = &T{"Hello"}
	i.M()
	descrive(i)
	i = F(math.Pi)
	i.M()
	descrive(i)
}

func descrive(i I) {
	fmt.Printf("(%v,%T)\n", i, i)
}
