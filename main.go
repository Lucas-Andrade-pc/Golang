package main

import (
	"fmt"
	"strings"
)

var remainingTickets int = 50

func welcome(conferenceName string, conferenceTicket int, remainingTickets int) {
	fmt.Printf("type variable conferenceName, conferenceTicket, remainingTickets ->  %T, %T, %T \n", conferenceName, conferenceTicket, remainingTickets)
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v \n", conferenceTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func validated(firstName string, lastName string, userTickets int, email string) (bool, bool, bool) {
	isValidateName := len(firstName) >= 2 && len(lastName) >= 2
	isValidateEmail := strings.Contains(email, "@")
	isValidateTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidateName, isValidateEmail, isValidateTickets
}

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTicket int = 50

	welcome(conferenceName, conferenceTicket, remainingTickets)

	bookings := []string{}

	for {
		var firstName string
		var lastName string
		var userTickets int
		var email string
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName) //referencia para local na memoria onde vai ser armazenada
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email:")
		fmt.Scan(&email)
		fmt.Println("Enter number tickets:")
		fmt.Scan(&userTickets) //referencia para local na memoria onde vai ser armazenada

		funcValidatName, funcValidatLastEmail, funcValidatTicket := validated(firstName, lastName, userTickets, email)

		if funcValidatName && funcValidatLastEmail && funcValidatTicket {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName)

			fmt.Printf("User %v booked %v tickect \n", firstName, userTickets)
			fmt.Printf("Total tickets avaliable %v \n", remainingTickets)
			first := []string{}
			for _, bobooking := range bookings {
				var names = strings.Fields(bobooking)
				first = append(first, names[0])
			}

			noTickects := remainingTickets == 0
			if noTickects {
				fmt.Printf("Sould out")
				break
			}

			fmt.Printf("List user :%v\n", first)
		} else {
			if !funcValidatName {
				fmt.Printf("Your input is invalid -> %v %v\n", firstName, lastName)
			}
			if !funcValidatLastEmail {
				fmt.Printf("Your email is invalid -> %v\n", email)
			}
			if !funcValidatTicket {
				fmt.Printf("Your Tickets is invalid -> %v\n", userTickets)
			}
		}
	}
}
