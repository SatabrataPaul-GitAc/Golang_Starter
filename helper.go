package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, tickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := tickets > 0 && tickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
