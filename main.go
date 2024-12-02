package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
	// "strconv"
)

// struct declaration as map data type struct mean structure
type  UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}
// package level variable must be declared using var, const keyword
const conferenceTickets int = 50
var conferenceName string = "Go Conference"
var remainingTickets uint = 50

// var bookings []string // map[string]string
// var bookings = make([]map[string]string, 0) // map[string]string

var bookings = make([]UserData, 0) // map[string]string

// it will be wait for goroutine task completion
// go is basically concurrency will carate goroutine as green theread its not create OS lavel thread like as queue
// Abstraction of an actual thread
// we can run hundreads of thoudands of millions of goroutines wihout creating OS lavel thread and wihout affecting the performance
// Build-in Functionality for goroutines to talk with one another
//
var wg = sync.WaitGroup{}



func main() {

	// formated print
	greetUsers()

	// reference variable with user input value
	//fixed element in array

	// var bookings = [50]string{}
	// if we dont use loop then wheen complete the main thread task then applicaiton will exit
	// so main therad will not wait for goroutine
	// for this we need to fix for goroutine task completion then main thread will exit
	// package "sync" provided basis synchronization fucntionality
	//
	// for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// call function for booking
			bookTicket(userTickets, firstName, lastName, email)

			// call function for sending ticket using concurrency as goroutine with go keyword
			// A goroutine is a thread that is running in the background
			// A goroutine is a light-weight thread managed by the Go runtime
			// need to be wait for goroutine task completion in main thread with use wg.Wait() and add wg.Add(1)
			wg.Add(1) // Sets the number of goroutines to wait for (increases the counter by the provided number)

			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("These first names of bookings are: %v\n", firstNames)

			// if else
			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				// break
			}
		} else if userTickets == remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book any other tickets\n", remainingTickets)
			// break

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

		wg.Wait() // Block untill the waitGroup counter is 0

	// }

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
		var firstName = booking.firstName
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
	// var userData = map[string]string{
		var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
		// numberOfTickets: strconv.FormatUint(uint64(userTickets), 10), // convert uint to string for store data to userData map
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

func sendTicket(userTickets uint, firstName string, lastName string, email string)  {
	time.Sleep(10 * time.Second) // when run this it will be block the code and wait for 10 seconds so we need to concurrency 
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket: \n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done() // Decrements the waitGroup counter by 1 So this is called by the goroutine to indicate that the goroutine is done/completed

}