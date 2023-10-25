package infra

import (
	"github.com/myantyuWorld/animal_healthcate/domain/model"
	"github.com/myantyuWorld/animal_healthcate/domain/repository"
)

/***
* repository --|> infra
* 実現したい振る舞いをrepositoryで定義し、ここで実装していく
***/
type ScheduleRepository struct {
	SqlHandler
}

// Create implements repository.ScheduleRepository.
func (scheduleRepo *ScheduleRepository) Create(schedule *model.Schedule) (*model.Schedule, error) {
	err := scheduleRepo.Conn.Create(schedule).Error
	return schedule, err
}

func NewScheduleRepository(sqlHandler SqlHandler) repository.ScheduleRepository {
	scheduleRepository := ScheduleRepository{sqlHandler}
	return &scheduleRepository
}
