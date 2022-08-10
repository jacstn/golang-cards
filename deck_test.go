package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(deck))
	}

	if deck[0] != "As" {
		t.Errorf("Expected As, but got %v", deck[0])
	}
}

func TestSaveDeckAndLoadDeck(t *testing.T) {
	fileName := ("__filetesting")
	os.Remove(fileName)

	deck := newDeck()

	deck.saveDeck(fileName)

	loadedDeck := loadDeck(fileName)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected loaded deck length of 16, but got %v", len(loadedDeck))
	}

	os.Remove(fileName)
}
