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

// Package b02a is a support package for package b02 to benchmark calls from Go into C and from C into Go.
package b02a

// Cgo directives are available starting from Go 1.23
// #cgo nocallback call_go_into_c_nc

// #include <stddef.h>
// void call_go_into_c(void*,size_t);
// void call_go_into_c_nc(void*,size_t);
// void call_go_into_c_into_go(void*,size_t);
import "C"
import "unsafe"

// CallGo is for comparison with other functions.
func CallGo(arr []uint) {
	arr[0]++
	for i := 1; i < len(arr); i++ {
		if arr[i-1]-1 > arr[i] {
			arr[i]++
		}
	}
}

// CallGoIntoC calls into C.
func CallGoIntoC(arr unsafe.Pointer, length int) {
	C.call_go_into_c(arr, C.size_t(length))
}

// CallGoIntoCNoCallback calls into C, but with a C function that should have the Cgo directive nocallback enabled.
// (This package does not have nocallback enabled, because the directive is available starting from Go 1.23.)
func CallGoIntoCNoCallback(arr unsafe.Pointer, length int) {
	C.call_go_into_c_nc(arr, C.size_t(length))
}

// CallGoIntoCIntoGo calls into C, then calls into Go.
func CallGoIntoCIntoGo(arr unsafe.Pointer, length int) {
	C.call_go_into_c_into_go(arr, C.size_t(length))
}

//export cIntoGo
func cIntoGo(arr *C.uint, length C.size_t) {
	arrGo := unsafe.Slice(arr, length)
	arrGo[0]++
	for i := 1; i < len(arrGo); i++ {
		if arrGo[i-1]-1 > arrGo[i] {
			arrGo[i]++
		}
	}
}
