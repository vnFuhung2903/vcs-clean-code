package abstractFactory

type IFactory interface {
	GetName() string
}

func NewFactory(name string) IFactory {
	return &Factory{
		name,
	}
}
