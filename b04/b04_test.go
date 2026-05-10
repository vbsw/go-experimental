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

package b04

import (
	gobytes "bytes"
	"github.com/vbsw/go-lib/tabformat"
	"os"
	"testing"
)

type tBenchmarkData struct {
	id, entry int
	buffer    [][]int
}

var benchmarkData *tBenchmarkData

func newBenchmarkData() *tBenchmarkData {
	benchmarkData := new(tBenchmarkData)
	benchmarkData.buffer = make([][]int, 148*3)
	for i := 0; i < 148*3; i += 3 {
		benchmarkData.buffer[i+0] = make([]int, 7)   // chunks
		benchmarkData.buffer[i+1] = make([]int, 7)   // deltas
		benchmarkData.buffer[i+2] = make([]int, 8*7) // bounds
	}
	return benchmarkData
}

func toInt(bytes []byte) int {
	var value int
	for _, b := range bytes {
		if b >= '0' && b <= '9' {
			value = value*10 + int(b-'0')
		} else {
			return -1
		}
	}
	return value
}

func (d *tBenchmarkData) readId(value, idName []byte) {
	index := gobytes.IndexByte(value, '=')
	if index == 2 && gobytes.Equal(value[:2], idName) {
		d.id = toInt(value[index+1:])
	} else {
		d.id = -1
	}
}

func (d *tBenchmarkData) readChunksDeltas(value []byte, offset int) {
	if d.id >= 0 {
		var i, n int
		for j := 0; j < len(value); j++ {
			if value[j] == ' ' {
				d.buffer[d.id*3+offset][n] = toInt(value[i:j])
				n, i = n+1, j+1
			}
		}
		d.buffer[d.id*3+offset][n] = toInt(value[i:])
	}
}

func (d *tBenchmarkData) readCollection(value []byte) {
	// nothing
}

func (d *tBenchmarkData) readEntry(value []byte) {
	if len(value) > 0 && value[0] >= '0' && value[0] <= '9' {
		d.entry = int(value[0] - '0')
	} else {
		d.entry = -1
	}
}

func (d *tBenchmarkData) readBounds(value []byte) {
	if d.id >= 0 && d.entry >= 0 {
		var i, n int
		for j := 0; j < len(value); j++ {
			if value[j] == ' ' {
				d.buffer[d.id*3+2][d.entry*8+n] = toInt(value[i:j])
				n, i = n+1, j+1
			}
		}
		d.buffer[d.id*3+2][d.entry*8+n] = toInt(value[i:])
	}
}

func TestReadFile(t *testing.T) {
	bytes, err := os.ReadFile("testdata.txt")
	if err != nil {
		t.Error("failed load file:", err.Error())
	} else if len(bytes) == 0 {
		t.Error("loaded 0 bytes")
	} else {
		var parser tabformat.ByteParser
		elementName := []byte("element")
		chunksName := []byte("chunks")
		deltasName := []byte("deltas")
		collectionName := []byte("collection")
		entryName := []byte("entry")
		boundsName := []byte("bounds")
		idName := []byte("id")
		benchmarkData = newBenchmarkData()
		parser.IgnoreOpenEnd = true
		for parser.Next(bytes) {
			switch parser.Indent {
			case 0:
				if gobytes.Equal(parser.Key(bytes), elementName) {
					benchmarkData.readId(parser.Value(bytes), idName)
				}
			case 1:
				key := parser.Key(bytes)
				if gobytes.Equal(key, chunksName) {
					benchmarkData.readChunksDeltas(parser.Value(bytes), 0)
				} else if gobytes.Equal(key, deltasName) {
					benchmarkData.readChunksDeltas(parser.Value(bytes), 1)
				} else if gobytes.Equal(key, collectionName) {
					benchmarkData.readCollection(parser.Value(bytes))
				}
			case 2:
				if gobytes.Equal(parser.Key(bytes), entryName) {
					benchmarkData.readEntry(parser.Value(bytes))
				}
			case 3:
				if gobytes.Equal(parser.Key(bytes), boundsName) {
					benchmarkData.readBounds(parser.Value(bytes))
				}
			}
		}
		if benchmarkData.id != 147 && benchmarkData.buffer[len(benchmarkData.buffer)-3][0] != 20 {
			t.Error("wrong result:", benchmarkData.id, benchmarkData.buffer[len(benchmarkData.buffer)-3][0])
		}
		if benchmarkData.buffer[100*3][0] != 10 {
			t.Error("wrong result:", benchmarkData.buffer[100*3][0])
		}
		if benchmarkData.buffer[100*3+1][3] != 11 {
			t.Error("wrong result:", benchmarkData.buffer[100*3+1][3])
		}
		if benchmarkData.buffer[50*3+2][10] != 180 {
			t.Error("wrong result:", benchmarkData.buffer[50*3+2][10])
		}
	}
}

func BenchmarkTabFormatReader(b *testing.B) {
	bytes, err := os.ReadFile("testdata.txt")
	if err != nil {
		b.Fatal("failed load file:", err.Error())
	} else if len(bytes) == 0 {
		b.Fatal("loaded 0 bytes")
	} else {
		var parser tabformat.ByteParser
		elementName := []byte("element")
		chunksName := []byte("chunks")
		deltasName := []byte("deltas")
		collectionName := []byte("collection")
		entryName := []byte("entry")
		boundsName := []byte("bounds")
		idName := []byte("id")
		benchmarkData = newBenchmarkData()
		b.ResetTimer()
		parser.IgnoreOpenEnd = true
		for parser.Next(bytes) {
			switch parser.Indent {
			case 0:
				if gobytes.Equal(parser.Key(bytes), elementName) {
					benchmarkData.readId(parser.Value(bytes), idName)
				}
			case 1:
				key := parser.Key(bytes)
				if gobytes.Equal(key, chunksName) {
					benchmarkData.readChunksDeltas(parser.Value(bytes), 0)
				} else if gobytes.Equal(key, deltasName) {
					benchmarkData.readChunksDeltas(parser.Value(bytes), 1)
				} else if gobytes.Equal(key, collectionName) {
					benchmarkData.readCollection(parser.Value(bytes))
				}
			case 2:
				if gobytes.Equal(parser.Key(bytes), entryName) {
					benchmarkData.readEntry(parser.Value(bytes))
				}
			case 3:
				if gobytes.Equal(parser.Key(bytes), boundsName) {
					benchmarkData.readBounds(parser.Value(bytes))
				}
			}
		}
		b.StopTimer()
		if benchmarkData.id != 147 && benchmarkData.buffer[len(benchmarkData.buffer)-3][0] != 20 {
			b.Fatal("wrong result:", benchmarkData.id, benchmarkData.buffer[len(benchmarkData.buffer)-3][0])
		}
		if benchmarkData.buffer[100*3][0] != 10 {
			b.Fatal("wrong result:", benchmarkData.buffer[100*3][0])
		}
		if benchmarkData.buffer[100*3+1][3] != 11 {
			b.Fatal("wrong result:", benchmarkData.buffer[100*3+1][3])
		}
		if benchmarkData.buffer[50*3+2][10] != 180 {
			b.Fatal("wrong result:", benchmarkData.buffer[50*3+2][10])
		}
		benchmarkData = nil
	}
}

func BenchmarkTabFormatReaderRaw(b *testing.B) {
	bytes, err := os.ReadFile("testdata.txt")
	if err != nil {
		b.Fatal("failed load file:", err.Error())
	} else if len(bytes) == 0 {
		b.Fatal("loaded 0 bytes")
	} else {
		var parser tabformat.ByteParser
		benchmarkData = newBenchmarkData()
		b.ResetTimer()
		parser.IgnoreOpenEnd = true
		for parser.Next(bytes) {
			benchmarkData.id++
		}
		b.StopTimer()
		if benchmarkData.id != 2664 {
			b.Fatal("wrong result:", benchmarkData.id)
		}
		benchmarkData = nil
	}
}

func BenchmarkTabFormatReaderRawC(b *testing.B) {
	bytes, err := os.ReadFile("testdata.txt")
	if err != nil {
		b.Fatal("failed load file:", err.Error())
	} else if len(bytes) == 0 {
		b.Fatal("loaded 0 bytes")
	} else {
		benchmarkData = newBenchmarkData()
		b.ResetTimer()
		counter := parseC(bytes)
		b.StopTimer()
		if counter != 2664 {
			b.Fatal("wrong result:", counter)
		}
		benchmarkData = nil
	}
}
