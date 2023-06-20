package fen

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Ctx struct {
	*fiber.Ctx
}

func WrapCtx(c *fiber.Ctx) *Ctx {
	ctx := &Ctx{}
	ctx.Ctx = c
	return ctx
}

// Deadline returns that there is no deadline (ok==false) when c.Request has no Context.
func (c *Ctx) Deadline() (deadline time.Time, ok bool) {
	return c.Context().Deadline()
}

// Done returns nil (chan which will wait forever) when c.Request has no Context.
func (c *Ctx) Done() <-chan struct{} {
	return c.Context().Done()
}

func (c *Ctx) Err() error {
	return c.Context().Err()
}

func (c *Ctx) Value(key any) any {
	return c.Context().Value(key)
}
