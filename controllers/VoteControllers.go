package controllers

import (
	"backend-micro-feature/dtos"
	"backend-micro-feature/models"
	"backend-micro-feature/services"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ControllerVote struct {
	VoteService services.VoteService
}

func VoteController(VoteService services.VoteService) *ControllerVote {
	return &ControllerVote{VoteService}
}

func (controllers *ControllerVote) FindVotes(c echo.Context) error {
	vote, err := controllers.VoteService.FindVotes()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dtos.SuccessResult{
		Code: http.StatusOK,
		Data: convertMultipleVotesResponse(vote)})
}

func (controllers *ControllerVote) GetVoteByPaslonID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	votes, err := controllers.VoteService.GetVoteByPaslonID(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()},
		)
	}

	return c.JSON(http.StatusOK, dtos.SuccessResult{
		Code: http.StatusOK,
		Data: convertMultipleVotesResponse(votes)},
	)
}

func (controllers *ControllerVote) CreateVote(c echo.Context) error {
	request := new(dtos.VoteRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	data := models.Vote{
		PaslonID:  request.PaslonID,
		VoterName: request.VoterName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	response, err := controllers.VoteService.CreateVote(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtos.SuccessResult{
		Code: http.StatusOK,
		Data: convertVoteResponse(response),
	})
}

func convertVoteResponse(vote models.Vote) dtos.VoteResponse {
	response := dtos.VoteResponse{
		ID:        vote.ID,
		VoterName: vote.VoterName,
		Paslon: dtos.PaslonResponseToVote{
			ID:   vote.Paslon.ID,
			Name: vote.Paslon.Name,
			Visi: vote.Paslon.Visi,
		},
	}

	return response
}

func convertMultipleVotesResponse(votes []models.Vote) []dtos.VoteResponse {
	var responses []dtos.VoteResponse

	for _, vote := range votes {
		response := dtos.VoteResponse{
			ID:        vote.ID,
			VoterName: vote.VoterName,
			Paslon: dtos.PaslonResponseToVote{
				ID:   vote.Paslon.ID,
				Name: vote.Paslon.Name,
				Visi: vote.Paslon.Visi,
			},
		}
		responses = append(responses, response)
	}

	return responses
}
