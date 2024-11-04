package system

type ErrorData interface {
	Error() error
	ErrorText() string
	Code() int
}

func NewErrorResponse(code int, err error) ErrorData {
	return &ErrorResponse{
		code:  code,
		error: err,
	}
}

type ErrorResponse struct {
	code  int
	error error
}

func (e *ErrorResponse) Error() error {
	return e.error
}

func (e *ErrorResponse) ErrorText() string {
	return e.error.Error()
}

func (e *ErrorResponse) Code() int {
	return e.code
}
