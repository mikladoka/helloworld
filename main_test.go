package main

import "testing"

func TestHelloWorld(t *testing.T) {
	if helloworld() != "Helo World!!" {
		t.Fatal("Test fail")
	}
}
