# Project - cards

## description 
Create a deck of cards with functions:
  * newDeck  
    * Create a list of playing cards. Basically an array of srings  
  * print  
    * logs/displays/prints contents of a ceck of cards  
  * shuffle  
    * shuffles al the cards in a deck  
  * deal  
    * create a 'hand' of cards  
  * saveToFile  
    * save a list of cards to a file locally  
  * newDeckFromFile  
    * Load a list of cards from the local storage  


## go - about packages: Files in the same package can freely call functions defined in other files.
  Question 5:
  Here's a tough one!  
  Imagine we have two files in the same folder  
  with the same package name.  Will the following  
  code successfully compile?  This might take  
  a little experimentation on your side.  
  If you do try this out yourself, run your code  
  with the command go run main.go state.go .

In main.go:
```
package main
 
func main() {
    printState()
}
```

In a separate file called state.go:
```
package main
 
import "fmt"
 
func printState() {
    fmt.Println("California")
}
```
