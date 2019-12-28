package main

func main() {

	//hand1, hand2 := deal(cards, 5)
	//hand1.print()
	//hand2.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
