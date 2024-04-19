package main

import "testing"

func TestNewDeck(t *testing.T){
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected length of deck is 16, but got %d", len(d))
	}

	if d[0] != "Aces of Spades"{
		t.Errorf("Expected \"Aces of Spades\" at first value, but got \"%s\"", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected \"Four of Clubs\" at last value, but got \"%s\"", d[len(d) -1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T){
	
}