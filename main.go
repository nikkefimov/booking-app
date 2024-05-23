package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// *OLD PART* fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// getting data from Users
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames) //list with names, who booked

		//logic for stop booking loop when tickets out
		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			fmt.Println("Sorry, but our conference is booked out. Come back next year.")

		}
	} else {
		if !isValidName {
			fmt.Println("ERROR! First name or last name you entered is too short, please start new registration form.")
		}
		if !isValidEmail {
			fmt.Println("ERROR! Email address you entered doesnt contain @ sign, please start new registration form.")
		}
		if !isValidTicketNumber {
			fmt.Println("ERROR! Number of tickets you entered is invalid, please start new registration form.")
		}
	}
	wg.Wait()

	// *OLD PART*
	/*city := "London"
	switch city {
	case "New York":
		  // execute code
	case "London":
		  // execute code
	case "Paris":
		  // execute code
	case "Amsterdam":
		  // execute code
	case "Copenhagen":
		  // execute code
	    case "Berlin", "Munich":
			//execute code for both

		  // execute code
	    default:
		// execute if none there true for example: fmt.Println("No valid city selected")}*/
}

// great function with info for user
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	// adding second slice for hiding surnames from main slice with bookings, we used there for(each) loop
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// logic for user input information
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

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

// final part booking tickets function with user data, adding new in slice and show list of booking
func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	//struct with user input data
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	// info for user after booking
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets reamining for %v\n", remainingTickets, conferenceName)

}

// *OLD PART*
//fmt.Println(":::::::::::::::::::::::::::::::::::::::<SLICE DATA>:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
//fmt.Printf("The whole slice: %v\n", bookings)
//fmt.Printf("The first value: %v\n", bookings[0])
//fmt.Printf("Slice type: %T\n", bookings)
//fmt.Printf("Slice length: %v\n", len(bookings))
//fmt.Println(":::::::::::::::::::::::::::::::::::::::<SLICE DATA>:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")

// sending tickets info for user
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	fmt.Println("Processing...")
	time.Sleep(10 * time.Second) // fake delay
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("-----------------")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("-----------------")
	fmt.Println("SUCCESSFUL! Tickets have been sent")
	wg.Done()
}
