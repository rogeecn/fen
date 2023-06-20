package fen

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

type BusError struct {
	HttpCode int
	ErrCode  int
	Message  string
	Errors   error
	Data     interface{}
}

func NewBusError(httpCode, code int, msg string) BusError {
	return BusError{HttpCode: httpCode, ErrCode: code, Message: msg}
}

func (b BusError) GetHttpCode() int {
	if AlwaysStatusOK {
		return http.StatusOK
	}
	return b.HttpCode
}

func (b BusError) Format(params ...interface{}) BusError {
	if strings.Contains(b.Message, "%") {
		b.Message = fmt.Sprintf(b.Message, params...)
	}
	return b
}

func (b BusError) SetData(err interface{}) BusError {
	b.Data = err
	return b
}

func (b BusError) Wrap(err error) BusError {
	b.Errors = errors.Wrap(err, b.Message)
	return b
}

func (b BusError) Wrapf(err error, args ...interface{}) BusError {
	return b.Format(args...).Wrap(err)
}

func (b BusError) Error() string {
	return b.Message
}

func (b BusError) Stack() string {
	return fmt.Sprintf("%+v", b.Errors)
}

func (b BusError) StackAsList() []string {
	stack := strings.ReplaceAll(b.Stack(), "\t", "        ")
	if stack == "<nil>" {
		return nil
	}
	return strings.Split(stack, "\n")
}

func (b BusError) JSON(ctx *Ctx, errorStack bool) JsonResponse {
	var json JsonResponse
	json = &JSON{}

	if v := ctx.Value(JsonResponseKey); v != nil {
		json = v.(JsonResponse)
	}

	json.SetCode(b.ErrCode).SetMessage(b.Message)
	if errorStack {
		json.SetErrorStack(b.StackAsList())
	}

	return json
}
