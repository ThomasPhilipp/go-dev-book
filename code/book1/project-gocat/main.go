package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Args) // [0]: programm path

	if len(os.Args) == 1 {
		fmt.Println("A file as arugment is necessary")
		os.Exit(1)
	}

	for _, fname := range os.Args[1:] {
		f, err := os.Open(fname)
		if err != nil {
			log.Printf("Error by opening file: %s\n%s", fname, err)
			f.Close()
			continue
		}

		_, err = io.Copy(os.Stdout, f)

		if err != nil {
			log.Printf("Error by printing out: %s\n%s", fname, err)
		}
		f.Close()
	}

}
