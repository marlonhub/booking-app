package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings []string

const conferenceTickets int = 50

type confAttendings struct {
	FirstName   string `json:"fname"`
	LastName    string `json:"lname"`
	Email       string `json:"email"`
	UserTickets uint   `json:"tickets"`
}

func main() {

	greetUsers()

	for {
		userData := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(userData.FirstName, userData.LastName, userData.Email, userData.UserTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userData.UserTickets, userData.FirstName, userData.LastName, userData.Email)
			go sendTicket(userData.UserTickets, userData.FirstName, userData.LastName, userData.Email)

			// call function print first names
			firstNames := getFirstNames()

			fmt.Printf("Current bookings are: %v\n", firstNames)
			stringify := EncodeJson(userData.FirstName, userData.LastName, userData.Email, userData.UserTickets)
			fmt.Println("User: ", stringify)
			if remainingTickets == 0 {
				fmt.Print("The conference is sold out!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First and/or last name is not valid!")
			}
			if !isValidEmail {
				fmt.Println("Not a valid email address!")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets entered is invalid!")
			}

		}

	}

}

// functions
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend here! ")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
		break

	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= uint(remainingTickets)
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() confAttendings {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Print("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email:")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Println(firstName, lastName, email, userTickets)
	return confAttendings{FirstName: firstName, LastName: lastName, Email: email, UserTickets: uint(userTickets)}
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - uint(userTickets)
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("\n#######################")
	fmt.Printf("Sending ticket: \n %v \nto email address %v\n", ticket, email)
	fmt.Println("#######################")
}

func EncodeJson(firstName string, lastName string, email string, userTickets uint) string {
	attendingUsers := confAttendings{
		FirstName: firstName, LastName: lastName, Email: email, UserTickets: uint(userTickets),
	}

	fmt.Println(" ", attendingUsers)

	finalJson, err := json.MarshalIndent(attendingUsers, "", "    ")
	if err != nil {
		panic(err)
	}

	return string(finalJson[:])
}
