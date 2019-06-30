package main

import (
	"awesome-portal-api/internal/service"
	"awesome-portal-api/internal/storage"
	"awesome-portal-api/internal/transport"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
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

	// Migrate db
	storage.MigrateAll(db)

	// Load storage
	studentStorage := storage.NewStudentStorage(db)
	programStorage := storage.NewProgramStorage(db)
	facultyStorage := storage.NewFacultyStorage(db)
	subjectStorage := storage.NewSubjectStorage(db)
	typeSubStorage := storage.NewTypeSubStorage(db)
	enrollStorage := storage.NewEnrollStorage(db)
	scoreStorage := storage.NewScoreStorage(db)

	// Load service
	studentService := service.StudentService{
		StudentStorage: studentStorage,
		ProgramStorage: programStorage,
		FacultyStorage: facultyStorage,
		SubjectStorage: subjectStorage,
		TypeSubStorage: typeSubStorage,
		EnrollStorage:  enrollStorage,
		ScoreStorage:   scoreStorage,
	}
	programService := service.ProgramService{ProgramStorage: programStorage}
	facultyService := service.FacultyService{FacultyStorage: facultyStorage}
	subjectService := service.SubjectService{
		SubjectStorage: subjectStorage,
		ProgramStorage: programStorage,
		FacultyStorage: facultyStorage,
		TypeSubStorage: typeSubStorage,
	}
	typeSubService := service.TypeSubService{TypeSubStorage: typeSubStorage}
	enrollService := service.EnrollService{EnrollStorage: enrollStorage}
	scoreService := service.ScoreService{ScoreStorage: scoreStorage}

	// Load transport
	studentTransport := transport.StudentTransport{StudentService: &studentService}
	programTransport := transport.ProgramTransport{ProgramService: &programService}
	facultyTransport := transport.FacultyTransport{FacultyService: &facultyService}
	subjectTransport := transport.SubjectTransport{SubjectService: &subjectService}
	typeSubTransport := transport.TypeSubTransport{TypeSubService: &typeSubService}
	enrollTransport := transport.EnrollTransport{EnrollService: &enrollService}
	scoreTransport := transport.ScoreTransport{ScoreService: &scoreService}

	// Config gin
	gin.SetMode(os.Getenv("GIN_MODE"))
	gin.DisableConsoleColor()
	route := gin.Default()

	route.GET("/students/:mssv", studentTransport.StudentByMSSV)
	route.POST("/students", studentTransport.Save)
	route.DELETE("/students/:mssv", studentTransport.DeleteByMSSV)
	route.POST("/auth/login", studentTransport.Validate)

	route.POST("/programs", programTransport.Save)
	route.POST("/faculties", facultyTransport.Save)

	route.GET("/subjects/:id", subjectTransport.Subject)
	route.POST("/subjects", subjectTransport.Save)
	route.POST("/type_subs", typeSubTransport.Save)

	route.POST("/enrolls", enrollTransport.Save)
	route.POST("/scores", scoreTransport.Save)

	// Run gin
	if err := route.Run(":" + os.Getenv("PORT")); err != nil {
		log.Println(err)
		log.Fatal("Error running gin")
	}
}
