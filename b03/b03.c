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

#define NONE 0
#define SKIPPING 1
#define SKIPPING_ESCAPE 2
#define ESCAPE 3

void is_match(const void *const ppattern, const int pattern_len, const void *const pstr, const int str_len, int *const ret) {
	const char *const pattern = (const char *const) ppattern;
	const char *const str = (const char *const) pstr;
	if (str_len > 0) {
		if (pattern_len > 0) {
			int i = 0, j = 0, state = NONE;
			while (i < pattern_len && j < str_len) {
				const char p = pattern[i], s = str[j];
				switch (state) {
				case NONE:
					if (p == '*') {
						i = i+1, state = SKIPPING;
					} else if (p == '?') {
						i = i+1, j = j+1;
					} else if (p == '\\') {
						if (i+1 == pattern_len) {
							*ret = (int)(p == s && j+1 == str_len); return;
						} else {
							i = i+1, state = ESCAPE;
						}
					} else if (p != s) {
						*ret = 0; return;
					} else {
						i = i+1, j = j+1;
					}
					break;
				case SKIPPING:
					if (p == '*') {
						i++;
					} else if (p == '\\') {
						if (i+1 == pattern_len) {
							*ret = (int)(p == s && j+1 == str_len); return;
						} else {
							i = i+1, state = SKIPPING_ESCAPE;
						}
					} else if (p == '?') {
						i = i+1, j = j+1;
					} else if (p != s) {
						j++;
					} else {
						i = i+1, j = j+1, state = NONE;
					}
					break;
				case SKIPPING_ESCAPE:
					if (p == s) {
						i = i+1, j = j+1, state = NONE;
					} else {
						j++;
					}
					break;
				case ESCAPE:
					if (p == s) {
						i = i+1, j = j+1, state = NONE;
					} else {
						*ret = 0; return;
					}
					break;
				}
			}
			if (i == pattern_len) {
				if (j == str_len) {
					*ret = 1; return;
				} else {
					*ret = (int)(state == SKIPPING); return;
				}
			} else {
				if (j == str_len && pattern[i-1] != '\\') {
					for (; i < pattern_len; i++) {
						if (pattern[i] != '*') {
							*ret = 0; return;
						}
					}
				} else {
					*ret = 0; return;
				}
			}
		} else {
			*ret = 0; return;
		}
	} else if (pattern_len > 0) {
		int i;
		for (i = 0; i < pattern_len; i++) {
			if (pattern[i] != '*') {
				*ret = 0; return;
			}
		}
	}
	*ret = 1;
}
