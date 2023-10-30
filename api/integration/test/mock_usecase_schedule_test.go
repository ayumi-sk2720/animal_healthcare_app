package mock_repository

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/myantyuWorld/animal_healthcate/domain/model"
	"github.com/myantyuWorld/animal_healthcate/usecase"
)

func TestCreateForOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected *model.Schedule
	var err error

	mockView := NewMockScheduleRepository(ctrl)
	mockView.EXPECT().Create(expected).Return(expected, err)

	scheduleUsecase := usecase.NewScheduleUsecase(mockView)
	err = scheduleUsecase.Create(expected)

	if err != nil {
		t.Error("スケジュールの追加　エラーです")
	}
}
