package main

import (
	"fmt"
	"time"
	"sync"
)

var conferenceName string = "Go Conference"

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

	greetuser(conferenceName, conferenceTickets, remainingTickets)

	//for {

		firstName, lastName, email, userTicket := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTicket, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(remainingTickets, userTicket, firstName, lastName, email)
			
			wg.Add(1)
			go sendTicket(userTicket, firstName, lastName, email)

			//for _, booking := range bookings { //_ is a blank identifier to ignore a variable you don't want to use
			//	var names = strings.Fields(booking)
			//	firstNames = append(firstNames, names[0]) //Fields Splits the stringwith white space as separator
			//}
			//noTicketsRemaining := remainingTickets == 0
			//if noTicketsRemaining {

			firstNames := getFirstNames(bookings)
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("All tickets are sold out. Come back newxt year!")
				//break
			}

		} else {
			if !isValidName {
				fmt.Println("First Name or last name you enteres is too short. Please enter atleast 2 characters.")
			}
			if !isValidEmail {
				fmt.Println("Email is invalid. Please enter a valid email.")
			}
			if !isValidTicketNumber {
				fmt.Println("Invalid Ticket Number. Please enter a valid ticket number.")
			}
			//continue  //Skip the rest of the code and start from the beginning of the loop

			//userName = "Tom"
			//userTicket = 2
			//fmt.Println("User",firstName ,"has booked",userTicket,"tickets.")

			//formatting the print statement
			//fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

		}
		wg.Wait()
	}
//}

func greetuser(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Println("Welcome to the Go Conference\n", conferenceName)
	//fmt.Printf("conferenceTickets is %T, remainingTickets %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your ticket here to attend")
}

func getFirstNames(bookings []UserData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// func validateUserInput(firstName string, lastName string, email string, userTicket uint, remainingTickets uint) (bool, bool, bool) {

// 	isValidName := len(firstName) >= 2 && len(lastName) >= 2
// 	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
// 	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTickets
// 	return isValidName, isValidEmail, isValidTicketNumber

// } //Present in helper.go

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTicket uint
	//ask user for their name

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) //Pointer = Memory Address

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email name: ")
	fmt.Scan(&email)

	fmt.Println("Enter Number of Tickets: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTicket(remainingTickets uint, userTicket uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTicket //bookings[0] = firstName + " " + lastName

	//Create a map for user booking
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTicket,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of booking: %v\n", bookings)

	//fmt.Printf("The whole slice: %v\n", bookings)
	//fmt.Printf("The First value: %v\n", bookings[0])
	//fmt.Printf("Slice Type: %T\n", bookings)
	//fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("Thank You %v %v for booking %v tickets. You will recevie a confirmation email at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for the %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(30 * time.Second)
	var ticket = fmt.Sprintf("%v Tickets for %v %v have been sent to your email %v\n", userTicket, firstName, lastName, email)
	fmt.Printf("####################\n")
	fmt.Printf("SENDING TICKET: %v\n", ticket)
	fmt.Printf("####################\n")
	wg.Done()
}

// In GO Language concurrency is Cheap ans easy to implement with go keyword.
