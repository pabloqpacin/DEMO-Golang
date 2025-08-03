package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	err := Greet(&buffer, "Chris")
	if err != nil {
		t.Errorf("Greet returned an error: %v", err)
	}

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
