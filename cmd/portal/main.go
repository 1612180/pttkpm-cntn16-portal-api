package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mattn/go-colorable"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	fmt.Println(port)

	gin.SetMode(os.Getenv("gin.mode"))
	if gin.Mode() == gin.DebugMode {
		gin.DefaultWriter = colorable.NewColorableStdout()
	} else {
		gin.DisableConsoleColor()
	}

	route := gin.Default()

	route.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hi")
	})

	if err := route.Run(":" + os.Getenv("PORT")); err != nil {
		log.Println(err)
		log.Fatal("Error running gin")
	}
}
