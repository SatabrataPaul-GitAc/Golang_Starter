package helper

import "strings"

//Capitalizing the function name, which servers as to explicitly export this function from this package (helper) to be usedin other packages
func ValidateUserInput(firstName string, lastName string, email string, tickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := tickets > 0 && tickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
