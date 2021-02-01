package errCode

import (
	"log"
	"net/http"
)

type Error struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var codes = map[int]string{}

func newError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		log.Fatalf("错误码:%d已存在，请再更换一个", code)
	}
	codes[code] = msg
	return &Error{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func (e *Error) StatusCode() int {
	switch e.Code {
	case Success.Code:
		return http.StatusOK
	case DatabaseError.Code:
		return http.StatusInternalServerError
	case ServerError.Code:
		return http.StatusInternalServerError
	case NotFound.Code:
		return http.StatusNotFound
	case BindError.Code:
		return http.StatusInternalServerError
	case GenderURLError.Code:
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}
