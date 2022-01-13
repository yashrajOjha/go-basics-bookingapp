package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Annual Conference"

//building a ticked application
var remainingtickets uint = 50

//arrays
// var bookings [50]string
//slice - size unknown
// var bookings = make([]map[string]string, 0) //creating empty slice of maps, of initial size of 0, which will increase when we add new elements
//instead of using a map for booking we make it a struct
var bookings = make([]user, 0)

//creating a struct
type user struct {
	fname           string
	lname           string
	emailid         string
	numberoftickets uint
}

var wg = sync.WaitGroup{}

func main() {
	// var userint int
	// fmt.Scan(&userint) // asking for input, we will initally not be able enter anyting
	// fmt.Println(&userint)
	// fmt.Println("Value entered by user", userint)

	//for loop
	// for remainingtickets > 0 && len(bookings) <= 50 {

	greetUsers(conferenceName, remainingtickets)
	firstname, lastname, email, ticketcount := userInput()

	//validation of input function
	namevalid, emailvalid, usertickets := helper.UserInputValidation(firstname, lastname, email, ticketcount, remainingtickets)
	//we put the userInputValidation function  to a different file because then we can use it as a package, to make it run
	//we use the command : go run userinput.go helper.go
	if namevalid && emailvalid && usertickets {
		fmt.Printf("\nThank you %v for boooking %v tickets with us.\n", firstname, ticketcount)

		remainingtickets = remainingtickets - ticketcount //updating the value of the remaining tickets
		fmt.Printf("\nThe remaining number of tickets are: %v", remainingtickets)

		//map for User (key -value) pairs
		//we cannot mix different datatypes in a map
		var userData = user{
			fname:           firstname,
			lname:           lastname,
			emailid:         email,
			numberoftickets: ticketcount,
		}
		// userData["firstName"] = firstname
		// userData["lastName"] = lastname
		// userData["email"] = email
		// userData["numberTickets"] = strconv.FormatUint(uint64(ticketcount), 10) //converting uint to string
		//appending a value to the slice
		bookings = append(bookings, userData)

		fmt.Printf("\nList of bookings %v\n", bookings)

		// fmt.Printf("\nThe whole bookings array: %v\n", bookings)
		// fmt.Printf("The length of the array: %v", len(bookings))

		// fmt.Printf("\nThe whole bookings slice: %v\n", emails)
		// fmt.Printf("The length of the slice: %v", len(emails))

		firstNames := getFirstNames()
		fmt.Printf("\nThe first name of bookings are %v", firstNames)

		wg.Add(1) //based on the number of functions
		//to make an application concurrent just add go infront of it
		go sendTicket(ticketcount, firstname, lastname, email)

		if remainingtickets == 0 {
			fmt.Println("Ticket sold out!")
			// break
		}
		// else if ticketcount == remainingtickets can be used if required
	} else { //just to make sure the user asks for correct number of tickets
		// fmt.Printf("\nWe only have %v tickets remaining. You cant book %v tickets.", remainingtickets, ticketcount)
		//using continue so user can book the correct number of tickets

		if !namevalid {
			fmt.Println("First Name or Last Name you entered is too short (length <2)")
		}
		if !emailvalid {
			fmt.Println("Email entered is invalid.")
		}
		if !usertickets {
			fmt.Println("Number of tickets entered is invalid.")
		}
		// continue
	}
	// }
	/*Switch Statements
	city:="London"
	switch city {
	case "New York":
		//writing some piece of code for new york
	case "London":
		//writing some piece of code for london
	case "Honkong":
		//writing some piece of code for honkong
	case "Berlin","Delhi":
		//same code for Berlin and Delhi
	default:
		//when none of the above cities has been entered
	}
	*/
	wg.Wait() //waits untill wg.Add() gets done
}
func greetUsers(conferenceName string, remainingtickets uint) {
	fmt.Printf("\n\n\t--Hello and Welcome to the %v Ticket Booking System--", conferenceName)
	fmt.Printf("\n\t\t--Total number of reamaining tickets are %v--", remainingtickets)
}
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// firstNames = append(firstNames, booking["firstName"]) //this is to access the key of a map
		firstNames = append(firstNames, booking.fname)
	}
	return firstNames
}

func userInput() (string, string, string, uint) {
	var firstname string
	fmt.Print("\nEnter your first name:")
	fmt.Scan(&firstname)

	var lastname string
	fmt.Print("\nEnter your last name:")
	fmt.Scan(&lastname)

	var email string
	fmt.Print("\nEnter your email:")
	fmt.Scan(&email)

	var ticketcount uint //only positive number of tickets can be booked, and so we use unsigned int
	fmt.Print("\nEnter number of tickets:")
	fmt.Scan(&ticketcount)

	return firstname, lastname, email, ticketcount
}
func sendTicket(ticketcount uint, firstname string, lastname string, email string) {

	//sending a ticket may take some time, so we need to simulate the delay time
	time.Sleep(10 * time.Second)
	// fmt.Printf("%v tickets %v %v", ticketcount, firstname, lastname)
	var ticket = fmt.Sprintf("%v tickets %v %v", ticketcount, firstname, lastname)
	fmt.Println("\n#################")
	fmt.Printf("\nSending ticket: %v to email: %v", ticket, email)
	fmt.Println("\n#################")
	wg.Done() //removes thread from waiting list
}
