package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Cookier(id string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get cookie if it exists
		if _, err := c.Cookie("wordpress_test_cookie"); err == nil {
			c.Next()
		} else {
			// set a cookie
			// -1 ttl for session cookie
			sessionId := uuid.New().String()
			c.SetCookie("wordpress_test_cookie", sessionId, 300, "/", "", false, true)
			c.Next()
		}
		return
	}
}

func setupRouter() *gin.Engine {
	id := uuid.New().String()
	fmt.Println(id)

	router := gin.Default()
	router.Use(Cookier(id))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.GET("/id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": id,
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/pong", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})

	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		// c.String(http.StatusOK, "Hello %s", name)
		message := fmt.Sprintf("Hello, %s", name)
		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})
	return router
}

func main() {
	router := setupRouter()
	// listen and serve on 0.0.0.0:8080
	// use "localhost:8080" on windows
	router.Run(":8080")
}
