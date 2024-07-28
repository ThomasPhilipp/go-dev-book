package main

import "fmt"

func print(s []int) {
	fmt.Printf("%p - len: %d cap: %d %#v\n", s, len(s), cap(s), s)
}

func main() {
	var a []int
	print(a)
	a = append(a, 1)
	print(a)
	a = append(a, 2)
	print(a)
	a = append(a, 3)
	print(a)
	a = append(a, 4)
	print(a)
	a = append(a, 5)
	print(a)
}
