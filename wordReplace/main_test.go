package main

import (
	"reflect"
	"testing"
)

func TestFirst(t *testing.T) {
	if !reflect.DeepEqual(true, WordReplace("first", "first")) {
		t.Error("first and first case")
	}

	if !reflect.DeepEqual(true, WordReplace("first", "firsd")) {
		t.Error("first and firsd case")
	}

	if !reflect.DeepEqual(false, WordReplace("first", "firzd")) {
		t.Error("first and firzd case")
	}

	if !reflect.DeepEqual(true, WordReplace("first", "irst")) {
		t.Error("first and irst case")
	}

	if !reflect.DeepEqual(true, WordReplace("first", "firsts")) {
		t.Error("first and firsts case")
	}

	if !reflect.DeepEqual(false, WordReplace("firsts", "second")) {
		t.Error("firsts and second case")
	}

	if !reflect.DeepEqual(true, WordReplace("first你好", "firsd你好")) {
		t.Error("first你好 and firsd cut你好")
	}
}