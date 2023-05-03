package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// get the file name from the command line argument list
	// os.Args is the list of command line arguments passed to our program
	// the first element in that slice of type string is going to be our executable file that
	// represents our main.go process
	// fmt.Println(os.Args[1])

	// ------
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
