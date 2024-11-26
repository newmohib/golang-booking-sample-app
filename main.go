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
			firstNames = append(firstNames,  names[0])
		}
		fmt.Printf("These first names of bookings are: %v\n", firstNames)

		// exit
		// if else


	}


}
