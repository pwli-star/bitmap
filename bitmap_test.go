package bitmap_test

import (
	"bitmap"
	"testing"
)

func TestBitmapSort(t *testing.T) {
	bmap := bitmap.NewBitmap(15)

	original := [6]uint64{4, 6, 8, 1, 7, 15}
	expected := [6]uint64{1, 4, 6, 7, 8, 15}
	actual := [6]uint64{}

	for _, offset := range original {
		bmap.Set(offset, 1)
	}

	var i uint64
	var offset, maxpos uint64 = 0, bmap.Maxpos() + 1
	for ; offset < maxpos; offset++ {
		if bmap.Get(offset) == 1 {
			actual[i] = offset
			i++
		}
	}

	if expected != actual {
		t.Errorf("expected:%#v, actual:%#v", expected, actual)
	}
}

func TestBitmap_Reset(t *testing.T) {
	bmap := bitmap.Default()
	original := [6]uint64{4, 6, 8, 1, 7, 15}

	for _, offset := range original {
		bmap.Set(offset, 1)
	}
	bmap.Reset()
	for i := 0; i < 16; i++ {
		if bmap.Get(uint64(i)) != 0 {
			t.Fatalf("error: not clear at offset %d", i)
		}
	}
}
