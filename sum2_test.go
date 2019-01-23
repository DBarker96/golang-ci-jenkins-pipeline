package main

import (
	"testing"
)

func TestBrokenSum(t *testing.T) {
	actual := Sum(4, 2)
	expected := 6
	if actual != expected {
		t.Errorf("Sum(1, 2): actual %v, expected %v", actual, expected)
	}
}
