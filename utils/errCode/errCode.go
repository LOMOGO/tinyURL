package errCode

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Error struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Details interface{} `json:"details"`
	Data    interface{} `json:"data"`
}

var codes = map[int]string{}

func newError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		log.Fatalf("错误码:%d已存在，请再更换一个", code)
	}
	codes[code] = msg
	return &Error{
		Code:    code,
		Msg:     msg,
		Details: nil,
		Data:    nil,
	}
}

func (e *Error) ResponseJson(c *gin.Context) {
	c.JSON(e.StatusCode(), e)
}

func (e *Error) WithDetails(details interface{}) *Error {
	newE := e
	newE.Details = details
	return newE
}

func (e *Error) WithData(data interface{}) *Error {
	newE := e
	newE.Data = data
	return newE
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
	case InvalidParams.Code:
		return http.StatusBadRequest
	case GenderURLError.Code:
		return http.StatusInternalServerError
	case GenderQRCodeError.Code:
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}
