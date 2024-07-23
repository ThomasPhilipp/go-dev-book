+++
title = 'Quick Start'
date = 2024-07-23T22:40:11+02:00
+++

To create a new Go project execute:

`
$ go mod init https://<sample-url>/<user>/<project-name>
`

> For a local project any arbitrary URL can be used. For a professional projects you must shall use the Git URL.

The previous command automatically creates the `go.mod` and a `main.go` file. The latter one also belongs to the `main` package which marks it  as an *executeable application*.

A package name (like *fmt*) also defines a *namespace* which includes functions, types, variables and constants.

To build the application execute:

`$ go build`

This creates an executeable application in our folder. To run the application directly, execute:

`$ go run main.go`

Here, the application also gets compiled into a temporary folder and directly started from there. When finished this file gets automatically deleted.