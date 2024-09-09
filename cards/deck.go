package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of 'deck'
//which is a slice of strings

type deck []string

/*
In go there is no OOP concept instead we want to "extend" a base type
and add some extra functionality to it

	type deck []string - tell go we want to create an array og strings and add a bunch of functions specifically made to work with it.

	A function with a receiver is like a "method" - a function that belongs to an "instance"
*/

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "jack", "Qween", "King"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}
	return cards
}

/*
_ - use the underscore to tell go that variable is not being used
*/

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/*
	func (d deck) print()
		in above (d deck) is the receiver on a function.
		In above example deck is slice, d accept deck that means it's a slice, all variable is using deck can have access to print func
		d - the actual copy of the deck we are working with is available in the function as a variable called 'd'
		deck - Every variable of type 'deck' can call this function on itself

*/

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}

/*

slice[startIndexIncluding : upToNotIncluding]
slice[: x] - starting from 0 to x not including
slice[y:] - starting from y to size
return multiple values in the above function.
*/

func (d deck) toString() string {

	// basically deck is a slice of string so it's possible in conversion
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0777)
}

/*
In the above line we are type casting string to []byte
*/

func newDeckFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)
	if err != nil {
		// option1 : log the error and return a call to newDeck()
		// option2 : Log the error and quit the program.
		fmt.Println("Error:", err)
		os.Exit(1)

	}

	s := string(bs)
	data := strings.Split(s, ",")
	return deck(data)

}

/*
Value of type 'error' if nothing went wrong it will be nil
we are using strings split function for splting the string
also casting it to string then deck
*/

func (d deck) shuffle() {
	rand.Int()
	for i := range d { // here we don't need the value so we can us this.
		np := rand.Intn(len(d) - 1) // using rand package from math/rand to generate random number

		d[i], d[np] = d[np], d[i] // here we are doing swap operation in a single line
	}

	/*
		if it's not working

		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source);/
		newPosition := r.Intn(len(d) - 1)
	*/
}
