package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("expected lneght of 16 but got %v", len(d))
	}
	if d[0] != "ace of spades" {
		t.Errorf("expected aos  but got %v", d[0])
	}
	if d[len(d)-1] != "four of clubs" {
		t.Errorf("expected foc but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("expected 16 but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
