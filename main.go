package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Hello world")
	fmt.Println("Welcome to our ", conferenceName, " booking application")
	fmt.Println("We have total of", conferenceTickets, "and ", remainingTickets, "are still remaiing")
	fmt.Println("Get your tickets here to attend")

	fmt.Println(conferenceName)
}
