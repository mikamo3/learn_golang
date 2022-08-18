package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {

	vp := &Vertex{5, 10}
	fmt.Println("Before scalling:%v, Abs: %v\n", vp, vp.Abs())
	vp.Scale(10)
	fmt.Println("After scalling:%v, Abs: %v\n", vp, vp.Abs())
}
