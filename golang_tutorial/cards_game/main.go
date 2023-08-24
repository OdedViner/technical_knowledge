package main

func main() {
	cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	cards.print()

	// hand, remainingCards := deal(cards, 2)
	// hand.print()
	// remainingCards.print()
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
}
