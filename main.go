package main

import (
	"github.com/controllers"
	"github.com/gin-gonic/gin"
	"github.com/middleware"
	"github.com/start"
)

func init() {
	start.LoadDotEnv()
	start.ConnectToDB()
	start.SyncDB()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.GET("/login", controllers.LogIn)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}
