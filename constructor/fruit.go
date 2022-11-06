package constructor

type Fruit struct {
	Name  string
	Color string
}

func (f *Fruit) GetColor() string {
	return f.Color
}

func (f *Fruit) GetName() string {
	return f.Name
}

func NewFruit(name, color string) *Fruit {
	return &Fruit{
		Name:  name,
		Color: color,
	}
}
