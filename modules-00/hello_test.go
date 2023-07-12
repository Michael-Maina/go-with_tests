package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Tadej")
	want := "Hello, Tadej"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
