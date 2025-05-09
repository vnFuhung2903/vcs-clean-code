package abstractFactory

type ICar interface {
	GetLogo() string
}

type Car struct {
	logo string
}

func (c *Car) GetLogo() string {
	return c.logo
}
