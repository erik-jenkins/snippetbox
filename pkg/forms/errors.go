package forms

// errors type will hold validation errors
// the key is the name of the form field
type errors map[string][]string

// Add an error to the map
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get the first error for a given field.
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
