package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// Package level variable cannot be created with these syntax
// conferenceName := "Go Conference"
const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

// var bookings = []string{}

func main() {
	// fmt.Println("Hello World")

	// var bookings = []string{}
	// var bookings [50]string
	// var booking = [50]string{"Rik", "Nana"}

	// fmt.Printf("conference name is %T, conference ticket is %T, remaining ticket is %T\n", conferenceName, conferenceTickets, remainingTickets)

	greetUsers()

	// for {

	firstName, lastName, email, userTickets := getUserInput()
	// ask user for input
	// Pointers in Go
	// fmt.Println(remainingTickets)
	// fmt.Println(&remainingTickets)
	// fmt.Print(firstName, lastName, email, userTickets)
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	// isValidCity := city == "Singapore" || city == "London"
	// isInalidCity := city != "Singapore" || city != "London"
	// above statement can also be written as !isValidCity
	// fmt.Print(isValidName, isValidEmail, isValidTicketNumber)
	if isValidName && isValidEmail && isValidTicketNumber {

		bookTickets(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		firstNames := getFirstNames()

		fmt.Printf("The first names of booking are %v\n", firstNames)
		// fmt.Println(firstNames)
		// fmt.Printf("The first names of our bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Printf("Our conference is booked out. Come back next year.")
			// break
		}
	} else {
		// fmt.Printf("We only have %v tickets remaining.\n", remainingTickets)
		// continue

		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain '@' sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
		// fmt.Printf("Your input is invalid. Please try again.\n")
	}
	wg.Wait()

}

// switch statement
// city := "Singapore"

// switch city {
// 	case "New York":
// 		// execute code for booking New York conference tickets.
// 	case "Singapore", "Hong Kong":
// 		// execute code for booking Singapore and Hong Kong conference tickets.
// 	case "London", "Berlin":
// 		// execute code for booking London and Berlin conference tickets.
// 	case "Mexico City":
// 		// execute code for booking Mexico City conference tickets.
// 	default:
// 		fmt.Println("No valid city selected")
// }
// }

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		// var names = strings.Fields(booking)
		// firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email:")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName

	// Map
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// Struct
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", userData)

	// fmt.Printf("The whole array: %v\n", bookings)
	// fmt.Printf("The first element: %v\n", bookings[0])
	// fmt.Printf("Slice type: %T\n", bookings)
	// fmt.Printf("Slice length %v\n", len(bookings))

	fmt.Printf("Thankyou %v %v for booking %v tickets. You will recieve confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##############")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, email)
	fmt.Println("##############")
	wg.Done()
}
