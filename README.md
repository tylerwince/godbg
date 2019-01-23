# godbg 🐛 [![Build Status](https://travis-ci.org/tylerwince/godbg.svg?branch=master)](https://travis-ci.org/tylerwince/godbg)

`godbg` is an implementation of the Rust2018 builtin debugging macro `dbg`.

The purpose of this package is to provide a better and more effective workflow for
people who are "print debuggers".

`go get github.com/tylerwince/godbg`

## The old way:

```go

package main

import "fmt"

func main() {
    a := 1
    fmt.Printf("My variable: a is equal to: %d", a)
}

```
outputs:

```
My variable: a is equal to: 1
```

## The _new_ (and better) way

```go

package main

import . "github.com/tylerwince/godbg"

func main() {
    a := 1
    Dbg(a)
}

```
outputs:

```
[main.go:7] a = 1
```

### This project is a work in progress and all feedback is appreciated.

The next features that are planned are:

- [ ] Tests
- [ ] Fancy Mode (display information about the whole callstack)
- [ ] Performance Optimizations
- [ ] Typing information
