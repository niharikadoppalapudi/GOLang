package main 

import "fmt"

func main() {
cards := [] string {newCard(),"Ace of Spades"}	
fmt.Println(cards)
cards= append(cards,"Six of spades")
fmt.Println(cards)
for index,cd := range cards{
	fmt.Println(index,cd)
}
card := newCard()
	fmt.Println(card)
	sliceOfCards := sliceOfCards()
	fmt.Println(sliceOfCards)
}
func newCard() string {
	return "Ace of Diamonds"
}
func sliceOfCards() [] string {
	cds := [] string {"spades","diamonds"}
	cds = append(cds,"Diamond")
	return cds
}