package internal

import (
	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
)

func KeyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func ErrorHandler(ctx *gin.Context, info ratelimit.Info) {
	ctx.String(429, "Too many requests!")
}
