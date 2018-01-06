//program to read a filename as a command line argument. Then prints the file to the console

package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open(os.Args[1]) //passing in command line args
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999) //creates an empty byte slice to hold up to 99999 items
	file.Read(bs)
	fmt.Println(string(bs))
}
