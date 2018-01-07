// program to check satus of websites in a slice of string
package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) //creates a channel of type string to use for our go subroutines

	for _, link := range links { //will iterate through links variable, we don't care about the indexes so we can use _
		go checkLink(link, c) //go keyword creates a new Go Routine to run checkLink so it wont be blocking to make code run faster
	}

	fmt.Println(<-c) //Wait for value to be sent to channel. When we get one, log it out right away
}

func checkLink(link string, c chan string) { //c is the channel variable, chan is the type and string is data type
	_, err := http.Get(link) //we don't care about the response, only the err so use _
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I thinnk" //sends message to channel
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep its up" //sends message to channel
}
