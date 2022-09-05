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
