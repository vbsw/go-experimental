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

package b03

import (
	"path/filepath"
	"testing"
)

func TestIsMatchStrStr(t *testing.T) {
	if !isMatchStrStr("*", "asdf") {
		t.Error("failed pattern \"*\"")
	}
	if !isMatchStrStr("*?", "asdf") {
		t.Error("failed pattern \"*?\"")
	}
	if !isMatchStrStr("*?*", "asdf") {
		t.Error("failed pattern \"*?*\"")
	}
	if !isMatchStrStr("?*", "asdf") {
		t.Error("failed pattern \"?*\"")
	}
	if !isMatchStrStr("ab*ef*x", "abcdefghx") {
		t.Error("failed pattern \"?*\"")
	}
	if isMatchStrStr("*\\*hx", "ab*cdefghx") {
		t.Error("failed pattern \"*\\*hx\"")
	}
	if isMatchStrStr("*\\**", "abcdefghx") {
		t.Error("failed pattern \"*\\**\"")
	}
}

func TestIsMatchStrStrFilepath(t *testing.T) {
	if result, err := filepath.Match("*", "asdf"); err != nil || result != isMatchStrStr("*", "asdf") {
		t.Error("failed pattern \"*\"")
	}
	if result, err := filepath.Match("*?", ""); err != nil || result != isMatchStrStr("*?", "") {
		t.Error("failed pattern \"*?\"")
	}
	if result, err := filepath.Match("*?*", ""); err != nil || result != isMatchStrStr("*?*", "") {
		t.Error("failed pattern \"*?*\"")
	}
	if result, err := filepath.Match("?*", ""); err != nil || result != isMatchStrStr("?*", "") {
		t.Error("failed pattern \"?*\"")
	}
	if result, err := filepath.Match("?*", "a"); err != nil || result != isMatchStrStr("?*", "a") {
		t.Error("failed pattern \"?*\"")
	}
	if result, err := filepath.Match("ab*ef*", "abcdefgh"); err != nil || result != isMatchStrStr("ab*ef*", "abcdefgh") {
		t.Error("failed pattern \"ab*ef*\"")
	}
	if result, err := filepath.Match("ab*ef*x", "abcdefgh"); err != nil || result != isMatchStrStr("ab*ef*x", "abcdefgh") {
		t.Error("failed pattern \"ab*ef*x\"")
	}
	if result, err := filepath.Match("ab*ef*x", "abcdefghx"); err != nil || result != isMatchStrStr("ab*ef*x", "abcdefghx") {
		t.Error("failed pattern \"ab*ef*x\"")
	}
	if result, err := filepath.Match("?*ab", "aab"); err != nil || result != isMatchStrStr("?*ab", "aab") {
		t.Error("failed pattern \"?*ab\"")
	}
}

func BenchmarkIsMatchByteByte(b *testing.B) {
	result, str := true, []byte("abcdefghijklmnopqrstuvwxyz")
	pA, pB, pC := []byte("*"), []byte("*asdf"), []byte("asdf*")
	pD, pE := []byte("abcdefghijklmnopqr*"), []byte("*jklmnopqrstuvwxyz")
	pF, pG := []byte("*fghijklmnopqrst*"), []byte("*fghijklmnopqrst")
	pH := []byte("*efghijklm?opqrstuvwxyz")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = result && (isMatchByteByte(pA, str) == true)
		result = result && (isMatchByteByte(pB, str) == false)
		result = result && (isMatchByteByte(pC, str) == false)
		result = result && (isMatchByteByte(pD, str) == true)
		result = result && (isMatchByteByte(pE, str) == true)
		result = result && (isMatchByteByte(pF, str) == true)
		result = result && (isMatchByteByte(pG, str) == false)
		result = result && (isMatchByteByte(pH, str) == true)
	}
	b.StopTimer()
	if !result {
		b.Fatal("wrong result")
	}
}

func BenchmarkIsMatchByteStr(b *testing.B) {
	result, str := true, "abcdefghijklmnopqrstuvwxyz"
	pA, pB, pC := []byte("*"), []byte("*asdf"), []byte("asdf*")
	pD, pE := []byte("abcdefghijklmnopqr*"), []byte("*jklmnopqrstuvwxyz")
	pF, pG := []byte("*fghijklmnopqrst*"), []byte("*fghijklmnopqrst")
	pH := []byte("*efghijklm?opqrstuvwxyz")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = result && (isMatchByteStr(pA, str) == true)
		result = result && (isMatchByteStr(pB, str) == false)
		result = result && (isMatchByteStr(pC, str) == false)
		result = result && (isMatchByteStr(pD, str) == true)
		result = result && (isMatchByteStr(pE, str) == true)
		result = result && (isMatchByteStr(pF, str) == true)
		result = result && (isMatchByteStr(pG, str) == false)
		result = result && (isMatchByteStr(pH, str) == true)
	}
	b.StopTimer()
	if !result {
		b.Fatal("wrong result")
	}
}

func BenchmarkIsMatchStrStr(b *testing.B) {
	result, str := true, "abcdefghijklmnopqrstuvwxyz"
	pA, pB, pC := ("*"), ("*asdf"), ("asdf*")
	pD, pE := ("abcdefghijklmnopqr*"), ("*jklmnopqrstuvwxyz")
	pF, pG := ("*fghijklmnopqrst*"), ("*fghijklmnopqrst")
	pH := ("*efghijklm?opqrstuvwxyz")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = result && (isMatchStrStr(pA, str) == true)
		result = result && (isMatchStrStr(pB, str) == false)
		result = result && (isMatchStrStr(pC, str) == false)
		result = result && (isMatchStrStr(pD, str) == true)
		result = result && (isMatchStrStr(pE, str) == true)
		result = result && (isMatchStrStr(pF, str) == true)
		result = result && (isMatchStrStr(pG, str) == false)
		result = result && (isMatchStrStr(pH, str) == true)
	}
	b.StopTimer()
	if !result {
		b.Fatal("wrong result")
	}
}

func BenchmarkIsMatchClosures(b *testing.B) {
	result, str := true, "abcdefghijklmnopqrstuvwxyz"
	pA, pB, pC := ("*"), ("*asdf"), ("asdf*")
	pD, pE := ("abcdefghijklmnopqr*"), ("*jklmnopqrstuvwxyz")
	pF, pG := ("*fghijklmnopqrst*"), ("*fghijklmnopqrst")
	pH := ("*efghijklm?opqrstuvwxyz")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = result && (isMatchClosures(pA, str) == true)
		result = result && (isMatchClosures(pB, str) == false)
		result = result && (isMatchClosures(pC, str) == false)
		result = result && (isMatchClosures(pD, str) == true)
		result = result && (isMatchClosures(pE, str) == true)
		result = result && (isMatchClosures(pF, str) == true)
		result = result && (isMatchClosures(pG, str) == false)
		result = result && (isMatchClosures(pH, str) == true)
	}
	b.StopTimer()
	if !result {
		b.Fatal("wrong result")
	}
}

func BenchmarkIsMatchC(b *testing.B) {
	result, str := true, "abcdefghijklmnopqrstuvwxyz"
	pA, pB, pC := ("*"), ("*asdf"), ("asdf*")
	pD, pE := ("abcdefghijklmnopqr*"), ("*jklmnopqrstuvwxyz")
	pF, pG := ("*fghijklmnopqrst*"), ("*fghijklmnopqrst")
	pH := ("*efghijklm?opqrstuvwxyz")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = result && (isMatchC(pA, str) == true)
		result = result && (isMatchC(pB, str) == false)
		result = result && (isMatchC(pC, str) == false)
		result = result && (isMatchC(pD, str) == true)
		result = result && (isMatchC(pE, str) == true)
		result = result && (isMatchC(pF, str) == true)
		result = result && (isMatchC(pG, str) == false)
		result = result && (isMatchC(pH, str) == true)
	}
	b.StopTimer()
	if !result {
		b.Fatal("wrong result")
	}
}

func BenchmarkIsMatchCNoLoad(b *testing.B) {
	result, str := true, "abcdefghijklmnopqrstuvwxyz"
	pA, pB, pC := ("*"), ("*asdf"), ("asdf*")
	pD, pE := ("abcdefghijklmnopqr*"), ("*jklmnopqrstuvwxyz")
	pF, pG := ("*fghijklmnopqrst*"), ("*fghijklmnopqrst")
	pH := ("*efghijklm?opqrstuvwxyz")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = result && (isMatchCNoLoad(pA, str) == true)
		result = result && (isMatchCNoLoad(pB, str) == true)
		result = result && (isMatchCNoLoad(pC, str) == true)
		result = result && (isMatchCNoLoad(pD, str) == true)
		result = result && (isMatchCNoLoad(pE, str) == true)
		result = result && (isMatchCNoLoad(pF, str) == true)
		result = result && (isMatchCNoLoad(pG, str) == true)
		result = result && (isMatchCNoLoad(pH, str) == true)
	}
	b.StopTimer()
	if !result {
		b.Fatal("wrong result")
	}
}

func BenchmarkFilepathMatch(b *testing.B) {
	result, str := true, "abcdefghijklmnopqrstuvwxyz"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		match, err := filepath.Match("*", str)
		result = result && err == nil && match == true
		match, err = filepath.Match("*asdf", str)
		result = result && err == nil && match == false
		match, err = filepath.Match("asdf*", str)
		result = result && err == nil && match == false
		match, err = filepath.Match("abcdefghijklmnopqr*", str)
		result = result && err == nil && match == true
		match, err = filepath.Match("*jklmnopqrstuvwxyz", str)
		result = result && err == nil && match == true
		match, err = filepath.Match("*fghijklmnopqrst*", str)
		result = result && err == nil && match == true
		match, err = filepath.Match("*fghijklmnopqrst", str)
		result = result && err == nil && match == false
		match, err = filepath.Match("*efghijklm?opqrstuvwxyz", str)
		result = result && err == nil && match == true
	}
	b.StopTimer()
	if !result {
		b.Fatal("wrong result")
	}
}

func BenchmarkFilepathMatchNoErr(b *testing.B) {
	result, str := true, "abcdefghijklmnopqrstuvwxyz"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		match, _ := filepath.Match("*", str)
		result = result && match == true
		match, _ = filepath.Match("*asdf", str)
		result = result && match == false
		match, _ = filepath.Match("asdf*", str)
		result = result && match == false
		match, _ = filepath.Match("abcdefghijklmnopqr*", str)
		result = result && match == true
		match, _ = filepath.Match("*jklmnopqrstuvwxyz", str)
		result = result && match == true
		match, _ = filepath.Match("*fghijklmnopqrst*", str)
		result = result && match == true
		match, _ = filepath.Match("*fghijklmnopqrst", str)
		result = result && match == false
		match, _ = filepath.Match("*efghijklm?opqrstuvwxyz", str)
		result = result && match == true
	}
	b.StopTimer()
	if !result {
		b.Fatal("wrong result")
	}
}
