package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
	//the above line is is also valid
}

/*
struct - will allow us to create a type with multiple properies
This is how we will create custom type in go
* above contact it's called embedding structs
*/

func main() {
	//alex := person{"Alex", "Anderson"}

	// alex := person{firstName: "Alex", lastName: "Anderson"}

	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	// jim := person{
	// 	firstName: "Jim",
	// 	lastName:  "Party",
	// 	contact: contactInfo{
	// 		zipCode: 32421,
	// 		email:   "jim@gmail.com"},
	// }
	// fmt.Println(jim)
	// commented for removing the contact field with contactInfo
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			zipCode: 32421,
			email:   "jim@gmail.com"},
	}
	// fmt.Println(jim)
	// jim.print()
	// jim.updateName("basil")
	//commented to use pointer in new func

	// jimPointer := &jim
	// here we are passing the address of memory
	//commenting for the pointer shortcut
	jim.updateName("basil")
	jim.print()
}

/*
1 -> in the order it will assign the value to the fields in struct
2 -> we can implicitly say the field and assign value to the field
3 -> create a variable without initializing it.
%+v -> prints all the different field names and it's values
*/

func (p person) print() {
	fmt.Printf("%+v", p)
}

// receiver function

// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// 	fmt.Println(p)
// }

// if we give it won't work because it's not modifiable the reciver
// it's updated in the inside function but not to the jim in main.

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

/*
&variable - give me the memory address of the value this variable is pointing at.

*pointer - Give me the value this memory address is pointing at

*person - is accepting the pointer as parameter

if we directly use the variable to a pointer to function as a call then go will convert the memory address for us
*/
