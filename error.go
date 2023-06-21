package cerrgo

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func IsCustomError(err error) bool {
	_, ok := err.(*Error)
	return ok
}

func GetErrResponse(err error) map[string]interface{} {
	if IsCustomError(err) {
		return map[string]interface{}{
			"code":    err.(*Error).Code,
			"message": err.(*Error).Message,
		}
	}

	return map[string]interface{}{
		"code":    500,
		"message": err.Error(),
	}
}

func NewError(code int, message string) error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
