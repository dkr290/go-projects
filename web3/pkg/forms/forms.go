package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) HasRequired(tagIDs ...string) {

	for _, tagID := range tagIDs {
		value := f.Get(tagID)

		if strings.TrimSpace(value) == "" {
			f.Errors.Add(tagID, "This field can't be blank")
		}
	}
}

func (f *Form) HasValue(tagID string, r *http.Request) bool {

	x := r.Form.Get(tagID)

	return x != ""

}

func (f *Form) MinLenght(tagID string, lenght int, r *http.Request) bool {

	x := r.Form.Get(tagID)
	if len(x) < lenght {
		f.Errors.Add(tagID, fmt.Sprintf("This filed have to be %d characters long or more", lenght))
		return false
	}

	return true

}

func (f *Form) IsEmail(tagID string) {
	if !govalidator.IsEmail(f.Get(tagID)) {
		f.Errors.Add(tagID, "Ivalid email")
	}
}

func (f *Form) Valid() bool {

	return len(f.Errors) == 0
}
