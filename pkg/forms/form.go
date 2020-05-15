package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

// Form anonymously holds the form data as well as any validation errors
type Form struct {
	url.Values
	Errors errors
}

// New creates a new Form object with the provided form data
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string]string{}),
	}
}

// Required ensures that the given fields are not empty
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MaxLength ensures the given field is at most `d` runes long
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("The field is too long (maximum is %d characters)", d))
	}
}

// PermittedValues checks that the given field is one of a provided list of options
func (f *Form) PermittedValues(field string, options ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	for _, option := range options {
		if value == option {
			return
		}
	}
	f.Errors.Add(field, "This field is invalid")
}

// Valid returns true if there are no errors in the form, false otherwise
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
