package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference" //Type inference
	// Another syntax for writing the above line that showcases typeinference
	// conferenceName := "Go Conferennce"
	const conferenceTickets = 100
	var remainingTickets uint = 100 //Explicitly specifying unit type, because remaining tickets cannot be negative
	//Users array
	//var users [100]string

	//Bookings slice
	var bookings []string //Alternative syntax-> bookings := []string{}

	//Taking a user's input
	var firstName string
	var lastName string
	var email string
	var tickets uint

	for {
		fmt.Printf("Welcome to %v booking appication !!!\n", conferenceName)
		fmt.Println("We have", conferenceTickets, "tickets in total and ", remainingTickets, "are still available.")
		fmt.Println("Get your tickets here to attend")
		fmt.Println()

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets to be booked")
		fmt.Scan(&tickets)

		isValidName := len(firstName) > 2 && len(lastName) > 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := tickets > 0 && tickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - uint(tickets)
			bookings = append(bookings, firstName+" "+lastName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("Thank You %v %v for booking %v tickets at %v. You will soon receive a confirmation email regarding the same !!!\n", firstName, lastName, tickets, conferenceName)
			fmt.Printf("Remaining tickets for %v is %v\n\n", conferenceName, remainingTickets)
			fmt.Printf("These first names of our bookings: %v \n\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("All tickets for %v are sold out. Come back next year !!!", conferenceName)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("The email you entered is invalid")
			}

			if !isValidTicketNumber {
				fmt.Println("The tickte number entered is invalid")
			}
		}
	}
}
