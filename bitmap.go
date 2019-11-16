/*
 * MIT License
 *
 * Copyright (c) 2019. pwli
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package bitmap

// the max number is 4294967296(1 << 32) by default
const DefaultSize = 0x01 << 32

type Bitmap struct {
	// inner byte slice
	data []byte
	// maxNum indicates the max number it can hold
	maxNum uint64
	// the first place where bit is set to 1
	maxpos uint64
}

// Default is shortcut for creating a BitMap with default max number
func Default() *Bitmap {
	return NewBitmap(DefaultSize)
}

// NewBitmap create a BitMap with a gaven max number
func NewBitmap(maxNum uint64) *Bitmap {
	if maxNum == 0 {
		maxNum = DefaultSize
	}
	return &Bitmap{data: make([]byte, (maxNum>>3)+1), maxNum: maxNum}
}

// Set will set the bit at offset position to a gaven value
func (bmap *Bitmap) Set(offset uint64, value uint8) bool {
	if bmap.maxNum < offset {
		return false
	}
	index, pos := offset/8, offset%8

	if value == 0 {
		// &^ clear the bit
		bmap.data[index] &^= 0x01 << pos
	} else {
		bmap.data[index] |= 0x01 << pos

		// Record the first place where bit is set to 1
		if bmap.maxpos < offset {
			bmap.maxpos = offset
		}
	}

	return true
}

// Get Returns the bit value at offset position
func (bmap *Bitmap) Get(offset uint64) uint8 {
	if bmap.maxNum < offset {
		return 0
	}

	index, pos := offset/8, offset%8
	return (bmap.data[index] >> pos) & 0x01
}

// Maxpos Returns the first position where bit is set to 1
func (bmap *Bitmap) Maxpos() uint64 {
	return bmap.maxpos
}

// Reset clear all the bits and set them to 0
func (bmap *Bitmap) Reset() {
	bmap.data = make([]byte, len(bmap.data))
	bmap.maxpos = 0
}
