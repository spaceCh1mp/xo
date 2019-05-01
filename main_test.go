package main

import "testing"

//test if a set is a subset of another set
func TestCheckWin(t *testing.T) {
	var g = []int{1, 2, 3, 4}
	var j = []int{5, 3, 4}
	i := CheckWin(g, j)
	if i == true {
		t.Error("Expected false, got ", i)
	}
}

//matrix coordinates test
func TestFindCoordinates(t *testing.T) {
	//the location of "16" in a 4x4 matrix
	value, size := 15, 4
	x, y := findCoordinates(value, size)
	if x != 3 || y != 2 {
		t.Errorf("Expected (3,2), got (%d,%d)", x, y)
	}
}

//test for who plays next
func TestFindTurn(t *testing.T) {
	got := findTurn(0)
	expected := 1
	if got != expected {
		t.Errorf("Got %d, expected %d", got, expected)
	}
}
