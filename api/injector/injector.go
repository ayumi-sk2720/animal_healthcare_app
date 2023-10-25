package injector

import (
	"github.com/myantyuWorld/animal_healthcate/domain/repository"
	"github.com/myantyuWorld/animal_healthcate/handler"
	"github.com/myantyuWorld/animal_healthcate/infra"
	"github.com/myantyuWorld/animal_healthcate/usecase"
)

func InjectDB() infra.SqlHandler {
	sqlhandler := infra.NewSqlHandler()
	return *sqlhandler
}

/*
TodoRepository(interface)に実装であるSqlHandler(struct)を渡し生成する。
*/
func InjectPetRepository() repository.PetRepository {
	sqlHandler := InjectDB()
	return infra.NewPetRepository(sqlHandler)
}
func InjectScheduleRepository() repository.ScheduleRepository {
	sqlHandler := InjectDB()
	return infra.NewScheduleRepository(sqlHandler)
}

func InjectPetUsecase() usecase.PetUsecase {
	PetRepo := InjectPetRepository()
	return usecase.NewPetUsecase(PetRepo)
}
func InjectScheduleUsecase() usecase.ScheduleUsecase {
	ScheduleRepo := InjectScheduleRepository()
	return usecase.NewScheduleUsecase(ScheduleRepo)
}

func InjectPetHandler() handler.PetHandler {
	return handler.NewPetHandler(
		InjectPetUsecase(),
		InjectScheduleUsecase(),
	)
}
