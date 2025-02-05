package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {
	h := md5.New()
	_, err := io.Copy(h, os.Stdin)
	if err != nil {
		fmt.Println("Error by reading:", err)
	}
	fmt.Printf("%x\n", h.Sum(nil)) // print hex value
}
