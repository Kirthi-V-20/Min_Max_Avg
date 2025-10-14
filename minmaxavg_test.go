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

func TestMin(t *testing.T) {
	actual := smallest([]int{5, 7, 3})
	expected := 3
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestMin1(t *testing.T) {
	actual := smallest([]int{20, 50, 10})
	expected := 10
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestAvg(t *testing.T) {
	actual := average([]int{5, 7, 3})
	expected := 5.0
	if actual != expected {
		t.Errorf("Expected %.2f but got %.2f", expected, actual)
	}
}

func TestAvg1(t *testing.T) {
	actual := average([]int{1, 2, 3})
	expected := 2.0
	if actual != expected {
		t.Errorf("Expected %.2f but got %.2f", expected, actual)
	}
}
