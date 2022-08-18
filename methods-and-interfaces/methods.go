package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
type Myfloat float64

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)

}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))
	f := Myfloat(-math.Sqrt2)
	fmt.Println((f.Abs()))
	v.Scale(10)
	fmt.Println(v.Abs())

	p := &Vertex{4, 3}
	p.Scale(3)
	Scale(p, 8)
	fmt.Println(v, p)

	fmt.Println(Abs(v))
	//fmt.Println(Abs(&v))

	fmt.Println(v.Abs())
	p = &v
	fmt.Println(p.Abs())

}
