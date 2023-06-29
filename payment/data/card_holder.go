package data

type cardHolder struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (c *cardHolder) GetName() string {
	return c.Name
}

func (c *cardHolder) SetName(value string) {
	c.Name = value
}

func (c *cardHolder) GetEmail() string {
	return c.Email
}
