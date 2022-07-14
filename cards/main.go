package main

import (
	d "otthtml/cards/deck"
)

func main() {
	cards := d.NewDeck()
	cards.SaveToFile("my_cards.txt")

	cards = d.NewDeckFromFile("my_cards.txt")
	cards.Shuffle()
	cards.SaveToFile("my_cards.txt")
}
