package router

import (
	"Go-Service/src/main/application/usecase"
	domainLogger "Go-Service/src/main/domain/interface/logger"
	"Go-Service/src/main/infrastructure/controller"
	"Go-Service/src/main/infrastructure/middleware"
	"Go-Service/src/main/infrastructure/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(r *gin.Engine, db *mongo.Database, logger domainLogger.Logger) {
	// Create API group
	api := r.Group("/api")
	r.Use(middleware.TraceIDMiddleware())

	// Setup controllers
	memeCoinRepository := repository.NewMemeCoinRepository(db, logger)
	memeCoinUsecase := usecase.NewMemeCoinUsecase(logger, memeCoinRepository)

	memeCoinController := controller.NewMemeCoinController(memeCoinUsecase)

	// User routes
	memeCoinRoutes := api.Group("/meme-coin")
	{
		memeCoinRoutes.POST("/", memeCoinController.Create)
		memeCoinRoutes.GET("/", memeCoinController.GetByID)
		memeCoinRoutes.PUT("/", memeCoinController.Update)
		memeCoinRoutes.DELETE("/", memeCoinController.Delete)
		memeCoinRoutes.POST("/poke", memeCoinController.Poke)
	}
}
