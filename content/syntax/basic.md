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

### UTF-8

Go uses *UTF-8* to encode text. This has the benefit, that there is no issue with special characters. A string consists of multiple Unicode-characters which requires between 1 and 4 bytes (!). A `string` data type is an alias for `[]byte` why we can use `range` to loop. Each character is from type `rune` which is an alias of `uint32`, meaning 4 bytes.

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

A struct is a composite type (aka "zusammengesetzter Typ"). 

```go
type address struct {
    steet string
    city string
}

// Composite Literals
a := address {
    street: "Heiligengeistplatz",
    city: "Klagenfurt",
}
```

A struct can also contain other structs (also known as *composition*):

```go 
type user struct {
    firstName string
    lastName string
    address
}

u := user {
    firstName: "Max",
    lastName: "Mustermann",
    // Shortcut
    address: address {"Villacherstraße", "Klagenfurt"}
}
```

A struct and its properties can be enriched with *tags*, which provides additional information and are mostly used for conversations.


```go
type address struct {
    steet string `json:"streetName"`
    city string `json:"cityName"`
}
```

## Functions

Functions in Go can have multiple input-values of the same data type and also multiple return values (an error is always the last parameter):

```go 
func add(a, b, c int) (int, error) {
    ...
}
```

Another special feature is a *variadische* function. Its characteristic is that it can have any input-parameters:

```go 
func add(a ...int) int {
    // Variable "a" is a slice -> []int
    ...
}
```

### First-class citizens

Functions are also *first-class citizens* meaning, that they behave like other data types. Certain functions can also be a type and assigned to variables: 

```go
func filterFunc func(string) bool // define signature as type

func myFilter(s []string, filter filterFunc) { // use type within another function
    ...
    if filter(str) { // call type which must match the defined signature
        ...
    }
}
```

A working example is given here: `code/book1/listening2-44`.

### Closures

With *closures* you can increase the flexibility by returning a type-value:

```go
func lengthFilter(length int) filterFunc { // return a type-value
    return func(s string) bool { // define anyonmous function which matching signature
        ...
    }
}

// Or

func lengthFilter(length int) func(string) bool {
    ...
}
```

A working example is given here: `code/book1/listening2-47`.

### Defer

A *defer* is used to execute code at the end of a function, e.g. to clean up resources. Also in case of a return or a panic this code gets executed:

```go
func main() {
    myFunction()
}

func myFunction() {
    defer fmt.Println(1)
    defer fmt.Println(2)
    fmt.Println(3)
}

// Out:
// 3
// 2
// 1
```

As shown in previous code, in case of multiple defers, the LIFO (Last-In-First-Out) principle is used, which maybe is not expected. When multiple defers a grouped together, the order is different:


```go
func main() {
    myFunction()
}

func myFunction() {
    defer func() { // anonymous function
        fmt.Println(1)
        fmt.Println(2)
    } 
    fmt.Println(3)
}

// Out:
// 3
// 1
// 2
```

## Object Orientation

Go is **no** object-oriented programming (OOP) language. Classes and objects do not exist, although other OOP approaches can be used:

* Implement against interfaces (and not implementations)
* Use composition (no inheritance)

In Go, objects are concrete types (and not classes). If an object contains multiple properties, we use structs.

```go
type rectangle struct {
    length int
    width int
}

func area (r rectangle) int {
    return r.length * r.width
}

func main() {
    r := rectangle{length: 10, width: 5}
    fmt.Println(area(r))
}
```

### Methods

A method behaves like a function, but can be used only in the context of an object.

```go
func (r rectangle) area () int {
    return r.length * r.width
}

func main() {
    r := rectangle{length: 10, width: 5}
    fmt.Println(r.area())
}
```

When calling a method in Go, it only copies the value(s). When you want to change it, you must use pointers:

```go
func (r *rectangle) setLength(l int) {
    r.length = l
}
```

A working example is given here: `code/book1/listening2-56`.

## Arrays

An array can be used to store mulitple items of a type. Its size cannot be changed and must be defined:

```go
ip := [4]byte
ip[0] = 127
ip[1] = 0
ip[2] = 0
ip[3] = 1

// or simply

ip := [4]byte{127, 0, 0, 1}
```

## Slices

Slices are more flexible than arrays. Their size is not fixed. A slice points to an array and its current size is stored as `len` (length). The `cap` (capacity) of a slice is the overall size of the array.

```go
func print(s []int) {
    fmt.Printf("%p - len: %d cap: %d %#v\n", s len(s), cap(s), s)
}

// Out:
// 0x0 - len: 0 cap: 0 []int(nil)
// 0x1400000e100 - len: 1 cap: 1 []int{1} // new array
// 0x1400000e110 - len: 2 cap: 2 []int{1, 2} // new array
// 0x1400001e140 - len: 3 cap: 4 []int{1, 2, 3} // new array
// 0x1400001e140 - len: 4 cap: 4 []int{1, 2, 3, 4}
// 0x1400001a280 - len: 5 cap: 8 []int{1, 2, 3, 4, 5} // new array
```

A working example is given here: `code/book1/listening2-65`. Other possibilities to define a slice are:

```go
// Composite Literal
b := []int{} // an empty literal create a valid pointer

// make()
c := make([]int, 2/*len*/, 8/*cap*/)
```

To loop over a slice we use `range` which returns the index and a value: 

```go
for i, v := range s {
    fmt.Printf("%02d: %s\n", i, v)
}
```

A *slice of slice* is used to select a range within a slice:

```go
ip[:] // [127 0 0 1]
ip[0/*incl*/:3/*excl*/] // [127 0 0]
ip[1:] // [0 0 1]
```

The underlying array remains the same. When we want to modify the slice and do not change the original one, we have to `copy()` the values (which creates a new array).

## Maps

A map is a hash table. 

```go
// define a map
var m map[int/*key*/]string/*value*/
m = make(map[int]string) // create an empty map

// or use a Composite Literal
m := map[int]string {
    1: "A"
    ...
}

// or use the make() function
m := make(map[int]string)
```

## If

```go
if x := check(); x > 5 {
    ...
} else if x > 0 {
    ...
} else {
    ...
}
```

## Switch

The declaration of variables (similar as for an if) is also possible in a switch statement. At the end of a `case` the switch gets automatically terminated. When using a `fallthrough` within a case we can enforce the validation of the cases too. 

```go
switch a {
    case 0:
        ...
        fallthrough
    case 1, 2, 3:
        ...
    default: 
        ...
}
```

A special case is the type-switch:

```go
switch i := x.(type) {
    case int: // type of i is int
        ...
    ...
}
```

## For

```go
for i := 0; i <= 5; i++ {
    ...
}

// or

var i int
for i <= 5 {
    ...
    i++
}

// or

for { // inifinity loop
    ...
}

// or 

for i, v := range s { // use for arrays, slices, maps, strings, channels
    ...
}
```

Use `break` to cancel the loop or `continue` to jump of the next iteration.

## Labels and Goto

To jump within the code to a specific `Label`, you can use `goto`. Mostly they are used to jump e.g. from an inner-loop to the next iteration of the outer-loop.

