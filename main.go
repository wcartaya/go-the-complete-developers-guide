package main

func main() {
	/*
		cards := newDeck()

		hand, remainingCards := deal(cards, 5)

		hand.print()
		remainingCards.print()
	*/
	cards := newDeck()
	/*
		cards.saveToFile("temp.txt")

		newDecks := newDeckFromFile("temp.txt")
	*/
	cards.shuffle()
	cards.print()
}
