package agredrive

type Car struct {
	name string
	Engine
	Transmission
}

func (c *Car) AddEngine(engine Engine) {
	c.Engine = engine
}

func (c *Car) AddTransmision(transmission Transmission) {
	c.Transmission = transmission
}

func NewCar(name string) Car {
	return Car{
		name: name,
	}
}
