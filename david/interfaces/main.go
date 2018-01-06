package main

import "fmt"

type bot interface { //we create an interface that groups types that have the getGreeting function that returns a string into a type called bot. Can list multiple functions
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) { // by creating an interface of type bot, we can reuse code and call printGreeting on our englishBot and spanishBot
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// VERY custom logic for getnerating an english greeting
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
