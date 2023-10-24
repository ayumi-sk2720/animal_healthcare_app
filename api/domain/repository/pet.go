package repository

import "github.com/myantyuWorld/animal_healthcate/domain/model"

type PetRepository interface {
	TopView(petId string) (pet *model.TopInfo, err error)

	// 将来的には、CRUD操作を実現する（一旦、研修準備のレビューのため、ここまで）
}
