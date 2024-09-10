package main

import "fmt"

type bot interface {
	getGreeting() string
}

/*
Creating an interface
how interface work is in this program there is any type which has "getGreeting() string"
then u r a member of that interface
then the both struct can call the printGreeting
we don't need to type variable names in arguments for interface.
all functions needs to be used to qualify
*/
type englishBot struct {
}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hello!"
}

/*
In receiver if we are not using the variable then we can remove it
*/

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sp spanishBot,w string) {
// 	fmt.Println(sp.getGreeting())
// }

// method overloading with different parameters are not allowed in go.
// this is the issue of struct without interface we will use interface to solve this.

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

/*
1. we used interface to define a method, which ever method is used that method then it's a member
2. All members of the interface can use the interface methods
*/

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
