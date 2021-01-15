package stringOccurrences

import (
	"reflect"
	"testing"
)

func TestFirst(t *testing.T) {
	if !reflect.DeepEqual(false, IsStringOccurrences("eabc")) {
		t.Error("eabc error")
	}
	if !reflect.DeepEqual(true, IsStringOccurrences("abc")) {
		t.Error("abc error")
	}
	if !reflect.DeepEqual(true, IsStringOccurrences("aeabc")) {
		t.Error("aeabc error")
	}
	if !reflect.DeepEqual(true, IsStringOccurrences("defgh")) {
		t.Error("defgh error")
	}
	if !reflect.DeepEqual(true, IsStringOccurrences("clicag")) {
		t.Error("clicag error")
	}
	if !reflect.DeepEqual(true, IsStringOccurrences("clica")) {
		t.Error("clica error")
	}
	if reflect.DeepEqual(true, IsStringOccurrences("clicart")) {
		t.Error("clicart error")
	}
	if !reflect.DeepEqual(true, IsStringOccurrences("fiance")) {
		t.Error("fiance error")
	}
}