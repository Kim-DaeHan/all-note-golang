package errors

type CustomError struct {
	Message    string
	StatusCode int
	Err        error
}

func (e *CustomError) Error() string {
	return e.Message
}

func (e *CustomError) Status() int {
	return e.StatusCode
}
