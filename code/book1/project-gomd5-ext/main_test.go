package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintMD5(t *testing.T) {
	in := strings.NewReader("go") // create an io.Reader with specified string
	buf := &bytes.Buffer{}        // create an io.Writer
	want := "34d1f91fb2e514b8576fab1a75a89a6b\n"

	printMD5(in, buf) // test function in main.go

	got := buf.String() // read out value from buffer
	if got != want {
		t.Errorf("printMD5()=%s\nwant:%s\n", got, want)
	}
}
