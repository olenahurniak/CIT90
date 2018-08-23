package main

import (
	"fmt"
)

func main() {
	fmt.Println("Begin")
	
	foo("James Bond")

	x := "Miss Monypenny"
	foo(x)

	fmt.Println("about to end")
	
}

func foo(name string) {
	fmt.Println("Hello")
}

// foo takes a value 