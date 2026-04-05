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

// Package b01 does benchmark on sorting.
package b01

type iOrdered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

// BinarySearchCustom searches an element in list and returns its index. Second return value
// is true, if element is in list, otherwise false and the index returned is the
// insert index. Elements in list must be unique and in ascending order.
func BinarySearchCustom(list []int, element int) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}

// BinarySearchGEN is the same as BinarySearchCustom, but generic.
func BinarySearchGEN[T iOrdered](list []T, element T) (int, bool) {
	left := 0
	right := len(list) - 1
	for left <= right {
		middle := (left + right) / 2
		value := list[middle]
		if value < element {
			left = middle + 1
		} else if value > element {
			right = middle - 1
		} else {
			return middle, true
		}
	}
	return left, false
}
