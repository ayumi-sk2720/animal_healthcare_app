package mock_repository

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestCreateForOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

}
