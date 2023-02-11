package helper

import (
	"strings"
)

func ValidateUserInput(firstName string, lastName string, emailAddress string, numTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(emailAddress, "@")
	isValidTicketNumber := numTickets > 0 && numTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
