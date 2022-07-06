package fen

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
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

func Int(key string, errRet BusError) func(*fiber.Ctx) (int, error) {
	return func(ctx *fiber.Ctx) (int, error) {
		v, err := ctx.ParamsInt(key)
		if err != nil {
			return 0, errRet.Format(key).Wrap(err)
		}

		return v, nil
	}
}

func String(key string, errRet BusError) func(*fiber.Ctx) (string, error) {
	return func(ctx *fiber.Ctx) (string, error) {
		return ctx.Params(key), nil
	}
}

func Bind[T any](param T, errRet BusError) func(*fiber.Ctx) (T, error) {
	return func(ctx *fiber.Ctx) (T, error) {
		if err := ctx.BodyParser(param); err != nil {
			return param, errRet.Wrap(err)
		}

		if err := Validate(param); err != nil {
			return param, errRet.Datas(err)
		}

		return param, nil
	}
}

func Query[T any](param T, errRet BusError) func(*fiber.Ctx) (T, error) {
	return func(ctx *fiber.Ctx) (T, error) {
		if err := ctx.QueryParser(param); err != nil {
			return param, errRet.Wrap(err)
		}

		if err := Validate(param); err != nil {
			return param, errRet.Datas(err)
		}

		return param, nil
	}
}
