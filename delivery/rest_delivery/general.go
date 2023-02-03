package rest_delivery

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ErrorForm struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SuccessForm struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func MakeServerErrorForm() ErrorForm {
	return MakeErrorWithCode("server_error")
}

func MakeErrorWithCode(code string, message ...string) ErrorForm {
	msg := ""
	if len(message) != 0 {
		msg = message[0]
	}
	if msg == "" {
		msg = code
	}

	return ErrorForm{
		Status:  "error",
		Code:    code,
		Message: msg,
	}
}

func MakeSuccessWithData(data interface{}) SuccessForm {
	return SuccessForm{
		Status: "success",
		Data:   data,
	}
}

func MakeSuccess() SuccessForm {
	return SuccessForm{
		Status: "error",
		Data:   "ok",
	}
}

func MakeSuccessWithMessage(message string) SuccessForm {
	return SuccessForm{
		Status: "success",
		Data:   message,
	}
}

type UserError struct {
	error
	code string
}

func NewUserError(code string, err error) UserError {
	return UserError{
		error: err,
		code:  code,
	}
}

func (e *UserError) Code() string {
	return e.code
}

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) == 0 {
		return
	}

	firstError := c.Errors[0]
	switch t := firstError.Err.(type) {
	default:
		log.Error(t)
		c.JSON(500, MakeServerErrorForm())
		break

	case UserError:
		log.Debugf("Custom error [%s]", t.Code())
		c.JSON(500, MakeErrorWithCode(t.Code()))
		break
	}
}
