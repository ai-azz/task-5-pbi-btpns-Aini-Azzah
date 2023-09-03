package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/asaskevich/govalidator"
)

func main() {
	// Koneksi database 
	dsn := "root:@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrasi models
	db.AutoMigrate(&models.User{}, &models.Photo{})

	// Router
	r := gin.Default()

	// Middleware
	jwtMiddleware := middlewares.NewJWTMiddleware()

	// Controllers
	userController := controllers.NewUserController(db)
	photoController := controllers.NewPhotoController(db)

	// Endpoint User
	r.POST("/users/register", userController.Register) 
	r.POST("/users/login", userController.Login)
	r.PUT("/users/:userId", jwtMiddleware.Authenticate, userController.Update)
	r.DELETE("/users/:userId", jwtMiddleware.Authenticate, userController.Delete)

	// Endpoint Photo
	r.POST("/photos", jwtMiddleware.Authenticate, photoController.Create) 
	r.GET("/photos", photoController.GetAll) 
	r.GET("/photos/:photoId", photoController.GetByID) 
	r.PUT("/photos/:photoId", jwtMiddleware.Authenticate, photoController.Update) 
	r.DELETE("/photos/:photoId", jwtMiddleware.Authenticate, photoController.Delete) 

	// Run port 8080
	r.Run(":8080")
}