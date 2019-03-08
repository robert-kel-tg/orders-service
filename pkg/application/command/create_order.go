package command

type (
	// NewCreateOrder ...
	orderCreateCmd struct {
		ID          string
		Name        string
		Description string
	}
)

func NewCreateOrder(id, name, desc string) *orderCreateCmd {
	return &orderCreateCmd{id, name, desc}
}

// GetID ...
func (c *orderCreateCmd) GetID() string {
	return c.ID
}

// GetName ...
func (c *orderCreateCmd) GetName() string {
	return c.Name
}

// GetDesc ...
func (c *orderCreateCmd) GetDesc() string {
	return c.Description
}
