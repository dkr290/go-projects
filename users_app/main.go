package main

import (
	"fmt"
	"users_app/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	appUser, err := user.New(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println(err.Error())
		return

	}
	// ... do something awesome with that gathered data!

	appUser.OutPutDetails()

	admin := user.NewAdmin("test@example.com", "Password1")
	admin.PrintAll()
	admin.OutPutDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
