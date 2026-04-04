# a02

[![Go Reference](https://pkg.go.dev/badge/github.com/vbsw/go-experimental/a02.svg)](https://pkg.go.dev/github.com/vbsw/go-experimental/a02) [![Go Report Card](https://goreportcard.com/badge/github.com/vbsw/go-experimental/a02)](https://goreportcard.com/report/github.com/vbsw/go-experimental/a02)

## About
Package a02 tries to combine Packages a02a and a02b, but fails not compile. This is expected behaviour. Package a02 is published on <https://github.com/vbsw/go-experimental/a02>.

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

## Expected error is

	./main.go:44:13: cannot use printer (variable of type a02b.Printer) as a02a.Printer value in argument to a02a.Print: a02b.Printer does not implement a02a.Printer (wrong type for method Print)
		have Print(a02b.Driver)
		want Print(a02a.Driver)
