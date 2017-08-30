package main

func main() {

	// cards := newDeck()
	cards := newDeckFromFile("my_cards");
	// cards.saveToFile("my_cards")
	cards.shuffle()
	cards.print()
}