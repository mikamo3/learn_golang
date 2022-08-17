package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	s = s[1:]
	printSlice(s)
	s = s[1:1]
	printSlice(s)
	s = s[:2]
	printSlice(s)
	s = s[1:]
	printSlice(s)
	s = s[:3]
	printSlice(s)
	var ss []int
	fmt.Println(ss, len(ss), cap(ss))
	if ss == nil {
		fmt.Println("nil")
	}
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)
}
