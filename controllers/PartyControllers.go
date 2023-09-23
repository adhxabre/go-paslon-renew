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

type ControllerParty struct {
	PartyService services.PartyServices
}

func PartyController(partyService services.PartyServices) *ControllerParty {
	return &ControllerParty{partyService}
}

func (controllers *ControllerParty) FindParties(c echo.Context) error {
	partai, err := controllers.PartyService.FindParties()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dtos.SuccessResult{
		Code: http.StatusOK,
		Data: convertMultiplePartyResponses(partai)})
}

func (controllers *ControllerParty) GetParty(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	Party, err := controllers.PartyService.GetParty(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()},
		)
	}

	return c.JSON(http.StatusOK, dtos.SuccessResult{
		Code: http.StatusOK,
		Data: convertPartyResponse(Party)},
	)
}

func (controllers *ControllerParty) CreateParty(c echo.Context) error {
	request := new(dtos.PartyRequest)
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

	data := models.Party{
		Name:      request.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	response, err := controllers.PartyService.CreateParty(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtos.SuccessResult{
		Code: http.StatusOK,
		Data: convertPartyResponse(response),
	})
}

func (controllers *ControllerParty) UpdateParty(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	Party, err := controllers.PartyService.GetParty(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	request := new(dtos.PartyRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	if request.Name != "" {
		Party.Name = request.Name
	}

	Party.UpdatedAt = time.Now()

	response, err := controllers.PartyService.UpdateParty(Party)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dtos.SuccessResult{
		Code: http.StatusOK,
		Data: convertPartyResponse(response)})
}

func (controllers *ControllerParty) DeleteParty(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	Party, err := controllers.PartyService.GetParty(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := controllers.PartyService.DeleteParty(Party)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtos.SuccessResult{
		Code: http.StatusOK,
		Data: convertPartyResponse(data),
	})
}

func convertPartyResponse(u models.Party) dtos.PartyResponse {
	return dtos.PartyResponse{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func convertMultiplePartyResponses(parties []models.Party) []dtos.PartyResponse {
	partyResponses := make([]dtos.PartyResponse, len(parties))

	for i, party := range parties {
		partyResponses[i] = dtos.PartyResponse{
			ID:        party.ID,
			Name:      party.Name,
			CreatedAt: party.CreatedAt,
			UpdatedAt: party.UpdatedAt,
		}
	}

	return partyResponses
}
