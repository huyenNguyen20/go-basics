package main

import (
	"fmt"
	"io"
	"os"
)

type content  struct {}

func main() {
	
	filename := os.Args[1]

	file, err := os.Open(filename) // For read access.

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	ct := content{}
	io.Copy(ct, file)
}

func (content) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote that many bytes: ", len(bs))
	return len(bs), nil
}