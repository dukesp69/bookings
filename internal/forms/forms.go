package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form create a custom struct, embembs a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// Valid returns true if there are no errors, otherwhise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks if form fields are in post and not empty.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Has checks if form field is in post and not empty.
func (f *Form) Has(field string) bool {
	x := f.Get(field)
	if x == "" {
		//		f.Errors.Add(field, "This field cannot be blank")
		return false
	}
	return true
}

// MinLenght checks for string minimun lenght
func (f *Form) MinLenght(field string, lenght int) bool {
	x := f.Get(field)
	if len(x) < lenght {
		f.Errors.Add(field, fmt.Sprintf("This field must be at leat %d characters long", lenght))
		return false
	}
	return true
}

// IsEmail checks for valid email address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid Email Address")
	}
}
