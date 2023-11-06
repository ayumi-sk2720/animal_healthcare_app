package repository

import "github.com/myantyuWorld/animal_healthcate/domain/model"

type ScheduleRepository interface {
	Create(schedule *model.Schedule) (*model.Schedule, error)

	// 将来的には、CRUD操作を実現する（一旦、研修準備のレビューのため、ここまで）
	// FindAll(petid string) (schedules []*model.Schedule, err error)
}
