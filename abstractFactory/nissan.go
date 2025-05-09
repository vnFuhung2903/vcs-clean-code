package abstractFactory

type Nissan struct{}

func (factory *Nissan) Produce() ICar {
	return &Car{
		logo: "nissan",
	}
}
