
package main;

import "fmt";

func main() {
  const conferenceName = "Go Conference"
  const conferenceSeats = 50
  var conferenceFreeSeats = 30
  
  fmt.Printf("Welcome to %v booking application\nWe have total of %v tickets and %v are still available\nGet your tickets here to attend.", conferenceName, conferenceSeats, conferenceFreeSeats)

}
