package main

import "fmt"

func main() {
	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets int = 50

	fmt.Printf("welcome to my %v booking app\n", conferenceName)
	fmt.Printf("conferenceTickets:%v  - remainingTickets: %v \n", conferenceTickets, remainingTickets)
	fmt.Printf("conferenceticket is %T remainingTickets is %T\n", conferenceTickets, remainingTickets)
	var userName string
	var userTickets int
	//ask user for their name
	fmt.Scan(&userName)

	// userName = "Tahereh"
	userTickets = 2
	fmt.Printf("user %v booked %v tickets\n", userName, userTickets)

	println(remainingTickets)
	println(&remainingTickets) // address of remainingTickets

}
