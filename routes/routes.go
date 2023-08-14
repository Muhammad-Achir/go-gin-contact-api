package routes

import (
	"go-gin-contact-api/controllers"
	"go-gin-contact-api/models"
	"go-gin-contact-api/repositories"
	"go-gin-contact-api/services"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {
	dsn := "host=localhost user=postgres password=root dbname=go-gin-contact-api port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}

	db.AutoMigrate(&models.Contact{})

	contactRepository := repositories.NewContactRepository(db)
	contactService := services.NewContactService(contactRepository)

	contactController := controllers.NewContactController(contactService)
	r := gin.Default()

	// Grup rute v1
	v1 := r.Group("/v1")
	{
		v1.GET("/get-all-contact", contactController.GetAllContacs)
		v1.GET("/get-contact/:id", contactController.GetContact)
		v1.POST("/add-contact", contactController.CreateContact)
		v1.PUT("/update-contact/:id", contactController.UpdateContact)
		v1.DELETE("/delete-contact/:id", contactController.DeleteContact)
	}

	return r
}
