package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

const totalNumOfTickets = 50

var bookingPlatform = "standard gauge railway"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName    string
	lastName     string
	emailAddress string
	userTickets  uint
}

func main() {

	greetUsers()

	for {

		firstName, lastName, emailAddress, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, emailAddress, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, emailAddress)
			go sendTicket(userTickets, firstName, lastName, emailAddress)

			firstNames := getFirstNames()
			fmt.Printf("The first names of the booking are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("The train is bboked out,please check the next available time")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")

			}
		}

	}
}

func greetUsers() {

	fmt.Printf("Welcome the %v booking application\n", bookingPlatform)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", totalNumOfTickets, remainingTickets)
	fmt.Println("Get your tickets by clicking the link below")
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
	var userTickets uint
	// ask user for their name and other details
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&emailAddress)

	fmt.Println("Enter the number of tickets to purchase: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, emailAddress, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, emailAddress string) {
	remainingTickets = remainingTickets - userTickets

	// create a struct for a user
	var userData = UserData{
		firstName:    firstName,
		lastName:     lastName,
		emailAddress: emailAddress,
		userTickets:  userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets.You will receive a confirmation email at %v\n", firstName, lastName, userTickets, emailAddress)
	fmt.Printf("%v tickets are still remaining for booking\n", remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, emailAddress string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###########")
	fmt.Printf("sending ticket:\n %v to email address %v\n", ticket, emailAddress)
	fmt.Println("###########")
}
