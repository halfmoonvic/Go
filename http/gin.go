package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	r := gin.Default()

	logger, err := zap.NewProduction()

	if err != nil {
		panic(err)
	}

	r.Use(func(c *gin.Context) {
		fmt.Println("logger")
		s := time.Now()
		c.Next()

		//  log latency, response code, path
		logger.Info("incoming request",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)))

	}, func(c *gin.Context) {
		fmt.Println("id")
		c.Set("requestId", "rand.Int()")
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if rid, exists := c.Get("requestId"); exists {
			h["requestId"] = rid
		}
		c.JSON(200, h)
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	r.Run(":8809") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
