package routes

import (
	"backend-micro-feature/controllers"
	"backend-micro-feature/databases"
	"backend-micro-feature/services"

	"github.com/labstack/echo/v4"
)

func VoteRoutes(e *echo.Group) {
	service := services.ServiceVote(databases.DB)
	controller := controllers.VoteController(service)

	e.GET("/votes", controller.FindVotes)
	e.GET("/vote/:id", controller.GetVoteByPaslonID)
	e.POST("/vote", controller.CreateVote)
}
