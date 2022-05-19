package main

import (
	"example.com/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: " uri=${uri},method=${method}, status=${status}\n",
	}))

	e.POST("/sendfile", controller.Sendfile)
	e.GET("/getfiles", controller.Getfiles)

	e.Start(":8080")
}
