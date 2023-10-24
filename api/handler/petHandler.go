package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/myantyuWorld/animal_healthcate/usecase"
)

type PetHandler struct {
	petUsecase usecase.PetUsecase
}

func NewPetHandler(petUsecase usecase.PetUsecase) PetHandler {
	petHandler := PetHandler{petUsecase: petUsecase}
	return petHandler
}

func (handler *PetHandler) TopView() echo.HandlerFunc {
	return func(c echo.Context) error {
		models, err := handler.petUsecase.TopView(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, models)
		}
		return c.JSON(http.StatusOK, models)
	}
}
