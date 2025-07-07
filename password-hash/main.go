package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

type User struct {
	Username     string
	Salt         string
	PasswordHash string
}

var userDB = make(map[string]User)

func generateSalt() (string, error) {
	salt := make([]byte, 16)

	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

func hashPassword(password, salt string) string {
	hasher := sha256.New()
	hasher.Write([]byte(salt + password))
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}

func signup(username, password string) {
	fmt.Printf("🔐 SIGNUP PROCESS for user: %s\n", username)
	fmt.Printf("Password: %s\n", password)
	fmt.Println(strings.Repeat("-", 50))
	salt, err := generateSalt()
	if err != nil {
		fmt.Println("error generating the salt")
		return
	}
	fmt.Printf("1. Generated salt is %s\n", salt)

	hash := hashPassword(password, salt)
	fmt.Printf("2. Password +Salt: %s + %s", password, salt)
	fmt.Printf("3. SHA256 HASH %s\n", hash)

	userDB[username] = User{
		Username:     username,
		Salt:         salt,
		PasswordHash: hash,
	}
	fmt.Println("4. Stored in the db")
	fmt.Println()
}

func login(username, password string) bool {
	fmt.Printf("🔑 LOGIN ATTEMPT for user: %s\n", username)
	fmt.Printf("Provided password: %s\n", password)
	fmt.Println(strings.Repeat("-", 50))

	user, exists := userDB[username]
	if !exists {
		fmt.Println("User not found")
		return false
	}
	fmt.Printf("1. Retrieved from DB - Salt: %s\n", user.Salt)
	fmt.Printf("2. Retrieved from DB - Hash: %s\n", user.PasswordHash)
	providedHash := hashPassword(password, user.Salt)
	fmt.Printf("3. Provided password + stored salt: %s + %s\n", password, user.Salt)
	fmt.Printf("4. Computed hash: %s\n", providedHash)

	// Compare hashes
	if providedHash == user.PasswordHash {
		fmt.Println("5. ✅ MATCH! Login successful")
		return true
	} else {
		fmt.Println("5. ❌ NO MATCH! Login failed")
		return false
	}
}

func showDatabase() {
	fmt.Println("💾 DATABASE CONTENTS:")
	fmt.Println(strings.Repeat("=", 80))
	for username, user := range userDB {
		fmt.Printf("Username: %s\n", username)
		fmt.Printf("Salt:     %s\n", user.Salt)
		fmt.Printf("Hash:     %s\n", user.PasswordHash)
		fmt.Println(strings.Repeat("-", 40))
	}
}

func main() {
	fmt.Println("🎓 PASSWORD HASHING DEMO WITH SHA-256 + SALT")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println()

	users := []struct {
		username string
		password string
	}{
		{"alice", "mySecretPass123"},
		{"bob", "Password123r"},
		{"charlie", "charlieKey789"},
	}

	fmt.Println("STEP 1: USER REGISTRATION")
	fmt.Println()
	for _, user := range users {
		signup(user.username, user.password)
	}

	fmt.Println("STEP 2: Database state")
	showDatabase()
	fmt.Println()

	fmt.Println("STEP 3. Sucessfull LOGIN ATTEMPTS")
	fmt.Println()
	for _, user := range users {
		login(user.username, user.password)
		fmt.Println()
	}

	userwrongPassword := []struct {
		username string
		password string
	}{
		{"alice", "wrongpassword"},
		{"bob", "Password123r"},
		{"charlie", "charlieKey789"},
	}
	fmt.Println("STEP 4. Second Sucessfull or Failed LOGIN ATTEMPTS")
	fmt.Println()
	for _, user := range userwrongPassword {
		login(user.username, user.password)
		fmt.Println()
	}
}
