package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/IsaacDSC/pdv-golang/external/sqlc"
	entity "github.com/IsaacDSC/pdv-golang/src/domain"
	"github.com/IsaacDSC/pdv-golang/src/infra/settings"
)

type OrderRepositoryInterface interface {
	CreateOrder(ctx context.Context, input entity.Order) (err error)
}

type OrderRepository struct {
	db    *sqlc.Queries
	input *entity.Order
	ctx   context.Context
	tx    *sql.Tx
}

func (or *OrderRepository) CreateOrder(ctx context.Context, input entity.Order) (err error) {
	conn := settings.DbConn()
	tx, err := conn.Begin()
	defer tx.Rollback()
	if err != nil {
		return
	}

	or.input = &input
	or.tx = tx
	or.ctx = ctx
	or.db = sqlc.New(tx)
	if err = or.createClient(); err != nil {
		return
	}

	if err = or.createDelivery(); err != nil {
		return
	}

	if err = or.createOrder(); err != nil {
		return
	}

	if err = or.createOrderItems(); err != nil {
		return
	}

	return tx.Commit()
	// return nil
}

func (or *OrderRepository) createClient() (err error) {
	return or.db.CreateClient(or.ctx, sqlc.CreateClientParams{
		ClientID:  or.input.Client.ID,
		Email:     sql.NullString{String: or.input.Client.Email, Valid: true},
		Username:  sql.NullString{String: or.input.Client.Username, Valid: true},
		Telephone: sql.NullString{String: or.input.Client.Telephone, Valid: true},
	})
}

func (or *OrderRepository) createDelivery() (err error) {
	return or.db.CreateDelivery(or.ctx, sqlc.CreateDeliveryParams{
		DeliveryID:  or.input.Delivery.ID,
		Country:     or.input.Delivery.Country,
		State:       or.input.Delivery.State,
		Address:     or.input.Delivery.Address,
		NumberHome:  or.input.Delivery.NumberHome,
		Observation: sql.NullString{String: or.input.Delivery.Reference, Valid: true},
		Token:       or.input.Delivery.Token,
		ClientID:    or.input.Client.ID,
	})
}

func (or *OrderRepository) createOrder() (err error) {
	return or.db.CreateOrder(or.ctx, sqlc.CreateOrderParams{
		OrderID:       or.input.ID.String(),
		Price:         float64(or.input.Price),
		BillingType:   sql.NullString{String: or.input.BillingType, Valid: true},
		Discount:      sql.NullFloat64{Float64: 0, Valid: true},
		PriceDelivery: float64(or.input.PriceDelivery),
		IsDelivery:    sql.NullBool{Bool: or.input.IsDelivery, Valid: true},
		Status:        sql.NullString{String: or.input.Status, Valid: true},
		ClientID:      or.input.Client.ID,
		DeliveryID:    or.input.Delivery.ID,
	})
}

func (or *OrderRepository) createOrderItems() (err error) {
	const query = `INSERT INTO order_items (order_items_id, price, product_id, quantity, order_id, created_at) VALUES `
	var (
		counter int
		values  string
	)
	for index := range or.input.OrderItems {
		fmt.Println(counter, len(or.input.OrderItems)-1)
		if len(or.input.OrderItems) == 1 || len(or.input.OrderItems)-1 == counter {
			values += fmt.Sprintf(
				"( '%s', %f, '%s', %d, '%s', NOW() )",
				or.input.OrderItems[index].ID,
				or.input.OrderItems[index].Price,
				or.input.OrderItems[index].ProductID,
				or.input.OrderItems[index].Quantity,
				or.input.ID,
			)
			counter++
		} else {
			values += fmt.Sprintf(
				"( '%s', %f, '%s', %d, '%s', NOW() ),",
				or.input.OrderItems[index].ID,
				or.input.OrderItems[index].Price,
				or.input.OrderItems[index].ProductID,
				or.input.OrderItems[index].Quantity,
				or.input.ID,
			)
			counter++
		}
	}
	completeQuery := fmt.Sprintf("%s %s", query, values)
	_, err = or.tx.Exec(completeQuery)
	return
}
