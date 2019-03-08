package response

import "github.com/robertke/orders-service/pkg/domain"

type (
	// OrdersResponse ...
	OrdersResponse struct {
		Orders []*domain.Order `json:"orders"`
	}
)

// NewOrders ...
func NewOrders(orders []*domain.Order) *OrdersResponse {
	return &OrdersResponse{orders}
}

// NewOrdersFromOne ...
func NewOrdersFromOne(order *domain.Order) *OrdersResponse {
	var (
		orders []*domain.Order
	)
	orders = append(orders, order)
	return &OrdersResponse{orders}
}
