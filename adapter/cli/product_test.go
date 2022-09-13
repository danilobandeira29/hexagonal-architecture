package cli_test

import (
	"fmt"
	"github.com/danilobandeira29/hexagonal-architecture/adapter/cli"
	mock_application "github.com/danilobandeira29/hexagonal-architecture/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	productId := "1234"
	productName := "Product example"
	productPrice := 44.4
	productStatus := "enabled"
	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	serviceMock := mock_application.NewMockProductServiceInterface(ctrl)
	serviceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()
	expected := fmt.Sprint("Product ID 1234 with the name Product example has been created with the price 44.400000 and status enabled.")
	result, err := cli.Run(serviceMock, "create", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expected, result)
	expected = fmt.Sprint("Product Product example has been enabled.")
	result, err = cli.Run(serviceMock, "enable", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expected, result)
	expected = fmt.Sprint("Product Product example has been disabled.")
	result, err = cli.Run(serviceMock, "disable", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expected, result)
	expected = fmt.Sprint("Product ID: 1234\nName: Product example\nPrice: 44.400000\nStatus: enabled")
	result, err = cli.Run(serviceMock, "get", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expected, result)
}
