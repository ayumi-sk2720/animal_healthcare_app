package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/myantyuWorld/animal_healthcate/domain/model"
	"github.com/myantyuWorld/animal_healthcate/usecase"
)

type PetHandler struct {
	// TODO: こう考えたから、ここに、PetUsecase, ScheduleUsecase（や、他の画面用のusecaseも入れていく）を入れたことを明記
	petUsecase      usecase.PetUsecase
	scheduleUsecase usecase.ScheduleUsecase
}

func NewPetHandler(petUsecase usecase.PetUsecase, scheduleUsecase usecase.ScheduleUsecase) PetHandler {
	petHandler := PetHandler{
		petUsecase:      petUsecase,
		scheduleUsecase: scheduleUsecase,
	}

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

func (handler *PetHandler) PostSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		schedule := new(model.Schedule)
		if error := c.Bind(&schedule); error != nil {
			fmt.Println(error)
			return c.JSON(http.StatusBadRequest, "bad request")
		}
		schedule.PetId = c.Param("id")
		error := handler.scheduleUsecase.Create(schedule)
		return c.JSON(http.StatusOK, error)
	}
}
