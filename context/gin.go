package context

import (
	"context"

	"github.com/gin-gonic/gin"
)

// ctxKey defines how we label context in our requests
const ctxKey = "ctx"

// GetReqCtx obtains the context from the http request in gin
func GetReqCtx(c *gin.Context) context.Context {
	var ctx context.Context
	rCtx, ok := c.Get(ctxKey)
	if ok {
		ctx = rCtx.(context.Context)
	} else {
		ctx = context.Background()
	}
	return ctx
}

// SetReqCtx sets the context to gin's context
func SetReqCtx(c *gin.Context, ctx context.Context) {
	c.Set(ctxKey, ctx)
}
