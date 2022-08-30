/*
	A package is collection of .go files
	Every .go must contain declaration of its main package
	main - is default and mandatory name for executable package
		(vs reusable packages - dependencies)
*/
package main

/**/
import "fmt"

/**/
func main() {
	fmt.Println("Hi there!")
}
