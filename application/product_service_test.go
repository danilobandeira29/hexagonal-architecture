package application_test

import (
	"errors"
	"github.com/danilobandeira29/hexagonal-architecture/application"
	mock_application "github.com/danilobandeira29/hexagonal-architecture/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductService_Get_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{
		Persistence: persistence,
	}
	result, err := service.Get("1234")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Get_ErrWhenPersistenceReturnErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(nil, errors.New("error from persistence")).AnyTimes()
	service := application.ProductService{
		Persistence: persistence,
	}
	result, err := service.Get("1234")
	require.Nil(t, result)
	require.Error(t, err)
}
