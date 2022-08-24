package main

import "fmt"

func main() {
	conferenceName := "Go conference" //explicit variable assiging
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var userName string
	var lastName string
	var email string
	var userTickets uint
	//ask user for their name

	fmt.Println("Hello world")
	fmt.Println("Welcome to our ", conferenceName, " booking application")
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of  %v and %v are still remaining", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	fmt.Println(conferenceName)

	//userName = "Tom"

	//var userTickets int
	//userTickets = 2
	//fmt.Printf(" tickets %v\n", userTickets)

	fmt.Printf("conference tickeets is %T\n", conferenceTickets)

	//user input - pointer
	fmt.Print("Enter your name: ")
	fmt.Scan(&userName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter your user ticketes: ")
	fmt.Scan(&userTickets)

	//fmt.Printf("Username %v tickets %v\n", userName, userTickets)
	fmt.Printf("Thank you %v %v for booking %v tickets , You will get confirmation mail at %v ", userName, lastName, userTickets, email)

	remainingTickets = remainingTickets - userTickets
	fmt.Println("Remaining tickets in conference :", remainingTickets)

	//var bookings = [50]string{"p","v","vfv"} //array

	var bookings [50]string
	bookings[0] = userName + " " + lastName
	fmt.Printf("Whole array : %v\n", bookings)
	fmt.Printf("first array value : %v\n", bookings[0])
	fmt.Printf(" array data-type: %T\n", bookings)
	fmt.Printf("Whole array length: %v", len(bookings))
}
