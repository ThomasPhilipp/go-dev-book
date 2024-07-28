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

// First class citizens example
func main() {
	// Sample filter defined anonymous
	f := func(s string) bool {
		if len(s) < 3 {
			return false
		}
		return true
	}

	// Test strings
	s := []string{"a", "abcd", "abc", "ab"}
	fmt.Println("Sample string: ", s)
	// Verify output
	fmt.Println("Filtered string: ", myFilter(s, f))
}
