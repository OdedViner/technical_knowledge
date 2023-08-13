package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 6 {
		t.Error("Expected is 6, but actaul is", len(d))
	}
}
