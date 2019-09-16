package main

import (
	"rsc.io/quote"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := quote.Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
