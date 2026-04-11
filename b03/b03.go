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

// Package b03 benchmarks simple wildcard string matching implementations.
package b03

// void is_match(const void*,int,const void*,int,int*);
import "C"

type tState int

const (
	none tState = iota
	skipping
	skippingEscape
	escape
)

func isMatchByteByte(pattern, str []byte) bool {
	if len(str) > 0 {
		if len(pattern) > 0 {
			i, j, state := 0, 0, none
			for i < len(pattern) && j < len(str) {
				p, s := pattern[i], str[j]
				switch state {
				case none:
					if p == '*' {
						i, state = i+1, skipping
					} else if p == '?' {
						i, j = i+1, j+1
					} else if p == '\\' {
						if i+1 == len(pattern) {
							return p == s && j+1 == len(str)
						} else {
							i, state = i+1, escape
						}
					} else if p != s {
						return false
					} else {
						i, j = i+1, j+1
					}
				case skipping:
					if p == '*' {
						i++
					} else if p == '\\' {
						if i+1 == len(pattern) {
							return p == s && j+1 == len(str)
						} else {
							i, state = i+1, skippingEscape
						}
					} else if p == '?' {
						i, j = i+1, j+1
					} else if p != s {
						j++
					} else {
						i, j, state = i+1, j+1, none
					}
				case skippingEscape:
					if p == s {
						i, j, state = i+1, j+1, none
					} else {
						j++
					}
				case escape:
					if p == s {
						i, j, state = i+1, j+1, none
					} else {
						return false
					}
				}
			}
			if i == len(pattern) {
				if j == len(str) {
					return true
				} else {
					return state == skipping
				}
			} else {
				if j == len(str) && pattern[i-1] != '\\' {
					for ; i < len(pattern); i++ {
						if pattern[i] != '*' {
							return false
						}
					}
				} else {
					return false
				}
			}
		} else {
			return false
		}
	} else if len(pattern) > 0 {
		for _, b := range pattern {
			if b != '*' {
				return false
			}
		}
	}
	return true
}

func isMatchByteStr(pattern []byte, str string) bool {
	if len(str) > 0 {
		if len(pattern) > 0 {
			i, j, state := 0, 0, none
			for i < len(pattern) && j < len(str) {
				p, s := pattern[i], str[j]
				switch state {
				case none:
					if p == '*' {
						i, state = i+1, skipping
					} else if p == '?' {
						i, j = i+1, j+1
					} else if p == '\\' {
						if i+1 == len(pattern) {
							return p == s && j+1 == len(str)
						} else {
							i, state = i+1, escape
						}
					} else if p != s {
						return false
					} else {
						i, j = i+1, j+1
					}
				case skipping:
					if p == '*' {
						i++
					} else if p == '\\' {
						if i+1 == len(pattern) {
							return p == s && j+1 == len(str)
						} else {
							i, state = i+1, skippingEscape
						}
					} else if p == '?' {
						i, j = i+1, j+1
					} else if p != s {
						j++
					} else {
						i, j, state = i+1, j+1, none
					}
				case skippingEscape:
					if p == s {
						i, j, state = i+1, j+1, none
					} else {
						j++
					}
				case escape:
					if p == s {
						i, j, state = i+1, j+1, none
					} else {
						return false
					}
				}
			}
			if i == len(pattern) {
				if j == len(str) {
					return true
				} else {
					return state == skipping
				}
			} else {
				if j == len(str) && pattern[i-1] != '\\' {
					for ; i < len(pattern); i++ {
						if pattern[i] != '*' {
							return false
						}
					}
				} else {
					return false
				}
			}
		} else {
			return false
		}
	} else if len(pattern) > 0 {
		for _, b := range pattern {
			if b != '*' {
				return false
			}
		}
	}
	return true
}

func isMatchStrStr(pattern, str string) bool {
	if len(str) > 0 {
		if len(pattern) > 0 {
			i, j, state := 0, 0, none
			for i < len(pattern) && j < len(str) {
				p, s := pattern[i], str[j]
				switch state {
				case none:
					if p == '*' {
						i, state = i+1, skipping
					} else if p == '?' {
						i, j = i+1, j+1
					} else if p == '\\' {
						if i+1 == len(pattern) {
							return p == s && j+1 == len(str)
						} else {
							i, state = i+1, escape
						}
					} else if p != s {
						return false
					} else {
						i, j = i+1, j+1
					}
				case skipping:
					if p == '*' {
						i++
					} else if p == '\\' {
						if i+1 == len(pattern) {
							return p == s && j+1 == len(str)
						} else {
							i, state = i+1, skippingEscape
						}
					} else if p == '?' {
						i, j = i+1, j+1
					} else if p != s {
						j++
					} else {
						i, j, state = i+1, j+1, none
					}
				case skippingEscape:
					if p == s {
						i, j, state = i+1, j+1, none
					} else {
						j++
					}
				case escape:
					if p == s {
						i, j, state = i+1, j+1, none
					} else {
						return false
					}
				}
			}
			if i == len(pattern) {
				if j == len(str) {
					return true
				} else {
					return state == skipping
				}
			} else {
				if j == len(str) && pattern[i-1] != '\\' {
					for ; i < len(pattern); i++ {
						if pattern[i] != '*' {
							return false
						}
					}
				} else {
					return false
				}
			}
		} else {
			return false
		}
	} else if len(pattern) > 0 {
		for _, b := range pattern {
			if b != '*' {
				return false
			}
		}
	}
	return true
}
