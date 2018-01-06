package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{} //custom type that will have Write function

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// 1. can output this way but has extra lines of code to implement
	// bs := make([]byte, 99999) //make is a built in function that takes a type of slice as first arg and the number of elements you would want in the slice as second arg so we have space for 99999 elements
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// second exmplae
	//io.Copy(os.Stdout, resp.Body) //2. os.Stdout implements the writer interface and resp.Body has the writer interface

	//third example with custom type
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) { //3. custom implementation,  logWriter implements Writer interface since it has the Write function
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil

}
