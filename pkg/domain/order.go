package domain

type (
	// Order ...
	Order struct {
		ID          string `json:"ID"`
		Name        string `json:"Name"`
		Description string `json:"Desc"`
	}
)

// NewOrder ...
func NewOrder(id, name, desc string) *Order {
	return &Order{id, name, desc}
}
