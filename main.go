package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Arrays and variables pretaining to the name, the amount of tickets avaliable to purchase, as well as the array containing those who have booked.
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings []string

const conferenceTickets int = 50

//ConfAttending Struct which allows the bookings to be organized in a json format.
type confAttendings struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	UserTickets uint   `json:"tickets"`
}

func main() {
	// User's are being greeted by the greetUser function.
	greetUsers()

	// User's are validated according to a strict rule of thumb, first and last name,
	// email, and number of tickets being purchased is validated.
	for {
		// We grab the user's input which we then set to userData from the getUserInput function
		userData := getUserInput()
		// Variables from the validate userInput are being set to the function
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(userData.FirstName, userData.LastName, userData.Email, userData.UserTickets)

		// isValidName, isValidEmail, isValidTicketNumber are now
		if isValidName && isValidEmail && isValidTicketNumber {

			// Booking ticket and szenindg ticket function used to reserve
			// the ticket spot and confirmations.
			bookTicket(userData.UserTickets, userData.FirstName, userData.LastName, userData.Email)
			go sendTicket(userData.UserTickets, userData.FirstName, userData.LastName, userData.Email)

			// call function print first names
			firstNames := getFirstNames()

			// Prints the Users who have bookings planned.
			fmt.Printf("Current bookings are: %v\n", firstNames)

			// USes the encode json to display the informaion according to the json format with the data received
			stringify := EncodeJson(userData.FirstName, userData.LastName, userData.Email, userData.UserTickets)

			// User is presented with an if/else to further see the current bookings.
			var userChoice string
			fmt.Println("Would you like to see who is attending? Y/N")
			fmt.Scan(&userChoice)

			if userChoice == "y" {
				fmt.Println("User: ", stringify)

			} else if userChoice == "n" {

			}
			//An if/else which determines the number of tickets left and if sold out.
			if remainingTickets == 0 {
				fmt.Print("The conference is sold out!")
				break
			}
			// User is then infomred whether their first and or last name, email, or bumber of tickets is valid.
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

// Greeting user function is set. User's are welcome with infomration about the program.
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend here! ")
}

// The getfirstName function grabs the name of the attendees and stores them.
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
		break

	}
	return firstNames
}

// Fucntion to validate whether names, email, and usertickets are set to the rubrics.
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= uint(remainingTickets)
	return isValidName, isValidEmail, isValidTicketNumber
}

// User is prompted to enter their information which is also grabbed from a struct
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

// A booking ticket fucntion is written which keeps a live track of tickets while the application runs.
// While also presentating a cofirmation text.
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - uint(userTickets)
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

// Ticket confirmation/reciept is created based off the user's input
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("\n#######################")
	fmt.Printf("Sending ticket: \n %v \nto email address %v\n", ticket, email)
	fmt.Println("#######################")
}

// A json function which allows the value types to be stored and formatted into json.
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
