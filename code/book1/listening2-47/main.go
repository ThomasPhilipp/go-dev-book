package main

import "fmt"

type filterFunc func(string) bool

func myFilter(s []string, filter filterFunc) []string {
	var out []string

	for _, str := range s {
		if filter(str) {
			out = append(out, str)
		}
	}

	return out
}

func lengthFilter(length int) filterFunc {
	return func(s string) bool {
		if len(s) < length {
			return false
		}
		return true
	}
}

// Closure example
func main() {
	f2 := lengthFilter(2)
	f3 := lengthFilter(3)

	// Test strings
	s := []string{"a", "abcd", "abc", "ab"}
	fmt.Println("Sample string: ", s)

	// Verify output
	fmt.Println("f2 output: ", myFilter(s, f2))
	fmt.Println("f3 output: ", myFilter(s, f3))
}
