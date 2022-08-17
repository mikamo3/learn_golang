package main

import "fmt"

func fibonacci() func() int {
	x1 := -1
	x2 := -1
	return func() int {
		if x1 == -1 {
			x1 = 0
			return 0
		}
		if x2 == -1 {
			x2 = 1
			return 1
		}
		sum := x1 + x2
		x1 = x2
		x2 = sum
		return sum
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
