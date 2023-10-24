package usecase

import (
	"github.com/myantyuWorld/animal_healthcate/domain/model"
	"github.com/myantyuWorld/animal_healthcate/domain/repository"
)

type PetUsecase interface {
	TopView(petId string) (topInfo *model.TopInfo, err error)
}

type petUsecase struct {
	petRepo repository.PetRepository
}

// New
func NewPetUsecase(petRepo repository.PetRepository) PetUsecase {
	petUsecase := petUsecase{petRepo: petRepo}
	return &petUsecase
}

// Method
func (usecase *petUsecase) TopView(petId string) (topInfo *model.TopInfo, err error) {
	topInfo, err = usecase.TopView(petId)
	return
}
