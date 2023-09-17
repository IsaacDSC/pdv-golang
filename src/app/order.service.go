package service

import (
	"context"

	"github.com/IsaacDSC/pdv-golang/shared/dto"
	entity "github.com/IsaacDSC/pdv-golang/src/domain"
	"github.com/IsaacDSC/pdv-golang/src/infra/repositories"
)

type OrderService struct {
	entity     entity.OrderInterface
	repository repositories.OrderRepositoryInterface
}

type OrderServiceInterface interface {
	CreateOrder(ctx context.Context, input dto.OrderInputDto) (err error)
}

func NewOrderService(
	order entity.OrderInterface,
	repo repositories.OrderRepositoryInterface,
) OrderService {
	return OrderService{
		entity:     order,
		repository: repo,
	}
}

// TODO: add retry and trace logger errors
func (os *OrderService) CreateOrder(ctx context.Context, input dto.OrderInputDto) (err error) {
	os.entity.ToDomain(input)
	if err = os.entity.Validate(); err != nil {
		return err
	}
	os.entity.SetCalculateProduct()
	if err = os.repository.CreateOrder(ctx, os.entity.Get()); err != nil {
		return
	}
	return
}
