# b02

[![Go Reference](https://pkg.go.dev/badge/github.com/vbsw/go-experimental/b02.svg)](https://pkg.go.dev/github.com/vbsw/go-experimental/b02) [![Go Report Card](https://goreportcard.com/badge/github.com/vbsw/go-experimental/b02)](https://goreportcard.com/report/github.com/vbsw/go-experimental/b02)

## About
Package b02 benchmarks calls from Go into C and from C into Go. Package b02 is published on <https://github.com/vbsw/go-experimental/b02>.

## Copying
This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <http://unlicense.org>

## Compile
This package needs Cgo to compile and Cgo needs a C compiler.

**Linux**  
For Cgo install GCC, or configure another compiler like clang (see <https://stackoverflow.com/questions/44856124/can-i-change-default-compiler-used-by-cgo>).

**Windows**  
For Cgo install tdm-gcc (<https://jmeubank.github.io/tdm-gcc/>), or some other Go ABI compatible compiler like MinGW-w64.

## Execute
Default size is 10. (size decides the load per function call.)

	go test -bench=.

Custom size N must be greater than 0.

	go test -bench=. -args -size=N

## References
- https://go.dev/doc/install
- https://jmeubank.github.io/tdm-gcc/
- https://github.com/go101/go101/wiki/CGO-Environment-Setup
- https://git-scm.com/book/en/v2/Getting-Started-Installing-Git

