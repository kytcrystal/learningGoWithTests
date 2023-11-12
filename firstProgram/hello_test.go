package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()  //Declaring the got variable
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloChris(t *testing.T) {
	got := HelloName("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}