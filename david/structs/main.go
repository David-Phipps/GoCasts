package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	david := person{
		firstName: "David",
		lastName:  "Phipps",
		contact: contactInfo{
			email:   "david@gmail.com",
			zipCode: 78759, //every line needs a , even if its the last item
		},
	}
	david.updateName("Davey")
	david.print()
}

func (pointerToPerson *person) updateName(newFirstName string) { //*pointer gvies the value at the memory address, gives direct access to the value at the memory address. * in front of a type
	//is different than a * in front of in front of a pointer. * in front of a type means this is a description of a type (*pointer is a type of pointer to person and can only be called on a receiver
	//of pointer to person. *pointerToPerson means take this memory address and turn it into a value. )
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// Turn address into value with *. If we use a * in frnt of a type that means we are looking for a data type * to type
// Turn value into address with &

// func (pointerToPerson *person) updateName(newFirstName string) { //*pointer gvies the value at the memory address, gives direct access to the value at the memory address. * in front of a type
// 	//is different than a * in front of in front of a pointer. * in front of a type means this is a description of a type (*pointer is a type of pointer to person and can only be called on a receiver
// 	//of pointer to person. *pointerToPerson means take this memory address and turn it into a value. )
// 	(*pointerToPerson).firstName = newFirstName
// }
