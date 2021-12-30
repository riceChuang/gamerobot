package model

import (
	"errors"
	"fmt"
	"net/http"
)

type APIException struct {
	Code    int    `json:"-"`
	Message string `json:"error_code,omitempty"`
	Data    string `json:"message,omitempty"`
}

func (e APIException) Error() string {
	return fmt.Sprintf("error_code:%v, message:%v", e.Message, e.Message)
}

var (
	//error for js Error
	ErrInternalServerError = APIException{Code: http.StatusInternalServerError}
	//error not found
	ErrNotFound = APIException{Code: http.StatusNotFound}
	//parse error
	ErrParseFail = APIException{Code: http.StatusBadRequest, Message: "bad_request", Data: "parse fail"}
	//ErrInvalidRequest
	ErrInvalidRequest = APIException{Code: http.StatusBadRequest, Message: "bad_request", Data: "invalid_request"}

	ErrRecordNotFound = errors.New("record not found")
)
