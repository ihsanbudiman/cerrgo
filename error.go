package cerrgo

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func isCustomError(err error) bool {
	_, ok := err.(*Error)
	return ok
}

func GetErrResponse(err error) map[string]interface{} {
	if isCustomError(err) {
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

func New(message string, statusCode int) error {
	return &Error{
		Code:    statusCode,
		Message: message,
	}
}

func getErrCode(err error) int {
	if isCustomError(err) {
		return err.(*Error).Code
	}

	return 500
}

func SendResponse(w http.ResponseWriter, err error) {

	errData := GetErrResponse(err)
	jsonData, err := json.Marshal(errData)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"code":500,"message":"Internal Server Error"}`))
		return
	}

	// check if w.header is not set
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "application/json")
	}

	w.WriteHeader(getErrCode(err))
	w.Write(jsonData)
}
