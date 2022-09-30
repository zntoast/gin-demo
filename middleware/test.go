package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HelloMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("hello world")
	}
}
