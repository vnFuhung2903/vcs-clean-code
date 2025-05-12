package abstractFactory

type Factory struct {
	name string
}

func (factory *Factory) GetName() string {
	return factory.name
}
