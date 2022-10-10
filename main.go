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

		// User Input Validation
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		// Email Validator
		isValidEmail := strings.Contains(email, "@") // gives true or false back
		// User input Validator to make sure user input whole psoitive number greater than 0
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		//isValidCity := city = "Singapore" || city == "London"

		if isValidName && isValidEmail && isValidTicketNumber {
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

			if remainingTickets == 0 {
				// end program
				fmt.Printf("Our %v is booked out. Please come back next year!", conferenceName)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered does not contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you enetered is invalid")
			}

			// continue // skips to next iteration to start over again so they can have another chance
		}

	}

}
