package handler

import (
	"github.com/robertke/orders-service/pkg/domain"
)

type (
	// CreateOrder ...
	CreateOrder struct {
		orderRepo
	}

	orderCreateCmd interface {
		GetID() string
		GetName() string
		GetDesc() string
	}

	orderRepo interface {
		FindOneByID(orderID string) *domain.Order
		FindAll() []*domain.Order
		Save(*domain.Order) error
	}
)

// NewCreateOrder ...
func NewCreateOrder(orderRepo orderRepo) *CreateOrder {
	return &CreateOrder{orderRepo}
}

// Handle ...
func (h *CreateOrder) Handle(cmd orderCreateCmd) error {
	return h.orderRepo.Save(
		domain.NewOrder(cmd.GetID(), cmd.GetName(), cmd.GetDesc()),
	)
}
