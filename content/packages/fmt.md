+++
title = 'fmt'
date = 2024-07-23T22:59:42+02:00
+++

The `fmt` package (https://pkg.go.dev/fmt) is the abbreviation for *Format* and used to format input- and output-values. 

## Output

There are various possibilities to print output-values to the default output:

```go
fmt.Print("Hello ", "Print()\n") // Hello Print()
fmt.Println("Hello Println()") // Hello Println()

var s = "Printf()"
fmt.Printf("Hello , %s\n", s) // Hello Printf()
```

The number of inputs are not limited. The `fmt.Printf()` function is the most advanced and uses *verbs* for formating:

- `%v`: prints any value in its default format
- `%+v`: prints structs and adds it properties (aka field names)
- `%T`: prints its data-type
- `%p`: prints the address of a pointer
- `%s`: prints the string
- etc.

## Input

You can also store the values as string:

```go
s1 := fmt.Sprint("Hello ", "Print()\n") // Hello Print()
s2 := fmt.Sprintln("Hello Println()") // Hello Println()
s3 := fmt.Sprintf("%03d: Hello, %s\n", 7, "James Bond") // 007: Hello James Bond
```

To write values directly into files, webserver, etc. use the `io.Wirter` interface and the following functions:

```go
fmt.Fprint()
fmt.Fprintln()
fmt.Fprintf()
```