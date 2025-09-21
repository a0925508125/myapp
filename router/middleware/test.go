package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Test() gin.HandlerFunc {
	log.Printf("Test")
	return func(c *gin.Context) {

	}
}
