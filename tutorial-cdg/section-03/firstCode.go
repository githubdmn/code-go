package main

import "fmt"

// a variable can be declared outside of a function
// but cannot be initialized outside of a function

func afn() string {
	// var card string = newCard()
	// var card string = "Ace of Spades";
	// card:= "Ace of spades";
	// := is used only when a variable is initialised
	// = is used when variable value is reassigned or changed
	return "afn()"
}

func bfn() {
	// var cards string = []string{"Ace of Diamonds", newCard()}
	// 	- cannot use []string{…} (value of type []string) as type string in variable declaration
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	// fmt.Println(cards)
	for i, card := range cards {
		fmt.Println(i, card)
	}
	/*
		Every variable declared must be used in the code.
	*/
}

func newCard() string {
	return "Five of Diamonds"
}

//func main() {}
func run() {}
