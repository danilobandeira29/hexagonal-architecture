package application_test

import (
	"github.com/danilobandeira29/hexagonal-architecture/application"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable_ThrowWhenPriceIsLessOrEqualToZero(t *testing.T) {
	p := application.Product{
		ID:     "",
		Name:   "Product",
		Status: application.DISABLED,
		Price:  0,
	}
	err := p.Enable()
	require.Error(t, err, "the price must be greater than zero to enable the product")
}

func TestProduct_Enable_Success(t *testing.T) {
	p := application.Product{
		ID:     "",
		Name:   "Product",
		Status: application.DISABLED,
		Price:  1,
	}
	err := p.Enable()
	require.Nil(t, err)
}

func TestProduct_Disable_ThrowWhenPriceIsNotZero(t *testing.T) {
	p := application.Product{
		ID:     "",
		Name:   "",
		Status: application.ENABLED,
		Price:  1,
	}
	err := p.Disable()
	require.Error(t, err, "the price must be zero in order to have the product disabled")
}

func TestProduct_Disable_SuccessfullyWhenPriceIsZero(t *testing.T) {
	p := application.Product{
		ID:     "",
		Name:   "",
		Status: application.ENABLED,
		Price:  0,
	}
	err := p.Disable()
	require.Nil(t, err)
	require.Equal(t, application.DISABLED, p.GetStatus())
}
