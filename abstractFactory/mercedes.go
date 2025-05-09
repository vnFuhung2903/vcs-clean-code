package abstractFactory

type Mercedes struct{}

func (factory *Mercedes) Produce() ICar {
	return &Car{
		logo: "mercedes",
	}
}
