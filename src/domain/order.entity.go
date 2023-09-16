package entity

import (
	"errors"
	"fmt"

	"github.com/IsaacDSC/pdv-golang/src/infra/settings"
	"github.com/IsaacDSC/pdv-golang/src/shared/dto"
	"github.com/google/uuid"
)

type Order struct {
	ID            uuid.UUID
	OrderItems    map[string]orderItems
	BillingType   string
	Price         float32
	Discount      float32
	PriceDelivery float32
	TaxDelivery   bool
	Delivery      DeliveryEntity
	Client        ClientEntity
}

type orderItems struct {
	ID        uuid.UUID
	ProductID string
	Quantity  int32
	Price     float32
}

type OrderInterface interface {
	ToDomain(input dto.OrderInputDto)
	Validate() (err error)
	SetCalculateProduct()
	Get() Order
}

func NewOrderEntity() Order {
	return Order{}
}

func (o *Order) ToDomain(input dto.OrderInputDto) {
	o = &Order{
		ID:          uuid.New(),
		BillingType: input.BillingType,
		Client: ClientEntity{
			Name: input.Client.ClientName,
		},
		Delivery: DeliveryEntity{
			ID:         fmt.Sprintf("%s:%s:%s", input.Delivery.ID, input.Delivery.Address, input.Delivery.HomeNumber),
			Country:    "BR",
			State:      "RJ",
			City:       "BM",
			Address:    input.Delivery.Address,
			NumberHome: input.Delivery.HomeNumber,
			Reference:  input.Delivery.Observation,
		},
	}
	for index := range input.Cart {
		mapProducts := make(map[string]orderItems)
		mapProducts[input.Cart[index].ProductID] = orderItems{
			ID:        uuid.New(),
			ProductID: input.Cart[index].ProductID,
			Quantity:  int32(input.Cart[index].Qtd),
			Price:     input.Cart[index].Price,
		}
	}
}

func (o *Order) Validate() (err error) {
	template := settings.GetTemplateSingleton()
	var existDeliveryID bool
	for index := range template.Delivery {
		if template.Delivery[index].ID == o.Delivery.ID {
			existDeliveryID = true
			o.TaxDelivery = true
			o.PriceDelivery = template.Delivery[index].Price
			break
		}
	}
	if !existDeliveryID {
		err = errors.New("Not-Found-DeliveryID")
	}
	products := template.Pages[0].Home.Menu.Products
	mapProducts := make(map[string]orderItems)
	for index := range products {
		mapProducts[products[index].ID] = orderItems{
			ProductID: products[index].ID,
			Price:     products[index].Price,
		}
	}
	for index := range o.OrderItems {
		if o.OrderItems[index].ProductID != mapProducts[o.OrderItems[index].ProductID].ProductID {
			err = errors.New("Not-Found-Product-In-Schema")
			break
		}
	}
	if err != nil {
		return
	}
	return
}

func (o *Order) SetCalculateProduct() {
	for index := range o.OrderItems {
		o.Price += o.OrderItems[index].Price
	}
}

func (o *Order) Get() Order {
	return *o
}
