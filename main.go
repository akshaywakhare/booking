package main

import "fmt"

func main() {
	var conName = "GoCon"
	const conTickets = 50
	var remainingTickets = conTickets

	fmt.Println("Welcome!")
	fmt.Printf("%v Tickets remaining...\n", remainingTickets)
	fmt.Println("Proceed to booking for", conName, "---->")

}
