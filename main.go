package main

import (
	"fmt"
	"time"
)

var conferenceName string = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	greetUser()
	for {
		firstName, lastName, emailAddress, numTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, emailAddress, numTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(numTickets, firstName, lastName, emailAddress)
			go sendTicket(numTickets, firstName, lastName, emailAddress)
			firstNames := getFirstNames()
			fmt.Printf("First names of bookings: %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("No tickets remaining, closing sale.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Entered email does not contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid")
			}

			fmt.Printf("Your input data is invalid, try again \n\n")
		}
	}
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var emailAddress string
	var numTickets uint

	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email: ")
	fmt.Scan(&emailAddress)

	fmt.Println("Please enter number of tickets: ")
	fmt.Scan(&numTickets)
	return firstName, lastName, emailAddress, numTickets
}

func bookTickets(numTickets uint, firstName string, lastName string, emailAddress string) {
	remainingTickets = remainingTickets - numTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           emailAddress,
		numberOfTickets: numTickets,
	}
	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, numTickets, emailAddress)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(numTickets uint, firstName string, lastName string, emailAddress string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", numTickets, firstName, lastName)
	fmt.Println("########")
	fmt.Printf("Sending Ticket: %v\nTo email address %v\n", ticket, emailAddress)
	fmt.Println("########")
}
