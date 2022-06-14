package main

import "testing"

func TestFactorial(t *testing.T) {
	expected := 24
	testValue := 4
	got := factorial(testValue)
	if got != expected {
		t.Fail()
	}

}
