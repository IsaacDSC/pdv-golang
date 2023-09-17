package repositories

import (
	"context"
	"fmt"
	"testing"

	entity "github.com/IsaacDSC/pdv-golang/src/domain"
	"github.com/IsaacDSC/pdv-golang/src/infra/settings"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTestOrderRepository(t *testing.T) {
	orderItems := map[string]entity.OrderItems{}
	order_itemsID := uuid.New()
	orderItems[order_itemsID.String()] = entity.OrderItems{
		ID:        order_itemsID,
		ProductID: uuid.NewString(),
		Quantity:  10,
		Price:     uint32(19999),
	}
	order := entity.Order{
		ID:            uuid.New(),
		OrderItems:    orderItems,
		BillingType:   "PIX",
		Price:         (10 * uint32(19999)),
		Discount:      0,
		PriceDelivery: 10,
		IsDelivery:    true,
		Status:        "OPEN",
		Delivery: entity.DeliveryEntity{
			ID:         uuid.NewString(),
			Country:    "BR",
			State:      "RJ",
			City:       "BM",
			Address:    "meu endereco 6969",
			NumberHome: "969",
			Reference:  "Final da Rua",
			Token:      uuid.NewString(),
		},
		Client: entity.ClientEntity{
			ID:           uuid.NewString(),
			Name:         uuid.New().String()[0:10],
			Email:        fmt.Sprintf("%s@hotmail.com", uuid.New().String()[0:10]),
			Telephone:    "24988180688",
			CPF:          "",
			Username:     uuid.New().String()[0:10],
			Password:     "",
			RecoverToken: "token",
		},
	}
	t.Run("Should be ", func(t *testing.T) {
		repo := &OrderRepository{}
		err := repo.CreateOrder(context.Background(), order)
		assert.NoError(t, err)
	})

	t.Run("DELETE INSERT ORDER_ITEMS", func(t *testing.T) {
		conn := settings.DbConn()
		query := fmt.Sprintf("DELETE FROM \"order_items\" WHERE order_items_id = '%s';", order_itemsID)
		rows, err := conn.Exec(query)
		assert.NoError(t, err)
		rowsAffected, err := rows.RowsAffected()
		assert.NoError(t, err)
		assert.Equal(t, rowsAffected, int64(1))
	})

	t.Run("DELETE INSERT ORDER", func(t *testing.T) {
		conn := settings.DbConn()
		query := fmt.Sprintf("DELETE FROM \"order\" WHERE order_id = '%s';", order.ID)
		rows, err := conn.Exec(query)
		assert.NoError(t, err)
		rowsAffected, err := rows.RowsAffected()
		assert.NoError(t, err)
		assert.Equal(t, rowsAffected, int64(1))
	})

	t.Run("DELETE INSERT DELIVERY", func(t *testing.T) {
		conn := settings.DbConn()
		query := fmt.Sprintf("DELETE FROM \"delivery\" WHERE delivery_id = '%s';", order.Delivery.ID)
		rows, err := conn.Exec(query)
		assert.NoError(t, err)
		rowsAffected, err := rows.RowsAffected()
		assert.NoError(t, err)
		assert.Equal(t, rowsAffected, int64(1))
	})

	t.Run("DELETE INSERT CLIENT", func(t *testing.T) {
		conn := settings.DbConn()
		query := fmt.Sprintf("DELETE FROM \"client\" WHERE client_id = '%s';", order.Client.ID)
		rows, err := conn.Exec(query)
		assert.NoError(t, err)
		rowsAffected, err := rows.RowsAffected()
		assert.NoError(t, err)
		assert.Equal(t, rowsAffected, int64(1))
	})

}
