+++
title = 'os'
date = 2024-07-28T21:07:27+02:00
+++

## os.Args

Used to parse the arguments of a command line.

```go
os.Args[0] // program path
os.Args[1:] // program argument(s) 
```

## os.Getenv

Used to read environment variables, e.g. `os.Getenv("HOME")`.