package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {

	buffer := bytes.Buffer{}

	Greet(&buffer, "Sergio")

	got := buffer.String()
	want := "Hello, Sergio"


	
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}