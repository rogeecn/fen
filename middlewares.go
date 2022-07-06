package fen

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrProc  func(ctx *fiber.Ctx, err BusError) error
	DataProc func(ctx *fiber.Ctx, data interface{}) error
)

func init() {
	ErrProc = defaultErrProc
	DataProc = defaultDataProc
}
func defaultErrProc(ctx *fiber.Ctx, err BusError) error {
	// LOG.Error(err.Stack())
	return ctx.Status(err.HttpCode).JSON(err.JSON(DebugMode))
}
func defaultDataProc(ctx *fiber.Ctx, data interface{}) error {
	return ctx.JSON(JSON{
		Code:    http.StatusOK,
		Message: MessageSuccess,
		Data:    data,
	})
}
func Func(f func(*fiber.Ctx) error) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := f(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		return DataProc(ctx, nil)
	}
}
func Func1[P1 any](
	f func(*fiber.Ctx, P1) error,
	pf1 func(*fiber.Ctx) (P1, error),
) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		p, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		err = f(ctx, p)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		return DataProc(ctx, nil)
	}
}
func Func2[P1 any, P2 any](
	f func(*fiber.Ctx, P1, P2) error,
	pf1 func(*fiber.Ctx) (P1, error),
	pf2 func(*fiber.Ctx) (P2, error),
) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		p1, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p2, err := pf2(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		err = f(ctx, p1, p2)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		return DataProc(ctx, nil)
	}
}
func Func3[P1 any, P2 any, P3 any](
	f func(*fiber.Ctx, P1, P2, P3) error,
	pf1 func(*fiber.Ctx) (P1, error),
	pf2 func(*fiber.Ctx) (P2, error),
	pf3 func(*fiber.Ctx) (P3, error),
) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		p1, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p2, err := pf2(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p3, err := pf3(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		err = f(ctx, p1, p2, p3)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		return DataProc(ctx, nil)
	}
}
func Func4[P1 any, P2 any, P3 any, P4 any](
	f func(*fiber.Ctx, P1, P2, P3, P4) error,
	pf1 func(*fiber.Ctx) (P1, error),
	pf2 func(*fiber.Ctx) (P2, error),
	pf3 func(*fiber.Ctx) (P3, error),
	pf4 func(*fiber.Ctx) (P4, error),
) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		p1, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p2, err := pf2(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p3, err := pf3(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p4, err := pf4(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		err = f(ctx, p1, p2, p3, p4)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		return DataProc(ctx, nil)
	}
}
func Func5[P1 any, P2 any, P3 any, P4 any, P5 any](
	f func(*fiber.Ctx, P1, P2, P3, P4, P5) error,
	pf1 func(*fiber.Ctx) (P1, error),
	pf2 func(*fiber.Ctx) (P2, error),
	pf3 func(*fiber.Ctx) (P3, error),
	pf4 func(*fiber.Ctx) (P4, error),
	pf5 func(*fiber.Ctx) (P5, error),
) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		p1, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p2, err := pf2(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p3, err := pf3(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p4, err := pf4(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p5, err := pf5(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		err = f(ctx, p1, p2, p3, p4, p5)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		return DataProc(ctx, nil)
	}
}
func Func6[P1 any, P2 any, P3 any, P4 any, P5 any, P6 any](
	f func(*fiber.Ctx, P1, P2, P3, P4, P5, P6) error,
	pf1 func(*fiber.Ctx) (P1, error),
	pf2 func(*fiber.Ctx) (P2, error),
	pf3 func(*fiber.Ctx) (P3, error),
	pf4 func(*fiber.Ctx) (P4, error),
	pf5 func(*fiber.Ctx) (P5, error),
	pf6 func(*fiber.Ctx) (P6, error),
) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		p1, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p2, err := pf2(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p3, err := pf3(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p4, err := pf4(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p5, err := pf5(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p6, err := pf6(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		err = f(ctx, p1, p2, p3, p4, p5, p6)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		return DataProc(ctx, nil)
	}
}
func DataFunc[T any](f func(*fiber.Ctx) (T, error)) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		data, err := f(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		return DataProc(ctx, data)
	}
}
func DataFunc1[T any, P1 any](f func(*fiber.Ctx, P1) (T, error), pf1 func(*fiber.Ctx) (P1, error)) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		p, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		data, err := f(ctx, p)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		return DataProc(ctx, data)
	}
}
func DataFunc2[T any, P1 any, P2 any](
	f func(*fiber.Ctx, P1, P2) (T, error),
	pf1 func(*fiber.Ctx) (P1, error),
	pf2 func(*fiber.Ctx) (P2, error),
) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		p1, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p2, err := pf2(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		data, err := f(ctx, p1, p2)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		return DataProc(ctx, data)
	}
}
func DataFunc3[T any, P1 any, P2 any, P3 any](
	f func(*fiber.Ctx, P1, P2, P3) (T, error),
	pf1 func(*fiber.Ctx) (P1, error),
	pf2 func(*fiber.Ctx) (P2, error),
	pf3 func(*fiber.Ctx) (P3, error),
) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		p1, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p2, err := pf2(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p3, err := pf3(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		data, err := f(ctx, p1, p2, p3)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		return DataProc(ctx, data)
	}
}
func DataFunc4[T any, P1 any, P2 any, P3 any, P4 any](
	f func(*fiber.Ctx, P1, P2, P3, P4) (T, error),
	pf1 func(*fiber.Ctx) (P1, error),
	pf2 func(*fiber.Ctx) (P2, error),
	pf3 func(*fiber.Ctx) (P3, error),
	pf4 func(*fiber.Ctx) (P4, error),
) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		p1, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p2, err := pf2(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p3, err := pf3(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p4, err := pf4(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		data, err := f(ctx, p1, p2, p3, p4)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		return DataProc(ctx, data)
	}
}
func DataFunc5[T any, P1 any, P2 any, P3 any, P4 any, P5 any](
	f func(*fiber.Ctx, P1, P2, P3, P4, P5) (T, error),
	pf1 func(*fiber.Ctx) (P1, error),
	pf2 func(*fiber.Ctx) (P2, error),
	pf3 func(*fiber.Ctx) (P3, error),
	pf4 func(*fiber.Ctx) (P4, error),
	pf5 func(*fiber.Ctx) (P5, error),
) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		p1, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p2, err := pf2(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p3, err := pf3(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p4, err := pf4(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p5, err := pf5(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		data, err := f(ctx, p1, p2, p3, p4, p5)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		return DataProc(ctx, data)
	}
}
func DataFunc6[T any, P1 any, P2 any, P3 any, P4 any, P5 any, P6 any](
	f func(*fiber.Ctx, P1, P2, P3, P4, P5, P6) (T, error),
	pf1 func(*fiber.Ctx) (P1, error),
	pf2 func(*fiber.Ctx) (P2, error),
	pf3 func(*fiber.Ctx) (P3, error),
	pf4 func(*fiber.Ctx) (P4, error),
	pf5 func(*fiber.Ctx) (P5, error),
	pf6 func(*fiber.Ctx) (P6, error),
) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		p1, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p2, err := pf2(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p3, err := pf3(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p4, err := pf4(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p5, err := pf5(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		p6, err := pf6(ctx)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		data, err := f(ctx, p1, p2, p3, p4, p5, p6)
		if err != nil {
			return ErrProc(ctx, err.(BusError))
		}
		return DataProc(ctx, data)
	}
}
