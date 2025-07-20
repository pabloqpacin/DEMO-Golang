package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("pabloqpacin")
	want := "Hello, pabloqpacin"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
