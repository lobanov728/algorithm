package main

import (
	"reflect"
	"testing"
)

func TestFirst(t *testing.T) {
	if !reflect.DeepEqual(true, WordReplace("cat", "cat")) {
		t.Error("cat and cut case")
	}
	
	if !reflect.DeepEqual(true, WordReplace("cat", "cut")) {
		t.Error("cat and cut case")
	}

	if !reflect.DeepEqual(false, WordReplace("cat", "cur")) {
		t.Error("cat and cur case")
	}

	if !reflect.DeepEqual(true, WordReplace("cat", "at")) {
		t.Error("cat and at case")
	}

	if !reflect.DeepEqual(true, WordReplace("cat", "cats")) {
		t.Error("cat and cats case")
	}

	if !reflect.DeepEqual(false, WordReplace("cat", "dog")) {
		t.Error("cat and dog case")
	}

	if !reflect.DeepEqual(true, WordReplace("cat你好", "cut你好")) {
		t.Error("cat你好 and cut cut你好")
	}
}