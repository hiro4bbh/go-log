# Package golog provides a logging framework making your logging easier and happier.
[![Build Status](https://travis-ci.org/hiro4bbh/go-log.svg?branch=master)](https://travis-ci.org/hiro4bbh/go-log)
[![Report Status](https://goreportcard.com/badge/github.com/hiro4bbh/go-log)](https://goreportcard.com/report/github.com/hiro4bbh/go-log)

Copyright 2018- Tatsuhiro Aoshima (hiro4bbh@gmail.com).

# Introduction
Package golog provides a logging framework making your logging easier and happier.
Furthermore, golog provides the terminal emulator supporting full-width (zenkaku in Japanese) characters (other implementations support them partially or incorrectly).

See golog's document on [GoDoc](https://godoc.org/github.com/hiro4bbh/go-log).
You can test the examples:

```
GOLOG_MINLEVEL=debug go run example/main.go
GOLOG_MINLEVEL=debug go run example/main.go | cat
go run example/term.go
```
