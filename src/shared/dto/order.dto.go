package dto

type OrderInputDto struct {
	Cart        []CartInputDto   `json:"cart"`
	Delivery    DeliveryInputDto `json:"delivery"`
	Client      ClientInputDto   `json:"client"`
	BillingType string           `json:"billingType"`
}

type CartInputDto struct {
	Price     float32 `json:"price"`
	ProductID string  `json:"productId"`
	Qtd       int     `json:"qtd"`
}

type DeliveryInputDto struct {
	ID          string `json:"id"`
	Address     string `json:"address"`
	HomeNumber  string `json:"homeNumber"`
	Observation string `json:"observation"`
}

type ClientInputDto struct {
	ClientName string `json:"clientName"`
}

type OrderOutputDto struct {
	OrderID    string                `json:"orderId"`
	Price      float32               `json:"price"`
	DeliveryID string                `json:"deliveryId"`
	OrderItems []OrderItemsOutputDto `json:"orderItems"`
}

type OrderItemsOutputDto struct {
	ProductID string  `json:"productId"`
	Qtd       int32   `json:"qtd"`
	Price     float32 `json:"price"`
}
