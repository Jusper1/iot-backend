package main

import (
	"log"
	"os"

	"iot-backend/internal/config"
	"iot-backend/internal/handlers"
	"iot-backend/internal/middleware"
	"iot-backend/internal/repositories"
	"iot-backend/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {

	// DATABASE
	db := config.ConnectDatabase()

	// REPOSITORY
	adminRepository :=
		repositories.NewAdminRepository(db)

	passwordResetRepository :=
		repositories.NewPasswordResetRepository(db)

	inaprocRepository :=
		repositories.NewInaprocRepository(db)

	// SERVICE
	authService :=
		services.NewAuthService(
			adminRepository,
			passwordResetRepository,
		)

	inaprocService :=
		services.NewInaprocService(
			inaprocRepository,
		)

	// HANDLER
	authHandler :=
		handlers.NewAuthHandler(authService)

	inaprocHandler :=
		handlers.NewInaprocHandler(
			inaprocService,
		)



	router := gin.Default()

	// PUBLIC
	auth := router.Group("/auth")
	{
		auth.POST("/login",authHandler.Login)
		auth.POST("/forgot-password",authHandler.ForgotPassword)
		auth.POST("/reset-password",authHandler.ResetPassword)
	}


	api := router.Group("/")

	api.Use(middleware.JWTAuth())

	inaproc := api.Group("/inaproc")

	inaproc.POST("", inaprocHandler.Create)
	inaproc.GET("", inaprocHandler.FindAll)
	inaproc.GET("/:id", inaprocHandler.FindByID)
	inaproc.PUT("/:id", inaprocHandler.Update)
	inaproc.DELETE("/:id", inaprocHandler.Delete)

	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8080"
	}

	log.Println(
		"Server berjalan di http://localhost:" + port,
	)

	err := router.Run(":" + port)

	if err != nil {
		log.Fatal(err)
	}
}