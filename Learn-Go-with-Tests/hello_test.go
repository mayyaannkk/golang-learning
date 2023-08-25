package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		// will print out a message and fail the test
		t.Errorf("got %q want %q", got, want)
	}
}
