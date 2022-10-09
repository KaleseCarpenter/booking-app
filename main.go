package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Astro Tech Conference"
	const conferenceTickets = 50 // cannot change value
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
	// Allow people to keep booking
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		// this is like this because we get two values for each iteration
		//the position (index), and what you want to name the element itself (booking)
		for _, booking := range bookings {
			var names = strings.Fields(booking)       // gets first name, & lastname
			firstNames = append(firstNames, names[0]) // add variable & element you are adding to the slice, so first name
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
	}

}
