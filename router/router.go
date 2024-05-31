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
	photoRepository := repositories.NewPhotoRepository(db)

	userUsecase := usecases.NewUserUsecase(userRepository)
	photoUsecase := usecases.NewPhotoUsecase(photoRepository)

	userController := controllers.NewUserController(userUsecase)
	photoController := controllers.NewPhotoController(photoUsecase)

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
		userRoutes.GET("/login", middlewares.AuthMiddleware(), userController.GetLoginInfo)

		userRoutes.PUT("/:userID", middlewares.AuthMiddleware(), userController.Update)
		userRoutes.DELETE("/:userID", middlewares.AuthMiddleware(), userController.Delete)
	}

	photoRoutes := r.Group("/photos")
	{
		photoRoutes.POST("/", middlewares.AuthMiddleware(), photoController.AddPhoto)
		photoRoutes.GET("/:photoID", photoController.GetPhoto)
		photoRoutes.DELETE("/:photoID", middlewares.AuthMiddleware(), photoController.DeletePhoto)
	}

	return r
}
