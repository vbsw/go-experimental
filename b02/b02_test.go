// This is free and unencumbered software released into the public domain.
//
// Anyone is free to copy, modify, publish, use, compile, sell, or
// distribute this software, either in source code form or as a compiled
// binary, for any purpose, commercial or non-commercial, and by any
// means.
//
// In jurisdictions that recognize copyright laws, the author or authors
// of this software dedicate any and all copyright interest in the
// software to the public domain. We make this dedication for the benefit
// of the public at large and to the detriment of our heirs and
// successors. We intend this dedication to be an overt act of
// relinquishment in perpetuity of all present and future rights to this
// software under copyright law.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
// OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.
//
// For more information, please refer to <http://unlicense.org>

package b02

import (
	"flag"
	"github.com/vbsw/go-experimental/b02a"
	"testing"
)

var size int

func init() {
	// example:
	// go test -bench=. -args -size=100
	flag.IntVar(&size, "size", 10, "array size")
}

// BenchmarkCallGo does benchmark on calling a Go function.
func BenchmarkCallGo(b *testing.B) {
	data := make([]uint, size, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b02a.CallGo(data)
	}
	if data[0] <= 0 {
		b.Fatal("wrong result")
	}
	b.StopTimer()
}

// BenchmarkCallGo does benchmark on calling a Go function calling a C function.
func BenchmarkCallGoIntoC(b *testing.B) {
	data := allocCInt(size)
	defer freeCMemory(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b02a.CallGoIntoC(data, size)
	}
	if data == nil {
		b.Fatal("wrong result")
	}
	b.StopTimer()
}

// BenchmarkCallGoIntoCIntoGo does benchmark on calling a Go function calling a C function calling a Go function.
func BenchmarkCallGoIntoCIntoGo(b *testing.B) {
	data := allocCInt(size)
	defer freeCMemory(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b02a.CallGoIntoCIntoGo(data, size)
	}
	if data == nil {
		b.Fatal("wrong result")
	}
	b.StopTimer()
}

/*
// BenchmarkCallGoIntoCNoCallback does benchmark on calling a Go function calling a C function.
// (cgo directive nocallback must have been enabled on call_go_into_c_nc in b02a.)
func BenchmarkCallGoIntoCNoCallback(b *testing.B) {
	data := allocCInt(size)
	defer freeCMemory(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b02a.CallGoIntoCNoCallback(data, size)
	}
	if data == nil {
		b.Fatal("wrong result")
	}
	b.StopTimer()
}
*/
