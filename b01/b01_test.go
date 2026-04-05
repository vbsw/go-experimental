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

package b01

import (
	"sort"
	"testing"
)

// BenchmarkBinarySearchCustom does benchmark on function BinarySearchCustom.
func BenchmarkBinarySearchCustom(b *testing.B) {
	numbers := []int{1, 3, 10, 12, 15, 16, 17, 30, 32, 33, 34, 36, 38, 48, 51, 55, 61, 62, 63, 64, 67, 68, 74, 78, 80, 81, 82, 83, 84, 90, 95, 97, 99, 106, 110, 112, 118, 119, 130, 135, 137, 143, 145, 148, 150, 159, 160, 161, 170, 179, 182, 185, 192, 193, 195, 199, 201, 203, 210, 212, 215, 216, 217, 230, 232, 233, 234, 236, 238, 248, 251, 255, 261, 262, 263, 264, 267, 268, 274, 278, 280, 281, 282, 283, 284, 290, 295, 297, 299, 306, 310, 312, 318, 319, 330, 335, 337, 343, 345, 348, 350, 359, 360, 361, 370, 379, 382, 385, 392, 393, 395, 399, 401, 403, 410, 412, 415, 416, 417, 430, 432, 433, 434, 436, 438, 448, 451, 455, 461, 462, 463, 464, 467, 468, 474, 478, 480, 481, 482, 483, 484, 490, 495, 497, 499, 506, 510, 512, 518, 519, 530, 535, 537, 543, 545, 548, 550, 559, 560, 561, 570, 579, 582, 585, 592, 593, 595, 599, 601, 603, 610, 612, 615, 616, 617, 630, 632, 633, 634, 636, 638, 648, 651, 655, 661, 662, 663, 664, 667, 668, 674, 678, 680, 681, 682, 683, 684, 690, 695, 697, 699, 706, 710, 712, 718, 719, 730, 735, 737, 743, 745, 748, 750, 759, 760, 761, 770, 779, 782, 785, 792, 793, 795, 799, 801, 803, 810, 812, 815, 816, 817, 830, 832, 833, 834, 836, 838, 848, 851, 855, 861, 862, 863, 864, 867, 868, 874, 878, 880, 881, 882, 883, 884, 890, 895, 897, 899, 906, 910, 912, 918, 919, 930, 935, 937, 943, 945, 948, 950, 959, 960, 961, 970, 979, 982, 985, 992, 993, 995, 999}
	targets := []int{1, 500, 999, 22, 51, 351, 338, 571, 748, 935, 275, 557, 817, 840, 412, 413, 414, 630, 782, 25, 980}
	matches := []bool{true, false, true, false, true, false, false, false, true, true, false, false, true, false, true, false, false, true, true, false, false}
	var counter int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counter = 0
		for j := 0; j < len(targets); j++ {
			_, match := BinarySearchCustom(numbers, targets[j])
			if matches[j] == match {
				counter++
			}
		}
	}
	if counter != len(matches) {
		b.Fatal("wrong result", counter)
	}
}

// BenchmarkBinarySearchSTD does benchmark on standard function sort.Search.
func BenchmarkBinarySearchSTD(b *testing.B) {
	numbers := []int{1, 3, 10, 12, 15, 16, 17, 30, 32, 33, 34, 36, 38, 48, 51, 55, 61, 62, 63, 64, 67, 68, 74, 78, 80, 81, 82, 83, 84, 90, 95, 97, 99, 106, 110, 112, 118, 119, 130, 135, 137, 143, 145, 148, 150, 159, 160, 161, 170, 179, 182, 185, 192, 193, 195, 199, 201, 203, 210, 212, 215, 216, 217, 230, 232, 233, 234, 236, 238, 248, 251, 255, 261, 262, 263, 264, 267, 268, 274, 278, 280, 281, 282, 283, 284, 290, 295, 297, 299, 306, 310, 312, 318, 319, 330, 335, 337, 343, 345, 348, 350, 359, 360, 361, 370, 379, 382, 385, 392, 393, 395, 399, 401, 403, 410, 412, 415, 416, 417, 430, 432, 433, 434, 436, 438, 448, 451, 455, 461, 462, 463, 464, 467, 468, 474, 478, 480, 481, 482, 483, 484, 490, 495, 497, 499, 506, 510, 512, 518, 519, 530, 535, 537, 543, 545, 548, 550, 559, 560, 561, 570, 579, 582, 585, 592, 593, 595, 599, 601, 603, 610, 612, 615, 616, 617, 630, 632, 633, 634, 636, 638, 648, 651, 655, 661, 662, 663, 664, 667, 668, 674, 678, 680, 681, 682, 683, 684, 690, 695, 697, 699, 706, 710, 712, 718, 719, 730, 735, 737, 743, 745, 748, 750, 759, 760, 761, 770, 779, 782, 785, 792, 793, 795, 799, 801, 803, 810, 812, 815, 816, 817, 830, 832, 833, 834, 836, 838, 848, 851, 855, 861, 862, 863, 864, 867, 868, 874, 878, 880, 881, 882, 883, 884, 890, 895, 897, 899, 906, 910, 912, 918, 919, 930, 935, 937, 943, 945, 948, 950, 959, 960, 961, 970, 979, 982, 985, 992, 993, 995, 999}
	targets := []int{1, 500, 999, 22, 51, 351, 338, 571, 748, 935, 275, 557, 817, 840, 412, 413, 414, 630, 782, 25, 980}
	matches := []bool{true, false, true, false, true, false, false, false, true, true, false, false, true, false, true, false, false, true, true, false, false}
	var counter int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counter = 0
		for j := 0; j < len(targets); j++ {
			target := targets[j]
			index := sort.Search(len(numbers), func(k int) bool {
				return numbers[k] >= target
			})
			if matches[j] == bool(index < len(numbers) && numbers[index] == target) {
				counter++
			}
		}
	}
	if counter != len(matches) {
		b.Fatal("wrong result", counter)
	}
}

// BenchmarkBinarySearchGEN does benchmark on function BinarySearchGEN.
func BenchmarkBinarySearchGEN(b *testing.B) {
	numbers := []int{1, 3, 10, 12, 15, 16, 17, 30, 32, 33, 34, 36, 38, 48, 51, 55, 61, 62, 63, 64, 67, 68, 74, 78, 80, 81, 82, 83, 84, 90, 95, 97, 99, 106, 110, 112, 118, 119, 130, 135, 137, 143, 145, 148, 150, 159, 160, 161, 170, 179, 182, 185, 192, 193, 195, 199, 201, 203, 210, 212, 215, 216, 217, 230, 232, 233, 234, 236, 238, 248, 251, 255, 261, 262, 263, 264, 267, 268, 274, 278, 280, 281, 282, 283, 284, 290, 295, 297, 299, 306, 310, 312, 318, 319, 330, 335, 337, 343, 345, 348, 350, 359, 360, 361, 370, 379, 382, 385, 392, 393, 395, 399, 401, 403, 410, 412, 415, 416, 417, 430, 432, 433, 434, 436, 438, 448, 451, 455, 461, 462, 463, 464, 467, 468, 474, 478, 480, 481, 482, 483, 484, 490, 495, 497, 499, 506, 510, 512, 518, 519, 530, 535, 537, 543, 545, 548, 550, 559, 560, 561, 570, 579, 582, 585, 592, 593, 595, 599, 601, 603, 610, 612, 615, 616, 617, 630, 632, 633, 634, 636, 638, 648, 651, 655, 661, 662, 663, 664, 667, 668, 674, 678, 680, 681, 682, 683, 684, 690, 695, 697, 699, 706, 710, 712, 718, 719, 730, 735, 737, 743, 745, 748, 750, 759, 760, 761, 770, 779, 782, 785, 792, 793, 795, 799, 801, 803, 810, 812, 815, 816, 817, 830, 832, 833, 834, 836, 838, 848, 851, 855, 861, 862, 863, 864, 867, 868, 874, 878, 880, 881, 882, 883, 884, 890, 895, 897, 899, 906, 910, 912, 918, 919, 930, 935, 937, 943, 945, 948, 950, 959, 960, 961, 970, 979, 982, 985, 992, 993, 995, 999}
	targets := []int{1, 500, 999, 22, 51, 351, 338, 571, 748, 935, 275, 557, 817, 840, 412, 413, 414, 630, 782, 25, 980}
	matches := []bool{true, false, true, false, true, false, false, false, true, true, false, false, true, false, true, false, false, true, true, false, false}
	var counter int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counter = 0
		for j := 0; j < len(targets); j++ {
			_, match := BinarySearchGEN(numbers, targets[j])
			if matches[j] == match {
				counter++
			}
		}
	}
	if counter != len(matches) {
		b.Fatal("wrong result", counter)
	}
}
