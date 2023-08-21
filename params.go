package fen

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
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

func Integer[T constraints.Integer](key string, errRet BusError) func(*fiber.Ctx) (T, error) {
	return func(ctx *fiber.Ctx) (T, error) {
		v, err := ctx.ParamsInt(key)
		if err != nil {
			return T(0), errRet.Format(key).Wrap(err)
		}
		return T(v), nil
	}
}

func String(key string, errRet BusError) func(*fiber.Ctx) (string, error) {
	return func(ctx *fiber.Ctx) (string, error) {
		return ctx.Params(key), nil
	}
}

func Body[T any](errRet BusError) func(*fiber.Ctx) (*T, error) {
	return func(ctx *fiber.Ctx) (*T, error) {
		p := new(T)
		if err := ctx.BodyParser(p); err != nil {
			return nil, errRet.Wrap(err)
		}

		if err := Validate(p); err != nil {
			return nil, errRet.SetData(err)
		}

		return p, nil
	}
}

func Query[T any](errRet BusError) func(*fiber.Ctx) (*T, error) {
	return func(ctx *fiber.Ctx) (*T, error) {
		p := new(T)
		if err := ctx.QueryParser(p); err != nil {
			return nil, errRet.Wrap(err)
		}

		if err := Validate(p); err != nil {
			return nil, errRet.SetData(err)
		}

		return p, nil
	}
}

func Header[T any](errRet BusError) func(*fiber.Ctx) (*T, error) {
	return func(ctx *fiber.Ctx) (*T, error) {
		p := new(T)
		err := ctx.ReqHeaderParser(p)
		if err != nil {
			return nil, errRet.Wrap(err)
		}

		return p, nil
	}
}

const JwtCtxKey = "__jwtClaims{}"

func JwtClaim[T any](errRet BusError) func(*fiber.Ctx) (*T, error) {
	return func(ctx *fiber.Ctx) (*T, error) {
		claim, ok := ctx.Locals(JwtCtxKey).(*T)
		if !ok {
			return nil, errRet
		}
		return claim, nil
	}
}
