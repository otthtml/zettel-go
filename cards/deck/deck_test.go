package deck

import (
	"os"
	"testing"
)

func TestNewDec(t *testing.T) {
	d := NewDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestNewDeckFromFile(t *testing.T) {
	os.Remove("../my_cards.txt")

	d := NewDeck()
	d.SaveToFile("../my_cards.txt")

	nd := NewDeckFromFile("../my_cards.txt")

	if len(nd) != 16 {
		t.Errorf("Expected 16 cards, but got %v", len(nd))
	}

	os.Remove("../my_cards.txt")
}
