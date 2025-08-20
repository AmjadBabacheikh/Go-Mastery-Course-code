package main

import "testing"

func TestAdd(t *testing.T) {
	res := Add(2, 3)
	if res != 5 {
		t.Error("Expected 5 but we got ", res)
	}
}
