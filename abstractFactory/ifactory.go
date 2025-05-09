package abstractFactory

type IFactory interface {
	Produce() ICar
}

func NewFactory(name string) IFactory {
	if name == "mercedes" {
		return &Mercedes{}
	}
	if name == "nissan" {
		return &Nissan{}
	}
	return nil
}
