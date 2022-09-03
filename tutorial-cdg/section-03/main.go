package main

import "fmt"

// a variable can be declared outside of a function
// but cannot be initialized outside of a function

func main() {
	var card string = "Ace of spades"
	// card:= "Ace of spades"
	// := is used only when a variable is initialised
	// = is used when variable value is reassigned or changed

	fmt.Println(card)
}
