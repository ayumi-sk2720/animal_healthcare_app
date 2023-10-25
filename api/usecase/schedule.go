package usecase

import (
	"github.com/myantyuWorld/animal_healthcate/domain/model"
	"github.com/myantyuWorld/animal_healthcate/domain/repository"
)

type ScheduleUsecase interface {
	// ユースケースを増やす場合は、ここに宣言し、implements
	Create(scuedule *model.Schedule) (err error)
}

type scheduleUsecase struct {
	scheduleRepo repository.ScheduleRepository
}

// Create implements ScheduleUsecase.
func (usecase *scheduleUsecase) Create(scuedule *model.Schedule) (err error) {
	_, err = usecase.scheduleRepo.Create(scuedule)
	return
}

// New
func NewScheduleUsecase(scheduleRepo repository.ScheduleRepository) ScheduleUsecase {
	scheduleUsecase := scheduleUsecase{scheduleRepo: scheduleRepo}
	return &scheduleUsecase
}
