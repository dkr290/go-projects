package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createAt  time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "---",
			createAt:  time.Now(),
		},
	}

}

func (a *Admin) PrintAll() {
	fmt.Println("Email:", a.email)
	fmt.Println("First Name", a.firstName)

}

//we dont need to expose fields because the new function will set them up and not need

func (u *User) OutPutDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createAt)
}

// its like constructor function
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstname, lastname and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createAt:  time.Now(),
	}, nil
}
