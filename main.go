package main

import (
	"echo-mongo-api/configs"
	"echo-mongo-api/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Connect to database
	configs.ConnectDB()

	// Routes
	routes.UserRoutes(e)

	// CORS Middleware (consider using CORSWithConfig for more options)
	e.Use(middleware.CORS()) // Allows requests from same origin (localhost)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
