package router

import (
	"btpn-final/controllers"
	"btpn-final/middlewares"
	"btpn-final/repositories"
	"btpn-final/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userRepository := repositories.NewUserRepository(db)

	userUsecase := usecases.NewUserUsecase(userRepository)

	userController := controllers.NewUserController(userUsecase)

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)

		userRoutes.PUT("/:userID", middlewares.AuthMiddleware(), userController.Update)
		userRoutes.DELETE("/:userID", middlewares.AuthMiddleware(), userController.Delete)
	}

	return r
}
