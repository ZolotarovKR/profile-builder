package storage

type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "CV with given id doesn't exist"
}

func NewNotFoundError() error {
	return NotFoundError{}
}

type ValidationError string

func (e ValidationError) Error() string {
	return string(e)
}

func NewValidationError(s string) error {
	return ValidationError(s)
}
