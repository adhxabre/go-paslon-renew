package routes

import (
	"backend-micro-feature/controllers"
	"backend-micro-feature/databases"
	"backend-micro-feature/services"

	"github.com/labstack/echo/v4"
)

func PartyRoutes(e *echo.Group) {
	service := services.ServiceParty(databases.DB)
	controller := controllers.PartyController(service)

	e.GET("/parties", controller.FindParties)
	e.GET("/party/:id", controller.GetParty)
	e.POST("/party", controller.CreateParty)
	e.PATCH("/party/:id", controller.UpdateParty)
	e.DELETE("/party/:id", controller.DeleteParty)
}
