package main

import (
	"log"
	"os"

	"iot-backend/internal/config"
	"iot-backend/internal/handlers"
	"iot-backend/internal/middleware"
	"iot-backend/internal/repositories"
	"iot-backend/internal/services"

	"github.com/gin-contrib/cors"
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

	manualRepository :=
		repositories.NewManualRepository(db)
	
	spjRepository := repositories.NewSPJRepository(db)
	pembelianAlatRepository :=
	repositories.NewPembelianAlatRepository(db)

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

	manualService := services.NewManualService(manualRepository)
	spjService := services.NewSPJService(
	spjRepository,
)
pembelianAlatService :=
	services.NewPembelianAlatService(
		pembelianAlatRepository,
	)

	// HANDLER
	authHandler :=
		handlers.NewAuthHandler(authService)

	inaprocHandler :=
		handlers.NewInaprocHandler(
			inaprocService,
		)

	manualHandler := handlers.NewManualHandler(manualService)
	inaprocExcelHandler :=handlers.NewInaprocExcelHandler(inaprocService)
	manualExcelHandler := handlers.NewManualExcelHandler(manualService)
	spjHandler := handlers.NewSPJHandler(
	spjService,
	)	
	spjExcelHandler := handlers.NewSPJExcelHandler(
	spjService,
	)
pembelianAlatHandler :=
	handlers.NewPembelianAlatHandler(
		pembelianAlatService,
	)
	pembelianAlatExcelHandler :=
	handlers.NewPembelianAlatExcelHandler(
		pembelianAlatService,
	)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
	AllowOrigins: []string{
		"http://localhost:5173",
		"http://127.0.0.1:5173",
	},
	AllowMethods: []string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
		"OPTIONS",
	},
	AllowHeaders: []string{
		"Origin",
		"Content-Type",
		"Accept",
		"Authorization",
	},
	AllowCredentials: true,
}))

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
	{
	inaproc.POST("", inaprocHandler.Create)
	inaproc.GET("", inaprocHandler.FindAll)
	inaproc.GET("/export",inaprocExcelHandler.Export)
	inaproc.POST("/import",inaprocExcelHandler.Import)
	inaproc.GET("/:id", inaprocHandler.FindByID)
	inaproc.PUT("/:id", inaprocHandler.Update)
	inaproc.DELETE("/:id", inaprocHandler.Delete)
	}

	manualRoutes := api.Group("/pemesanan-manual")
	{
	manualRoutes.POST("", manualHandler.Create)
	manualRoutes.GET("", manualHandler.GetAll)
	manualRoutes.GET("/:id", manualHandler.GetByID)
	manualRoutes.GET("/export", manualExcelHandler.Export)
	manualRoutes.POST("/import", manualExcelHandler.Import)
	manualRoutes.PUT("/:id", manualHandler.Update)
	manualRoutes.DELETE("/:id", manualHandler.Delete)
	}

	spj := api.Group("/spj")
	{
		spj.GET("/export", spjExcelHandler.Export)
		spj.POST("/import", spjExcelHandler.Import)
		spj.POST("", spjHandler.Create)
		spj.GET("", spjHandler.FindAll)
		spj.GET("/:id", spjHandler.FindByID)
		spj.PUT("/:id", spjHandler.Update)
		spj.DELETE("/:id", spjHandler.Delete)
	}

	pembelianAlat := api.Group("/pembelian-alat")
	{
	pembelianAlat.POST("",pembelianAlatHandler.Create)
	pembelianAlat.GET("",pembelianAlatHandler.FindAll)
	pembelianAlat.GET("/export",pembelianAlatExcelHandler.Export)
	pembelianAlat.POST("/import",pembelianAlatExcelHandler.Import)
	pembelianAlat.GET("/:id",pembelianAlatHandler.FindByID)
	pembelianAlat.PUT("/:id",pembelianAlatHandler.Update)
	pembelianAlat.DELETE("/:id",pembelianAlatHandler.Delete)
	}

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