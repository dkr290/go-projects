package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	pw := "Password123"

	hPassword, err := bcrypt.GenerateFromPassword([]byte(pw), 12)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(hPassword))
	fmt.Println(time.Now())
}
