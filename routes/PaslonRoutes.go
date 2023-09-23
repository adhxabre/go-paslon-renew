package routes

import (
	"backend-micro-feature/controllers"
	"backend-micro-feature/databases"
	"backend-micro-feature/services"

	"github.com/labstack/echo/v4"
)

func PaslonRoutes(e *echo.Group) {
	service := services.ServicePaslon(databases.DB)
	controller := controllers.PaslonController(service)

	e.GET("/paslons", controller.FindPaslons)
	e.GET("/paslon/:id", controller.GetPaslon)
	e.POST("/paslon", controller.CreatePaslon)
	e.PATCH("/paslon/:id", controller.UpdatePaslon)
	e.DELETE("/paslon/:id", controller.DeletePaslon)
}
