package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits { //replace index variable with _ since it's not used
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() { //(d deck) is the receiver, its how we can tie functions to a data type to make it oop
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) { //we didn't use a receiver for deck because we are returning a type as deck
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",") //converts to a slice of string and joins on ,
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) //assigns byteSlice to bs and possible error object to err if there is an issue, otherwise will be nil
	if err != nil {
		// Log the error and quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() { //no need to return anything since it just randomizes
	source := rand.NewSource(time.Now().UnixNano()) // need to create a seed for our random number generator, need a int64 so using time UnixNano
	r := rand.New(source)                           //creates a new rand object

	for i := range d { //using i for index
		newPosition := r.Intn(len(d) - 1) // r is our rand object

		d[i], d[newPosition] = d[newPosition], d[i] //swaps values at index at each iteration
	}
}
