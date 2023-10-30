package mock_repository

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/myantyuWorld/animal_healthcate/domain/model"
	"github.com/myantyuWorld/animal_healthcate/usecase"
)

func TestViewForPetId2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected *model.TopInfo
	var err error

	mockView := NewMockPetRepository(ctrl)
	mockView.EXPECT().TopView("2").Return(expected, err)

	petUsecase := usecase.NewPetUsecase(mockView)
	result, err := petUsecase.TopView("2")

	if err != nil {
		t.Error("ペットなし")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Error("等しくない")
	}
}

// TODO: 異常系（petId = 999）とか指定した時に、存在せずアプリケーションロジックエラーのJSONが帰ってくることを確認
