package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kasihTakSampai/latih_api/config"
	"github.com/kasihTakSampai/latih_api/controller"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB = config.SetupDatabaseConnection{}
	authController          = controller.newAuthController{}
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	r.Run()
}