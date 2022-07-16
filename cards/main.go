package main

import "fmt"

func main() {
	fmt.Println("Generating new deck, shuffling and saving to file")

	cards := NewDeck()
	cards.SaveToFile("my_cards.txt")

	cards = NewDeckFromFile("my_cards.txt")
	cards.Shuffle()
	cards.SaveToFile("my_cards.txt")
}
