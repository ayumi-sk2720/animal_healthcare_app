package repository

import "github.com/myantyuWorld/animal_healthcate/domain/model"

type Schedule interface {
	FindAll(petid string) (schedules []*model.Schedule, err error)
	Create(schedule *model.Schedule, err error)

	// 将来的には、CRUD操作を実現する（一旦、研修準備のレビューのため、ここまで）
}
