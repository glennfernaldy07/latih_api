package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kasihTakSampai/latih_api/config"
	"github.com/kasihTakSampai/latih_api/controller"
	"github.com/kasihTakSampai/latih_api/middleware"
	"github.com/kasihTakSampai/latih_api/repository"
	"github.com/kasihTakSampai/latih_api/service"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
	// authController controller.AuthController = controller.NewAuthController()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("latih-api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("latih-api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	r.Run()
}
