package main

import (
	"testing"
)

func TestMax(t *testing.T) {
	actual := largest([]int{5, 7, 3})
	expected := 7
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestMax1(t *testing.T) {
	actual := largest([]int{20, 50, 10})
	expected := 50
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
