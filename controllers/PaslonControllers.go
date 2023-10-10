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

type ContorllerPaslon struct {
	PaslonService services.PaslonServices
}

func PaslonController(paslonService services.PaslonServices) *ContorllerPaslon {
	return &ContorllerPaslon{paslonService}
}

func (controllers *ContorllerPaslon) FindPaslons(c echo.Context) error {
	paslons, err := controllers.PaslonService.FindPaslons()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dtos.SuccessResult{
		Code: http.StatusOK,
		Data: convertMultiplePaslonsResponse(paslons)},
	)
}

func (controllers *ContorllerPaslon) GetPaslon(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	paslon, err := controllers.PaslonService.GetPaslon(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dtos.SuccessResult{
		Code: http.StatusOK,
		Data: convertPaslonResponse(paslon)})
}

func (controllers *ContorllerPaslon) CreatePaslon(c echo.Context) error {
	request := new(dtos.PaslonRequest)
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

	paslon := models.Paslon{
		Name:      request.Name,
		Vision:    request.Vision,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	var parties []models.Party
	for _, partyID := range request.Party {
		party := models.Party{ID: partyID}
		parties = append(parties, party)
	}

	paslon.Parties = parties
	response, err := controllers.PaslonService.CreatePaslon(paslon)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtos.SuccessResult{
		Code: http.StatusOK,
		Data: convertPaslonResponse(response),
	})
}

func (controllers *ContorllerPaslon) UpdatePaslon(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	paslon, err := controllers.PaslonService.GetPaslon(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	request := new(dtos.PaslonRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	if request.Name != "" {
		paslon.Name = request.Name
	}

	if request.Vision != "" {
		paslon.Vision = request.Vision
	}

	paslon.UpdatedAt = time.Now()

	response, err := controllers.PaslonService.UpdatePaslon(paslon)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dtos.SuccessResult{
		Code: http.StatusOK,
		Data: convertPaslonResponse(response)})
}

func (controllers *ContorllerPaslon) DeletePaslon(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	paslon, err := controllers.PaslonService.GetPaslon(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := controllers.PaslonService.DeletePaslon(paslon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtos.SuccessResult{
		Code: http.StatusOK,
		Data: data,
	})
}

func convertPaslonResponse(paslon models.Paslon) dtos.PaslonResponse {
	response := dtos.PaslonResponse{
		ID:        paslon.ID,
		Name:      paslon.Name,
		Vision:    paslon.Vision,
		CreatedAt: paslon.CreatedAt,
		UpdatedAt: paslon.UpdatedAt,
	}

	var parties []dtos.PartyResponseToPaslon
	for _, party := range paslon.Parties {
		partyResponse := dtos.PartyResponseToPaslon{
			ID:   party.ID,
			Name: party.Name,
		}
		parties = append(parties, partyResponse)
	}

	response.Parties = parties

	return response
}

func convertMultiplePaslonsResponse(paslons []models.Paslon) []dtos.PaslonResponse {
	var responses []dtos.PaslonResponse

	for _, paslon := range paslons {
		response := convertPaslonResponse(paslon)
		responses = append(responses, response)
	}

	return responses
}
