package main

import (
	"backend-micro-feature/databases"
	"backend-micro-feature/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	databases.DatabaseInit()
	databases.RunMigration()

	// routes
	routes.RouteInit(e.Group("/api/v1"))

	e.Logger.Fatal(e.Start("localhost:5000"))
}
