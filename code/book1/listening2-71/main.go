package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(a[0:3])
	fmt.Println(a[:3])
	fmt.Println(a[3:5])
	fmt.Println(a[3:])
	fmt.Println(a[:])
}
