+++
title = 'Basic'
date = 2024-07-24T21:51:51+02:00
+++

## Data types

bool
string
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64
byte, rune, float32, float64
complex64, complex128

## Variables

Naming of variables are follow the *Camel Case* style. Depending on their lifecycle, a name should be short or long. Abbreviations should be written in uppercase in one exception, when they are not public.

```go
var num int // def. w/o value
var name string = "Rob Pike" // def. with value
var count = 10 // def. w/o type and value

// shortcuts
cnt := 10 // single def.
a, b := 24, 12 // multiple def.
```

Variables can be also grouped for readability:

```go
var (
    home = os.Getenv("HOME")
    user = os.Getenv("USER")
    goPath = os.Getenv("GOPATH")
)
```

## Constants

The declaration of constants are very similar to *Variables*. The shortcut is not available.

For numerical constants a counter can be used:

```go
const (
    _ = iota // 0
    Mon // 1
    Tue // 2
    Wed // 3
    Thu // 4
    Fri // 5
    _ // 6
)
```

## Pointers

A pointer references a memory address in the RAM. A pointer does not contain any value.

```go
var a int
var b *int // define pointer to an int
a = 123
b = &a // get memory address of a and store it in b
fmt.Println(b, *b) // out: 0x40e020 123 (*b=de-reference)
*b = 100 // set b value (which has the same memory address as a))
fmt.Println(a) // out: 100 
```

## Custom Types

A custom type can be defined:

```go
type meter int
```

Such a type always uses a basis type, i.e. `int`. Through Go's compiler you cannot simply calculate with custom types. When necessary, you always have to cast its value, like:

```go
a := meter(10)
b := meter(5)

//cnt := a + b // won't work
cnt := int(a) + int(b) // type cast basis value
```

## Structs

TODO
