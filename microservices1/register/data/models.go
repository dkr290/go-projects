package data

import (
	"context"
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const dbTimeout = time.Second * 3

var db *sql.DB

type Models struct {
	User User
}

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) GetAll([]*User, error) {

}

// GetByEmail returns one user by email

// func (u *User) GetByEmail(email string) (*User, error) {

// }

// GetOne returns one user by id

// func (u *User) GetById(id int) (*User, error) {

// }

// Insert inserts a new user into the database, and returns the ID of the newly inserted row

func (u *User) Insert(email, password, fisrtname, lastname string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var user User
	user.Email = email
	user.Password = password
	user.FirstName = fisrtname
	user.LastName = lastname

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return 0, err
	}

	var newID int
	stmt := `insert into users (email, firstname, lastname, password, active, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7) returning id`

	err = db.QueryRowContext(ctx, stmt,
		user.Email,
		user.FirstName,
		user.LastName,
		hashedPassword,
		user.Active,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, err

}
