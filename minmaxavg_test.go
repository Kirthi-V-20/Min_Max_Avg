package main

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	actual := largest([]int{5, 7, 3})
	expected := 7
	if actual != expected {
		fmt.Errorf("Expected %d but got %d", expected, actual)
	}
}
