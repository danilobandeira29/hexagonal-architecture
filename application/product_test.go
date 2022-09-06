package application_test

import (
	"github.com/danilobandeira29/hexagonal-architecture/application"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable_ErrorWhenPriceIsLessOrEqualToZero(t *testing.T) {
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

func TestProduct_Disable_ErrorWhenPriceIsNotZero(t *testing.T) {
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

func TestProduct_IsValid_ErrorWhenStatusIsNotEnabledOrDisabled(t *testing.T) {
	p := application.Product{
		ID:     uuid.New().String(),
		Name:   "Product",
		Status: "invalid status",
		Price:  1,
	}
	_, err := p.IsValid()
	require.Error(t, err, "the status must be enabled or disabled")
}

func TestProduct_IsValid_ErrorWhenPriceIsLessThanZero(t *testing.T) {
	p := application.Product{
		ID:     uuid.New().String(),
		Name:   "Product",
		Status: application.ENABLED,
		Price:  -1,
	}
	_, err := p.IsValid()
	require.Equal(t, "the price must be greater or equal to zero", err.Error())
}

func TestProduct_IsValid_Success(t *testing.T) {
	p := application.Product{
		ID:     uuid.New().String(),
		Name:   "Product",
		Status: application.ENABLED,
		Price:  0,
	}
	_, err := p.IsValid()
	require.Nil(t, err)
}

func TestProduct_IsValid_ErrWhenIDIsNotUuid(t *testing.T) {
	p := application.Product{
		ID:     "not uuid",
		Name:   "Product",
		Status: application.ENABLED,
		Price:  0,
	}
	_, err := p.IsValid()
	require.Error(t, err)
}
