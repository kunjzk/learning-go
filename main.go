package main

import (
	"booking-app-go/helper"
	"fmt"
	"strings"
)

var conferenceName string = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets uint = 50
var bookings = []string{}

func main() {
	greetUser()
	for {
		firstName, lastName, emailAddress, numTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, emailAddress, numTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(numTickets, firstName, lastName, emailAddress)
			firstNames := getFirstNames(bookings)
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

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
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
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, numTickets, emailAddress)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
