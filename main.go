package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking applicaiton \n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	// formated print

	// reference variable with user input value
	//fixed element in array

	// var bookings = [50]string{}

	for {
		var userTickets uint
		var firstName string
		var lastName string
		var email string

		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("The whole slice %v\n", bookings)
			fmt.Printf("The first slice %v\n", bookings[0])
			fmt.Printf("slice Type %T\n", bookings)
			fmt.Printf("slice length %v\n", len(bookings))

			fmt.Printf("Thank you  %v %v for booking %v tickets. You will receive a confirmation email as %v \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			fmt.Printf("These are all our bookings: %v\n", bookings)

			firstNames := []string{}

			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These first names of bookings are: %v\n", firstNames)
			// if else
			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		}else if userTickets == remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book any other tickets\n", remainingTickets)
			break
			
		}else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
			// fmt.Printf(("We only have %v tickets remaining, so you can't book %v tickets\n"), remainingTickets, userTickets)
			//fmt.Printf("Your input is invalid. Try again.\n")
			// continue
		}

	}

	city := "New York"

	switch city {
	case "New York":
		fmt.Println("New York")
	case "London", "Berlin":
		fmt.Println("London")
	default:
		fmt.Println("Default case")
	}

}
