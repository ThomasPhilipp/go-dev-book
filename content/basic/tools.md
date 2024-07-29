+++
title = 'Tools'
date = 2024-07-29T20:15:01+02:00
+++

## go help

The command gives an overview about further Go tools. To get more details to a specific tool call `$ go help <fmt>`.

## go env

The commands prints out all configured environment variables. The most important variables are:
* `GOPATH`: requires read & write access and mostly points the `HOME`-folder.
* `GOROOT`: points to the Go installation which includes the compiler and the default library.

To print out more details about available environment variables execute `$ go help environment`. 

## gc (or gccgo)

The used Go compiler.

## gofmt

The tool is mostly deeply integrated into the IDE and automatically formats the code during saving your code changes. This process is executed by parsing the file and map it internally to a syntax tree before it gets formatted. 

The formatting process also can be started manually. To format a file, call `$ gofmt -w <fileName.go>`. If you want for format an folder, call `gofmt <path>`. 

## goimports

Also deeply integrated into IDEs is this tool, which automatically manages the imports. If not installed, call: `$ go get golang.org/x/tools/cmd/goimports`. 

## godoc

The tool is used to generate a documentation by using existing comments within a code. If not installed, call: `$ go get golang.org/x/tools/cmd/godoc`. To print out the documentation in a terminal call: `$ godoc <packageName>`. By adding the `-http=:<port>` option, you can read the documentation in a browser.

