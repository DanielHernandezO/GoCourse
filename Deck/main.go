package main

func main() {
	cards := newDeck()
	//fileName := "my_cards.txt"
	// cards.saveToFile(fileName)
	//cards2 := newDeckFromFile(fileName)
	cards.shuffle()
	cards.print()
}
