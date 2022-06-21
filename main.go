package main

//booking-app is the name ofthe go module, which also serves as n import path for packages
import (
	"booking-app/helper"
	"fmt"
	"strings"
)

//All the below variables are package level variables
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

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking appication !!!\n", conferenceName)
	fmt.Println("We have", conferenceTickets, "tickets in total and ", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")
	fmt.Println()
}

func displayFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func bookTickets() {
	remainingTickets = remainingTickets - uint(tickets)
	bookings = append(bookings, firstName+" "+lastName)
}

func main() {
	for {

		greetUsers(conferenceName, conferenceTickets, remainingTickets)

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets to be booked")
		fmt.Scan(&tickets)

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, tickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets()
			firstNames := displayFirstNames(bookings)

			fmt.Printf("Thank You %v %v for booking %v tickets at %v. You will soon receive a confirmation email regarding the same !!!\n", firstName, lastName, tickets, conferenceName)
			fmt.Printf("Remaining tickets for %v is %v\n\n", conferenceName, remainingTickets)
			fmt.Printf("The first names of our bookings are:  %v\n\n", firstNames)

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

	/*
		Example syntax for switch statement

		city := "London"

		switch city {
			case "New York":
				some code to be executed
			case "London":
				some to be executed
			case "Singapore":
				some code to be executed*
			case "Australia", "Berlin":
				some code to be executed when the choice of city is either Australia or Berlin
			default:
				some print statement , for telling the wrong choice
			}
	*/
}
