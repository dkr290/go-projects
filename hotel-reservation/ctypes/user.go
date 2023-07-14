package ctypes

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost = 12
)

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `bjson:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"EncryptedPassword" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encpwd, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encpwd),
	}, nil
}

func (params CreateUserParams) Validate() map[string]string {
	errors := make(map[string]string)

	if !govalidator.IsByteLength(params.FirstName, 2, 200) {
		errors["firstName"] = fmt.Sprintf("firstName length should be at least %d characters", 2)
	}
	if !govalidator.IsByteLength(params.LastName, 2, 200) {
		errors["lastName"] = fmt.Sprintf("lastName length should be at least %d characters", 2)
	}

	if !govalidator.IsByteLength(params.Password, 4, 200) {
		errors["password"] = fmt.Sprintf("password length should be at least %d characters", 4)
	}

	if !govalidator.IsEmail(params.Email) {
		errors["email"] = fmt.Sprintf("invalid Email %v", params.Email)
	}

	return errors

}
