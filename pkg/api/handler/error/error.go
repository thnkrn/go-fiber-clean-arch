package error

type ErrorBadRequest struct {
	Err error
}

func (e *ErrorBadRequest) Error() string {
	return e.Err.Error()
}

func NewErrorBadRequest(err error) error {
	return &ErrorBadRequest{Err: err}
}

type ErrorAuthentication struct {
	Err error
}

func (e *ErrorAuthentication) Error() string {
	return e.Err.Error()
}

func NewErrorAuthentication(err error) error {
	return &ErrorAuthentication{Err: err}
}
