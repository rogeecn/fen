package fen

import (
	"github.com/gofiber/fiber/v2"
)

var (
	ErrProc          func(ctx *Ctx, err error) error
	DataProc         func(ctx *Ctx, data interface{}) error
	ErrorConvertProc func(error) BusError
)

func init() {
	ErrProc = defaultErrProc
	DataProc = defaultDataProc
	ErrorConvertProc = defaultErrorConvert
}

func defaultErrProc(ctx *Ctx, err error) error {
	var busErr BusError
	switch err.(type) {
	case BusError:
		busErr = err.(BusError)
	default:
		busErr = ErrorConvertProc(err)
	}

	logger.Error(busErr.Stack())
	// _, _ = gin.DefaultErrorWriter.Write([]byte(busErr.Stack()))
	return ctx.Status(busErr.GetHttpCode()).JSON(busErr.JSON(ctx, DebugMode))
}

func defaultDataProc(ctx *Ctx, data interface{}) error {
	return ctx.JSON(data)
	// return ctx.JSON(JSON{
	// 	Code:    http.StatusOK,
	// 	Message: MessageSuccess,
	// 	Data:    data,
	// })
}

func Func(f func(*Ctx) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := WrapCtx(c)
		err := f(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		return DataProc(ctx, nil)
	}
}

func Func1[P1 any](
	f func(*Ctx, P1) error,
	pf1 func(*Ctx) (P1, error),
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := WrapCtx(c)
		p, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		err = f(ctx, p)
		if err != nil {
			return ErrProc(ctx, err)
		}
		return DataProc(ctx, nil)
	}
}

func Func2[P1 any, P2 any](
	f func(*Ctx, P1, P2) error,
	pf1 func(*Ctx) (P1, error),
	pf2 func(*Ctx) (P2, error),
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := WrapCtx(c)
		p1, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p2, err := pf2(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		err = f(ctx, p1, p2)
		if err != nil {
			return ErrProc(ctx, err)
		}
		return DataProc(ctx, nil)
	}
}

func Func3[P1 any, P2 any, P3 any](
	f func(*Ctx, P1, P2, P3) error,
	pf1 func(*Ctx) (P1, error),
	pf2 func(*Ctx) (P2, error),
	pf3 func(*Ctx) (P3, error),
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := WrapCtx(c)
		p1, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p2, err := pf2(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p3, err := pf3(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		err = f(ctx, p1, p2, p3)
		if err != nil {
			return ErrProc(ctx, err)
		}
		return DataProc(ctx, nil)
	}
}

func Func4[P1 any, P2 any, P3 any, P4 any](
	f func(*Ctx, P1, P2, P3, P4) error,
	pf1 func(*Ctx) (P1, error),
	pf2 func(*Ctx) (P2, error),
	pf3 func(*Ctx) (P3, error),
	pf4 func(*Ctx) (P4, error),
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := WrapCtx(c)
		p1, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p2, err := pf2(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p3, err := pf3(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p4, err := pf4(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		err = f(ctx, p1, p2, p3, p4)
		if err != nil {
			return ErrProc(ctx, err)
		}
		return DataProc(ctx, nil)
	}
}

func Func5[P1 any, P2 any, P3 any, P4 any, P5 any](
	f func(*Ctx, P1, P2, P3, P4, P5) error,
	pf1 func(*Ctx) (P1, error),
	pf2 func(*Ctx) (P2, error),
	pf3 func(*Ctx) (P3, error),
	pf4 func(*Ctx) (P4, error),
	pf5 func(*Ctx) (P5, error),
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := WrapCtx(c)
		p1, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p2, err := pf2(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p3, err := pf3(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p4, err := pf4(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p5, err := pf5(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		err = f(ctx, p1, p2, p3, p4, p5)
		if err != nil {
			return ErrProc(ctx, err)
		}
		return DataProc(ctx, nil)
	}
}

func Func6[P1 any, P2 any, P3 any, P4 any, P5 any, P6 any](
	f func(*Ctx, P1, P2, P3, P4, P5, P6) error,
	pf1 func(*Ctx) (P1, error),
	pf2 func(*Ctx) (P2, error),
	pf3 func(*Ctx) (P3, error),
	pf4 func(*Ctx) (P4, error),
	pf5 func(*Ctx) (P5, error),
	pf6 func(*Ctx) (P6, error),
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := WrapCtx(c)
		p1, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p2, err := pf2(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p3, err := pf3(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p4, err := pf4(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p5, err := pf5(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p6, err := pf6(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		err = f(ctx, p1, p2, p3, p4, p5, p6)
		if err != nil {
			return ErrProc(ctx, err)
		}
		return DataProc(ctx, nil)
	}
}

func DataFunc[T any](f func(*Ctx) (T, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := WrapCtx(c)
		data, err := f(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		return DataProc(ctx, data)
	}
}

func DataFunc1[T any, P1 any](f func(*Ctx, P1) (T, error), pf1 func(*Ctx) (P1, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := WrapCtx(c)
		p, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		data, err := f(ctx, p)
		if err != nil {
			return ErrProc(ctx, err)
		}
		return DataProc(ctx, data)
	}
}

func DataFunc2[T any, P1 any, P2 any](
	f func(*Ctx, P1, P2) (T, error),
	pf1 func(*Ctx) (P1, error),
	pf2 func(*Ctx) (P2, error),
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := WrapCtx(c)
		p1, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p2, err := pf2(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		data, err := f(ctx, p1, p2)
		if err != nil {
			return ErrProc(ctx, err)
		}
		return DataProc(ctx, data)
	}
}

func DataFunc3[T any, P1 any, P2 any, P3 any](
	f func(*Ctx, P1, P2, P3) (T, error),
	pf1 func(*Ctx) (P1, error),
	pf2 func(*Ctx) (P2, error),
	pf3 func(*Ctx) (P3, error),
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := WrapCtx(c)
		p1, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p2, err := pf2(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p3, err := pf3(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		data, err := f(ctx, p1, p2, p3)
		if err != nil {
			return ErrProc(ctx, err)
		}
		return DataProc(ctx, data)
	}
}

func DataFunc4[T any, P1 any, P2 any, P3 any, P4 any](
	f func(*Ctx, P1, P2, P3, P4) (T, error),
	pf1 func(*Ctx) (P1, error),
	pf2 func(*Ctx) (P2, error),
	pf3 func(*Ctx) (P3, error),
	pf4 func(*Ctx) (P4, error),
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := WrapCtx(c)
		p1, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p2, err := pf2(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p3, err := pf3(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p4, err := pf4(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		data, err := f(ctx, p1, p2, p3, p4)
		if err != nil {
			return ErrProc(ctx, err)
		}
		return DataProc(ctx, data)
	}
}

func DataFunc5[T any, P1 any, P2 any, P3 any, P4 any, P5 any](
	f func(*Ctx, P1, P2, P3, P4, P5) (T, error),
	pf1 func(*Ctx) (P1, error),
	pf2 func(*Ctx) (P2, error),
	pf3 func(*Ctx) (P3, error),
	pf4 func(*Ctx) (P4, error),
	pf5 func(*Ctx) (P5, error),
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := WrapCtx(c)
		p1, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p2, err := pf2(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p3, err := pf3(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p4, err := pf4(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p5, err := pf5(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		data, err := f(ctx, p1, p2, p3, p4, p5)
		if err != nil {
			return ErrProc(ctx, err)
		}
		return DataProc(ctx, data)
	}
}

func DataFunc6[T any, P1 any, P2 any, P3 any, P4 any, P5 any, P6 any](
	f func(*Ctx, P1, P2, P3, P4, P5, P6) (T, error),
	pf1 func(*Ctx) (P1, error),
	pf2 func(*Ctx) (P2, error),
	pf3 func(*Ctx) (P3, error),
	pf4 func(*Ctx) (P4, error),
	pf5 func(*Ctx) (P5, error),
	pf6 func(*Ctx) (P6, error),
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := WrapCtx(c)
		p1, err := pf1(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p2, err := pf2(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p3, err := pf3(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p4, err := pf4(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p5, err := pf5(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		p6, err := pf6(ctx)
		if err != nil {
			return ErrProc(ctx, err)
		}
		data, err := f(ctx, p1, p2, p3, p4, p5, p6)
		if err != nil {
			return ErrProc(ctx, err)
		}
		return DataProc(ctx, data)
	}
}
