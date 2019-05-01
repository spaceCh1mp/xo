package main

import "testing"

func TestCheckWin(t *testing.T) {
	var g = []int{1, 2, 3, 4}
	var j = []int{5, 3, 4}
	i := CheckWin(g, j)
	if i == true {
		t.Error("Expected false, got ", i)
	}
}
