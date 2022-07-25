
package main

import "fmt"

func main() {
  const conferenceName string = "Go Conference"
  const conferenceTickets int = 50
  var remainingTickets int = 30
  adHockString := "This form of declariation doesn't work with const"
  adHockNumber := 4

  fmt.Printf("conferenceName is %T, remainingTickets is %T, conferenceTickets is %T\n", conferenceName, remainingTickets, conferenceTickets)


}
