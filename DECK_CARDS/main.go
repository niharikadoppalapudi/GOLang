package main

//import "fmt"

func main() {
	//cards := deck{"Ace of Diamonds","Ace of Spades"}
	//cards = append(cards,"Six of Spades")
	cards := newDeck()
	//cards.print()
	//hand, remainingcards := deal(cards,6)
	//hand.print()
	//remainingcards.print()
	//Message:= "Hello GO!"
	//fmt.Println([]byte(Message))
	//fmt.Println(cards.toString())
	//cards.saveToFile("My_Cards")
	//cards := newDeckFromFile("My")
	cards.shuffle()
	cards.print()

}
