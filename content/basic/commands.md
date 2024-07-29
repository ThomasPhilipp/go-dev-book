+++
title = 'Commands'
date = 2024-07-29T20:15:01+02:00
+++

## go build

This command creates an executeable file. The name of the file is the same as the name of the folder, containing the `main`-package. To rename the file execute `$ go build -o <newName>`. 

### Cross compiling

Per default the previous commands always compiles the Go application for the own operating system. If you want to build it for other operating systems, you have to set the `GOOS` and `GOARCH` environment variables:

```go
env GOOS=<window> GOARCH=<amd64> go build
```

* `GOOS`: defines the target operating system
* `GOARCH`: defines the target architecture

### Build tags

Build tags are used when we only want to compile certain parts of a code for the desired operating system, e.g. Windows. Therefore, we create e.g. a file called `network_win.go` with the following build tag:

```go
// +build windows

package myPackage

...
```

For all other OS (MacOS and Linux) we create a file called `network_other.go` with: 

```go
// +build darwin linux
```

You can also define your own build tags and are not limited to environment variables only:

```go
// +build debug

package myPackage

const DBEUG = true

...
```

The custom build tag must be considered by compiling the code: `$ go build -tags debug`. Build tags can be also used for testing, e.g. `$ go test -tags <tagName>`. 

## go install

This commands compiles the Go application and installs it into the `bin`-folder of the `GOPATH`-variable. 