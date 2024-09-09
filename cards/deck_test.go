package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	expected := "Ace of Spades"
	actual := d[0]
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)

	}
	expected = "King of Clubs"
	actual = d[len(d)-1]
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

/*
t *testing.T - In the above code t is the test handler
%v - the value will take from the argument
*/

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	fn := "_decktesting"
	os.Remove(fn)
	d := newDeck()
	d.saveToFile(fn)
	ld := newDeckFromFile(fn)
	if len(ld) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(ld))
	}
	os.Remove(fn)
}
