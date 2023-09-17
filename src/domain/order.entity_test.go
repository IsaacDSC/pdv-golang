package entity

import (
	"testing"

	"github.com/IsaacDSC/pdv-golang/shared/dto"
	"github.com/stretchr/testify/assert"
)

func TestOrderDomain(t *testing.T) {
	order := new(Order)
	t.Run("Should be transform dto order in domain order Empty", func(t *testing.T) {
		order.ToDomain(dto.OrderInputDto{})
	})
	t.Run("Should be transform dto order in domain order", func(t *testing.T) {
		order.ToDomain(dto.OrderInputDto{})
		// order.Validate()
		// order.Get()
		// order.SetCalculateProduct()
	})

	t.Run("Should be validate successfully domain order", func(t *testing.T) {
		err := order.Validate()
		assert.NoError(t, err)
		// order.Get()
		// order.SetCalculateProduct()
	})

	t.Run("Should be validate successfully domain order", func(t *testing.T) {
		order.SetCalculateProduct()
		assert.Equal(t, order.Price, 0)
	})

	t.Run("", func(t *testing.T) {
		assert.Equal(t, order.Get(), order)
	})
}
