+++
title = 'Quick Start'
date = 2024-07-23T22:40:11+02:00
+++

To create a new Go project execute the following command:

```go
$ go mod init <url>/<project-name>
```

> For a local project any arbitrary URL can be used. For a professional projects you must shall use the Git URL.

The output of the command is a `go.mod` file, also known as a *Go module*. Each executeable application requires a file with a `main`-package and a `main()`-function. Typically, this file is named `main.go`. A project can contain multiple `main`-packages, e.g. a server and client. A best practice approach is to group them under the `cmd`-folder:

```go
project-x
    ...
    cmd
        client
            main.go
        server
            main.go
```

A package name (like *fmt*) also defines a *namespace* which includes functions, types, variables and constants.

To build the application execute:

`$ go build`

This creates an executeable application in our folder. To run the application directly, execute:

`$ go run cmd/client/main.go`

Here, the application also gets compiled into a temporary folder and directly started from there. When finished this file gets automatically deleted.


