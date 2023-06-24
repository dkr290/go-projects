package forms

import (
	"fmt"
	"net/url"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {

	return &Form{
		data,
		errors{},
	}
}

func (f *Form) FormNoValueError(tagId string) {

	f.Errors.Add(tagId, "Filed is empty")

}

func (f *Form) Valid() bool {

	return len(f.Errors) == 0

}

func (f *Form) MinLenght(tagId string, lenght string) {
	if !govalidator.MinStringLength(f.Get(tagId), lenght) {
		f.Errors.Add(tagId, fmt.Sprintf("Fild has minimum lenght of %s", lenght))

	}
}

func (f *Form) ValidateEmail(email string) {
	if !govalidator.IsEmail(email) {
		f.Errors.Add(email, fmt.Sprintf("Not a valid email %s", email))
	}
}
