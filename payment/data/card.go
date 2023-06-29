package data

import "fmt"

type card struct {
	Pan    string `json:"pan"`
	Expiry string `json:"expiry"`
}

func (c *card) String() string {
	return fmt.Sprintf("%v %v", c.GetPan(), c.GetExpiry())
}

func (c *card) SetPan(pan string) {
	c.Pan = pan
}

func (c *card) GetPan() string {
	return c.Pan
}

func (c *card) GetExpiry() string {
	return c.Expiry
}

func (c *card) SetExpiry(value string) {
	c.Expiry = value
}
