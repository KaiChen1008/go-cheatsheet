package ginutil

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, _ := uuid.NewV6()
		ctx.Header("X-Request-ID", uid.String())
		ctx.Next()
	}
}
