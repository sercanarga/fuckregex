package middleware

import (
	"fuckregex/internal"
	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"time"
)

func RateLimitMiddleware(app *gin.Engine) {
	limiter := ratelimit.RateLimiter(ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  1 * time.Second,
		Limit: 100,
	}), &ratelimit.Options{
		ErrorHandler: internal.ErrorHandler,
		KeyFunc:      internal.KeyFunc,
	})

	app.Use(limiter)
}
