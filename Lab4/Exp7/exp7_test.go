package main

import "testing"

func TestDivide(t *testing.T) {
	got, err := Divide(10, 4)
	if err != nil {
		t.Fatalf("Divide(10, 4) returned unexpected error: %v", err)
	}
	if got != 2.5 {
		t.Errorf("Divide(10, 4) = %g, want 2.5", got)
	}

	_, err = Divide(7, 0)
	if err == nil {
		t.Fatal("Divide(7, 0) should return an error, got nil")
	}
	want := "cannot divide 7 by zero"
	if err.Error() != want {
		t.Errorf("Divide(7, 0) error = %q, want %q", err.Error(), want)
	}
}

func TestSquareRoot(t *testing.T) {
	got, err := SquareRoot(81)
	if err != nil {
		t.Fatalf("SquareRoot(81) returned unexpected error: %v", err)
	}
	if got != 9 {
		t.Errorf("SquareRoot(81) = %g, want 9", got)
	}

	_, err = SquareRoot(-4)
	if err == nil {
		t.Fatal("SquareRoot(-4) should return an error, got nil")
	}
	want := "cannot take square root of negative number -4"
	if err.Error() != want {
		t.Errorf("SquareRoot(-4) error = %q, want %q", err.Error(), want)
	}
}

func TestAverage(t *testing.T) {
	got, err := Average([]float64{2, 4, 6, 8})
	if err != nil {
		t.Fatalf("Average returned unexpected error: %v", err)
	}
	if got != 5 {
		t.Errorf("Average([2 4 6 8]) = %g, want 5", got)
	}

	_, err = Average(nil)
	if err == nil {
		t.Fatal("Average(nil) should return an error, got nil")
	}
	want := "cannot find average of an empty list"
	if err.Error() != want {
		t.Errorf("Average(nil) error = %q, want %q", err.Error(), want)
	}
}
