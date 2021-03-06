package main

import "fmt"
import "github.com/gin-gonic/gin"
import "net/http"

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
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
    router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

