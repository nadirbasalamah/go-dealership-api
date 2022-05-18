package model

// ErrorResponse represents response for the validation error
type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}
