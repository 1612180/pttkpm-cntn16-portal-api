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
	reposInterface := repositories.NewReposGorm(db)
	studentRepo,
		accountRepo,
		programRepo,
		facultyRepo,
		subjectRepo,
		subjectTypeRepo,
		studentSubjectRepo := reposInterface.CreateAll()

	// Create service
	servicesInterface := services.NewMyServices(
		studentRepo,
		accountRepo,
		programRepo,
		facultyRepo,
		subjectRepo,
		subjectTypeRepo,
		studentSubjectRepo,
	)
	studentService,
		programService,
		facultyService,
		subjectService,
		subjectTypeService,
		studentSubjectService := servicesInterface.CreateAll()

	// Create handler
	handlersInterface := handlers.NewMyHandlers(
		studentService,
		programService,
		facultyService,
		subjectService,
		subjectTypeService,
		studentSubjectService,
	)
	studentHandler,
		programHandler,
		facultyHandler,
		subjectHandler,
		subjectTypeHandler,
		studentSubjectHandler := handlersInterface.CreateAll()

	// Set gin mode
	gin.SetMode(os.Getenv("GIN_MODE"))
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
			studentAPI.DELETE("/:mssv", studentHandler.DeleteByMSSV)
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

		facultyAPI := api.Group("/faculties")
		{
			facultyAPI.GET("", facultyHandler.FetchAll)
			facultyAPI.POST("", facultyHandler.Create)
		}

		subjectAPI := api.Group("/subjects")
		{
			subjectAPI.GET("", subjectHandler.FetchAll)
			subjectAPI.GET("/:id", subjectHandler.FindByID)
			subjectAPI.POST("", subjectHandler.Create)
			subjectAPI.DELETE("/:id", subjectHandler.DeleteByID)
		}

		subjectTypeAPI := api.Group("/subject_types")
		{
			subjectTypeAPI.GET("", subjectTypeHandler.FetchAll)
			subjectTypeAPI.POST("", subjectTypeHandler.Create)
		}

		studentSubjectAPI := api.Group("/student_subjects")
		{
			studentSubjectAPI.GET("", studentSubjectHandler.FetchAll)
			studentSubjectAPI.GET("/students/:student_id/subjects/:subject_id",
				studentSubjectHandler.FindByID)
			studentSubjectAPI.POST("", studentSubjectHandler.Create)
		}
	}

	// Run gin
	if err := route.Run(":" + os.Getenv("PORT")); err != nil {
		log.Println(err)
		log.Fatal("Error running gin")
	}
}
