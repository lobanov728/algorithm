package unionFind_2

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFirst(t *testing.T) {
	uf := NewUF(8)
	uf.Union(3, 4)

	if reflect.DeepEqual(false, uf.Connected(3, 4)) {
		t.Error("3 and 4 are connected")
	}

	if reflect.DeepEqual(true, uf.Connected(1, 2)) {
		t.Error("1 and 2 not connected")
	}

	uf.Union(1, 2)
	if reflect.DeepEqual(false, uf.Connected(1, 2)) {
		t.Error("1 and 2 connected")
	}

	uf.Union(2, 3)
	if reflect.DeepEqual(false, uf.Connected(1, 4)) {
		t.Error("1 and 4 connected")
	}

	uf.Union(1, 7)
	if reflect.DeepEqual(false, uf.Connected(7, 3)) {
		t.Error("3 and 7 connected")
	}

	uf1 := NewUF(13)
	uf1.Union(6, 8)
	uf1.Union(10, 8)
	uf1.Union(8, 0)

	uf1.Union(11,12)
	uf1.Union(11,9)
	uf1.Union(0, 9)

	fmt.Println(uf1.Connected(12,10))
}