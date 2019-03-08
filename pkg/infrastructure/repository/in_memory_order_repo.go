package repository

import "github.com/robertke/orders-service/pkg/domain"

type (
	// OrderRepo ...
	OrderRepo struct{}
)

var (
	orders = []*domain.Order{
		{
			"12121212", "Apple", "Order of apples",
		},
		{
			"343434", "Orange", "Order of oranges",
		},
		{
			"5565656", "Banana", "Order of bananas",
		},
	}
)

// NewInMemoryOrder ...
func NewInMemoryOrder() *OrderRepo {
	return &OrderRepo{}
}

// FindOneByID ...
func (r *OrderRepo) FindOneByID(orderID string) *domain.Order {

	for _, order := range orders {
		if order.ID == orderID {
			return order
		}
	}
	return nil
}

// FindAll ...
func (r *OrderRepo) FindAll() []*domain.Order {
	return orders
}

// Save ...
func (r *OrderRepo) Save(*domain.Order) error {
	return nil
}
