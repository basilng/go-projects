package main

func main() {
	// create a variable
	// var card string = "Ace of Spades"
	// alternative way

	// card := newCard()
	// fmt.Println(card)

	// Declaring a slice
	// cards := deck{newCard(), "Ace of Diamonds"}
	// //update replaced `[]sting`` with deck
	// cards = append(cards, "Six of Spades")
	//update commented all code to use new deck.go
	cards := newDeck()

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	//update replaced above code with deck print

	hand, remainingCards := deal(cards, 5)
	// deal is gonna return 2 values in the order we can store it to variable
	hand.print()
	remainingCards.print()

	cards.saveToFile("my_cards")
	// saving the content to a file named my_cards

	cards2 := newDeckFromFile("my_cards")
	cards2.print()

	cards.shuffle()
	cards.print()

}

/*
var - we are about to create a new variable
card - the name of the variable
string - only a "string" will ever be assinged to this variable
= - assigning the value
go is a statically typed lang
:= - telling go to infer the value, and it's only for new variable. (Initialization)

----
how to add new elements to a slice?
	append(cards, <>) - it won't modify the existing slice instead will return a new slice and assign to the variable
how to iterate over a slice?
	for <index>, <value> := range <slice|array> {}
		index - index of this elenebt in the array
		value - current value we are iterating over
		range <array> - take the slice of array and loop over it
		:= every time we reinitializing the index & value that's whhy we are using
*/

// creating new function that retun a string

// func newCard() string {

// 	return "Five of Diamonds"

// }
//update commented the newCard since we are using the deck.go newCard

/*
newCards : name of the function
 string : after the parathisis we have to tell what value it returns, it can be any type
 return : is the keyword for returning.

*/
