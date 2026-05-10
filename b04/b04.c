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
#include <stdbool.h>
#include <stdint.h>

#define NEW_LINE        0
#define NEW_LINE_PREFIX 1
#define INLINE_CHILD    2
#define INLINE_SIBLING  3

typedef struct {
	size_t keyBegin, keyEnd, keyLen;
	size_t valBegin, valEnd, valLen;
	size_t lineBegin, lineEnd, lineLen;
	size_t indent, lineNumber;
	size_t nextKeyBegin, nextLineBegin;
	uint8_t state;
	bool ignoreOpenEnd;
} parser_t;

static size_t iSkipWhitespaceAndChar(const char *const bytes, const size_t from, const size_t to, const char charToSkip) {
	for (size_t i = from; i < to; i++) {
		if ((bytes[i] < 0 || bytes[i] > 32) && bytes[i] != charToSkip)
			return i;
	}
	return to;
}

static bool isComment(const char *const bytes, const size_t from, const size_t to) {
	for (size_t i = from; i < to; i++) {
		if (bytes[i] < 0 || bytes[i] > 32) { // non whitespace
			if (bytes[i] == '#')
				return true;
			return false;
		} // else: whitespace
	}
	return false;
}


static size_t iSkipWhitespace(const char *const bytes, const size_t from, const size_t to) {
	for (size_t i = from; i < to; i++) {
		if (bytes[i] < 0 || bytes[i] > 32)
			return i;
	}
	return to;
}

static size_t iSkipWhitespaceReverse(const char *const bytes, const size_t from, const size_t to) {
	for (size_t i = to-1; i > from; i--) {
		if (bytes[i] < 0 || bytes[i] > 32)
			return i+1;
	}
	if (bytes[from] < 0 || bytes[from] > 32)
		return from+1;
	return from;
}

static bool parseLineBounds(parser_t *const p, const char *const bytes, const size_t size) {
	for (size_t i = p->nextLineBegin; i < size; i++) {
		if (bytes[i] == '\r') {
			if (i+1 < size) {
				p->lineNumber = p->lineNumber+1, p->lineBegin = p->nextLineBegin, p->lineEnd = i;
				if (bytes[i+1] == '\n')
					p->nextLineBegin = i+2;
				else
					p->nextLineBegin = i+1;
				return true;
			} else if (p->ignoreOpenEnd) {
				p->lineNumber = p->lineNumber+1, p->lineBegin = p->nextLineBegin, p->lineEnd = i, p->nextLineBegin = i+1;
				return true;
			}
			return false;
		} else if (bytes[i] == '\n') {
			p->lineNumber = p->lineNumber+1, p->lineBegin = p->nextLineBegin, p->lineEnd = i, p->nextLineBegin = i+1;
			return true;
		}
	}
	if (p->nextLineBegin < size && p->ignoreOpenEnd) {
		p->lineNumber = p->lineNumber+1, p->lineBegin = p->nextLineBegin, p->lineEnd = size, p->nextLineBegin = size;
		return true;
	}
	return false;
}

static void parseIndentation(parser_t *const p, const char *const bytes) {
	p->keyBegin = p->lineBegin, p->indent = 0;
	while (p->keyBegin < p->lineEnd && bytes[p->keyBegin] == '\t')
		p->indent = p->indent+1, p->keyBegin = p->keyBegin+1;
}

static bool parseInlineChildPrefix(parser_t *const p, const char *const bytes) {
	if (p->keyBegin < p->lineEnd && bytes[p->keyBegin] == '\\') {
		if (p->keyBegin+1 < p->lineEnd) {
			if (bytes[p->keyBegin+1] != '\\' && bytes[p->keyBegin+1] != '#' && bytes[p->keyBegin+1] != '|') {
				p->keyBegin = p->keyBegin+2, p->indent = p->indent+1;
				return true;
			}
		}
	}
	return false;
}

static void parseKey(parser_t *const p, const char *const bytes) {
	bool escape = false;
	for (size_t i = p->keyBegin; i < p->lineEnd; i++) {
		if (bytes[i] < 0 || bytes[i] > 32) { // non whitespace
			if (bytes[i] == '\\') {
				escape = !escape;
			} else if (bytes[i] == '#') {
				if (escape) {
					escape = false;
				} else {
					p->keyEnd = i, p->state = NEW_LINE;
					return;
				}
			} else if (bytes[i] == '|') {
				if (escape) {
					escape = false;
				} else {
					p->keyEnd = i, p->state = INLINE_SIBLING;
					return;
				}
			} else if (escape) {
				p->keyEnd = i-1, p->state = INLINE_CHILD;
				return;
			}
		} else { // whitespace
			p->keyEnd = i;
			return;
		}
	}
	p->keyEnd = p->lineEnd;
}

static void parseValue(parser_t *const p, const char *const bytes) {
	bool escape = false;
	for (size_t i = p->valBegin; i < p->lineEnd; i++) {
		if (bytes[i] < 0 || bytes[i] > 32) { // non whitespace
			if (bytes[i] == '\\') {
				escape = !escape;
			} else if (bytes[i] == '#') {
				if (escape) {
					escape = false;
				} else {
					p->valEnd = i, p->state = NEW_LINE;
					return;
				}
			} else if (bytes[i] == '|') {
				if (escape) {
					escape = false;
				} else {
					p->valEnd = i, p->state = INLINE_SIBLING;
					return;
				}
			} else if (escape) {
				p->valEnd = i-1, p->state = INLINE_CHILD;
				return;
			}
		} else if (escape) {
			p->valEnd = i-1, p->state = INLINE_CHILD;
			return;
		}
	}
	p->valEnd = p->lineEnd;
}

static void parseKeyValue(parser_t *const p, const char *const bytes) {
	const uint8_t stateOld = p->state;
	parseKey(p, bytes);
	if (stateOld == p->state) {
		p->valBegin = iSkipWhitespace(bytes, p->keyEnd, p->lineEnd);
		parseValue(p, bytes);
		p->nextKeyBegin = p->valEnd+1;
		p->valEnd = iSkipWhitespaceReverse(bytes, p->valBegin, p->valEnd);
		if (stateOld == p->state)
			p->state = NEW_LINE;
	} else {
		p->valBegin = p->keyEnd;
		p->valEnd = p->keyEnd;
		p->nextKeyBegin = p->keyEnd+1;
	}
	p->keyEnd = iSkipWhitespaceReverse(bytes, p->keyBegin, p->keyEnd);
	p->keyLen = p->keyEnd-p->keyBegin, p->valLen = p->valEnd-p->valBegin;
}

static bool next(parser_t *const p, const char *const bytes, const size_t size) {
	while (true) {
		switch(p->state) {
			case NEW_LINE:
				if (parseLineBounds(p, bytes, size)) {
					p->lineLen = p->lineEnd-p->lineBegin;
					parseIndentation(p, bytes);
					p->state = NEW_LINE_PREFIX;
					break;
				}
				return false;
			case NEW_LINE_PREFIX:
				p->keyBegin = iSkipWhitespaceAndChar(bytes, p->keyBegin, p->lineEnd, '|');
				if (isComment(bytes, p->lineBegin, p->lineEnd)) {
					p->state = NEW_LINE;
					break;
				} else if (parseInlineChildPrefix(p, bytes)) {
					p->state = NEW_LINE_PREFIX;
					break;
				}
				parseKeyValue(p, bytes);
				return true;
			case INLINE_CHILD:
				p->keyBegin = iSkipWhitespace(bytes, p->nextKeyBegin, p->lineEnd);
				while (parseInlineChildPrefix(p, bytes))
					p->keyBegin = iSkipWhitespace(bytes, p->keyBegin, p->lineEnd);
				if (isComment(bytes, p->keyBegin, p->lineEnd)) {
					p->state = NEW_LINE;
					break;
				}
				p->indent++;
				parseKeyValue(p, bytes);
				return true;
			case INLINE_SIBLING:
				p->keyBegin = iSkipWhitespaceAndChar(bytes, p->valEnd, p->lineEnd, '|');
				if (isComment(bytes, p->keyBegin, p->lineEnd)) {
					p->state = NEW_LINE;
				}
				parseKeyValue(p, bytes);
				return true;
		}
	}
	return false;
}

void parse(const void *const bytes, int size, int *const counter) {
	parser_t parser = {0};
	parser.ignoreOpenEnd = true;
	while (next(&parser, (char*)bytes, (size_t)size)) {
		counter[0]++;
	}
}
