package main

import "testing"

//func TestHello(t *testing.T) {
//	got := Hello()
//	want := "Hello, world!1"
//
//	if got != want {
//		t.Errorf("got %q want %q", got, want)
//	}
//}

func TestHelloYou(t *testing.T) {
	got := Hello("Omar")
	want := "Hello, Omar"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
