package service

import (
	entity "github.com/IsaacDSC/pdv-golang/src/domain"
	"github.com/IsaacDSC/pdv-golang/src/shared/dto"
)

type OrderService struct {
	entity     entity.OrderInterface
	repository Repository
}

type Repository interface {
	createOrder(entity.Order) (err error)
}

func NewOrderService(
	order entity.OrderInterface,
	repo Repository,
) OrderService {
	return OrderService{
		entity:     order,
		repository: repo,
	}
}

// TODO: add retry and trace logger errors
func (os *OrderService) CreateOrder(input dto.OrderInputDto) (err error) {
	os.entity.ToDomain(input)
	if err = os.entity.Validate(); err != nil {
		return err
	}
	os.entity.SetCalculateProduct()
	if err = os.repository.createOrder(os.entity.Get()); err != nil {
		return
	}
	return
}
