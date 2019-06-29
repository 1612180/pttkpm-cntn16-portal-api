package main

import (
	"awesome-portal-api/internal/handlers"
	"awesome-portal-api/internal/repositories"
	"awesome-portal-api/internal/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/mattn/go-colorable"
)

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Println("Error loading .env file")
	}

	// Load database
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}()

	// Create repo
	repos := repositories.NewReposGorm(db)
	studentRepo, accountRepo, programRepo := repos.CreateAll()

	// Create service
	sers := services.NewMyServices(studentRepo, accountRepo, programRepo)
	studentService, programService := sers.CreateAll()

	// Create handler
	hands := handlers.NewMyHandlers(studentService, programService)
	studentHandler, programHandler := hands.CreateAll()

	// Set gin mode
	gin.SetMode(os.Getenv("gin.mode"))
	if gin.Mode() == gin.DebugMode {
		gin.DefaultWriter = colorable.NewColorableStdout()
	} else {
		gin.DisableConsoleColor()
	}

	route := gin.Default()

	api := route.Group("/api")
	{
		studentAPI := api.Group("/students")
		{
			studentAPI.GET("", studentHandler.FetchAll)
			studentAPI.GET("/:id", studentHandler.FindByID)
			studentAPI.POST("", studentHandler.Create)
			studentAPI.DELETE("/:mssv", studentHandler.Delete)
		}

		authAPI := api.Group("/auth")
		{
			authAPI.POST("/login", studentHandler.Login)
		}

		programAPI := api.Group("/programs")
		{
			programAPI.GET("", programHandler.FetchAll)
			programAPI.POST("", programHandler.Create)
		}
	}

	// Run gin
	if err := route.Run(":" + os.Getenv("PORT")); err != nil {
		log.Println(err)
		log.Fatal("Error running gin")
	}
}
