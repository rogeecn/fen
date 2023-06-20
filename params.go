package fen

import (
	"github.com/go-playground/validator"
	"golang.org/x/exp/constraints"
)

var validate = validator.New()

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func Validate(in interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	if err := validate.Struct(in); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func Integer[T constraints.Integer](key string, errRet BusError) func(*Ctx) (T, error) {
	return func(ctx *Ctx) (T, error) {
		v, err := ctx.ParamsInt(key)
		if err != nil {
			return T(0), errRet.Format(key).Wrap(err)
		}
		return T(v), nil
	}
}

func String(key string, errRet BusError) func(*Ctx) (string, error) {
	return func(ctx *Ctx) (string, error) {
		return ctx.Params(key), nil
	}
}

func Bind[T any](param T, errRet BusError) func(*Ctx) (T, error) {
	return func(ctx *Ctx) (T, error) {
		if err := ctx.BodyParser(param); err != nil {
			return param, errRet.Wrap(err)
		}

		if err := Validate(param); err != nil {
			return param, errRet.SetData(err)
		}

		return param, nil
	}
}

func Query[T any](param T, errRet BusError) func(*Ctx) (T, error) {
	return func(ctx *Ctx) (T, error) {
		if err := ctx.QueryParser(param); err != nil {
			return param, errRet.Wrap(err)
		}

		if err := Validate(param); err != nil {
			return param, errRet.SetData(err)
		}

		return param, nil
	}
}

func Header[T any](param T, errRet BusError) func(*Ctx) (T, error) {
	return func(ctx *Ctx) (T, error) {
		err := ctx.ReqHeaderParser(&param)
		if err != nil {
			return param, errRet.Wrap(err)
		}

		return param, nil
	}
}
