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
