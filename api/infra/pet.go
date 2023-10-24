package infra

import (
	"fmt"
	"time"

	"github.com/myantyuWorld/animal_healthcate/domain/model"
	"github.com/myantyuWorld/animal_healthcate/domain/repository"
)

type PetRepository struct {
	SqlHandler
}

func NewPetRepository(sqlHandler SqlHandler) repository.PetRepository {
	petRepository := PetRepository{sqlHandler}
	return &petRepository
}

func (petRepo *PetRepository) TopView(petId string) (topInfo *model.TopInfo, err error) {
	pet := model.Pet{}

	result := petRepo.SqlHandler.Conn.Find(&pet, petId)
	if result.RecordNotFound() {
		fmt.Print("error")
		return
	}

	topInfo.Pet = pet
	// TODO: 将来的には、テーブル結合した結果を返却する --- start ---
	topInfo.DosageSchedules.Today = model.Task{
		Title: "hogehoge",
		Date:  time.Now(),
	}
	topInfo.DosageSchedules.Next = model.Task{
		Title: "fugafuga",
		Date:  time.Now(),
	}
	topInfo.PhysicalCondition.Food = 3
	topInfo.PhysicalCondition.Sweat = 2
	topInfo.PhysicalCondition.Toilet = 1
	topInfo.Memo = model.Task{
		Title: "今月に入って飲む水の量が\n増えた気がする\n次に通院した時に先生に相談する",
		Date:  time.Now(),
	}
	topInfo.Schedules = []model.Schedule{
		{
			PetId:    petId,
			Title:    "トリミング",
			Date:     time.Now(),
			Location: "ペテモ立川店",
		},
		{
			PetId:    petId,
			Title:    "通院",
			Date:     time.Now(),
			Location: "ホゲホゲ病院",
		},
	}
	// TODO: 将来的には、テーブル結合した結果を返却する --- start ---

	return
}
