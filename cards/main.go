package main

func main() {
	cards := NewDeck()
	cards.SaveToFile("my_cards.txt")

	cards = NewDeckFromFile("my_cards.txt")
	cards.Shuffle()
	cards.SaveToFile("my_cards.txt")
}
