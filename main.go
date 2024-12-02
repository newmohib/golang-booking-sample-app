package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

// package level variable must be declared using var, const keyword
const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50

// var bookings []string // map[string]string
var bookings = make([]map[string]string, 0) // map[string]string

func main() {

	// formated print
	greetUsers()

	// reference variable with user input value
	//fixed element in array

	// var bookings = [50]string{}

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// call function for booking
			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("These first names of bookings are: %v\n", firstNames)

			// if else
			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else if userTickets == remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book any other tickets\n", remainingTickets)
			break

		} else {
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

func greetUsers() {

	fmt.Printf("Welcome to %v booking applicaiton \n", conferenceName)
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking applicaiton \n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	//
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		var firstName = booking["firstName"]
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for user
	var userData = map[string]string{
		"firstName":       firstName,
		"lastName":        lastName,
		"email":           email,
		"numberOfTickets": strconv.FormatUint(uint64(userTickets), 10), // convert uint to string for store data to userData map
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("The whole slice %v\n", bookings)
	fmt.Printf("The first slice %v\n", bookings[0])
	fmt.Printf("slice Type %T\n", bookings)
	fmt.Printf("slice length %v\n", len(bookings))

	fmt.Printf("Thank you  %v %v for booking %v tickets. You will receive a confirmation email as %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Printf("These are all our bookings: %v\n", bookings)

}
