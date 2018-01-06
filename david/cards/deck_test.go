package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { //t is our test handler
	d := newDeck()

	if len(d) != 16 { //checking if we have 16 cards in our deck
		t.Errorf("Expected deck length of 16, but got %v", len(d)) //%v will allow interpolation with the value
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Epected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
