package forms

import (
	"net/url"
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
