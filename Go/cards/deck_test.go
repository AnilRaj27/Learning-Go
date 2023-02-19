package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("expected 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("expected aces but got %v", d[0])
	}
}
