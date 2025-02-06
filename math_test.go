package main

import (
	"pt_aka_tech_test/math"

	"testing"
)

func TestAddition(t *testing.T){
    got := math.Add(4, 6)
    want := 10
    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestDivision(t *testing.T){
    got, _ := math.Divide(10, 2)
    want := 5
    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }

}

func TestDivideByZero(t *testing.T) {
	_, err := math.Divide(10, 0)
	want := "division by zero"
	
	if err == nil || err.Error() != want { // Compare the error messages
		t.Errorf("got %q, wanted %q", err, want)
	}
}