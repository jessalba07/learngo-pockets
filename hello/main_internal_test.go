package main

import "testing"

func TestGreet( t *testing.T) {
	want := "Hello, World!"

	got := greet()

	if got != want {
		// mark this test as failed
		t.Errorf("expected: %q, got: %q", want, got)
	}
}