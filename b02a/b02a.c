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

#include <stddef.h>

/* Go functions can not be passed to c directly.            */
/* They can only be called from c.                          */
/* This code is an indirection to call Go callbacks.        */
/* _cgo_export.h is generated automatically by cgo.         */
#include "_cgo_export.h"

void call_go_into_c(void *const arr, size_t length) {
	size_t i; unsigned int *const arr_int = (unsigned int *)arr;
	arr_int[0]++;
	for (i = 1; i < length; i++) {
		if (arr_int[i-1]-1 > arr_int[i]) {
			arr_int[i]++;
		}
	}
}

void call_go_into_c_nc(void *const arr, size_t length) {
	size_t i; unsigned int *const arr_int = (unsigned int *)arr;
	arr_int[0]++;
	for (i = 1; i < length; i++) {
		if (arr_int[i-1]-1 > arr_int[i]) {
			arr_int[i]++;
		}
	}
}

void call_go_into_c_into_go(void *const arr, size_t length) {
	unsigned int *const arr_int = (unsigned int *)arr;
	cIntoGo(arr_int, length);
}
