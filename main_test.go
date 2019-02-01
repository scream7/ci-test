package main

import "testing"

func TestF(t *testing.T) {
	if f("aa") != "Hello, aa" {
		t.Fail()
	}
}
