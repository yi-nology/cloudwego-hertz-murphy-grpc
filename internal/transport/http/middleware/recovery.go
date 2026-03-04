package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"{{.module_name}}/pkg/response"
)

func Recovery() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		defer func() {
			if err := recover(); err != nil {
				response.InternalError(c, "Internal server error")
				c.Abort()
			}
		}()
		c.Next(ctx)
	}
}
