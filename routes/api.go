package routes

import (
	"inv-app/controllers"
	"inv-app/middlewares"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(router *gin.Engine) *gin.Engine {

	api := router.Group("/api/v1")
	{
		api.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })
		auth := api.Group("/auth")
		{
			auth.GET("/register", controllers.AuthController().Register)
			auth.POST("/login", controllers.AuthController().Login)
			auth.GET("/logout", controllers.AuthController().Logout)
		}

		protected := api.Group("/")
		protected.Use(middlewares.JWTAuthMiddleware())
		{
			user := protected.Group("/user")
			{
				user.GET("/", controllers.UserController().Index)
				// user.GET("/create", controllers.UserController().Create)
				// user.GET("/show/:id", controllers.UserController().Index)
				// user.POST("/store", controllers.UserController().Store)
				// user.PUT("/update/:id", controllers.UserController().Update)
				// user.PUT("/update/status/:id", controllers.UserController().UpdateStatus)
			}
		}

		// invoice := api.Group("/invoice")
		// {
		// 	invoice.GET("/", controllers.InvoiceController().Index)
		// 	invoice.GET("/create", controllers.InvoiceController().Create)
		// 	invoice.GET("/show/:id", controllers.InvoiceController().Index)
		// 	invoice.POST("/store", controllers.InvoiceController().Store)
		// 	invoice.PUT("/update/:id", controllers.InvoiceController().Update)
		// 	invoice.PUT("/update/status/:id", controllers.InvoiceController().UpdateStatus)
		// }
	}

	return router
}
