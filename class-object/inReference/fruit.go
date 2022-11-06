package inReference

type Fruit struct {
	Name  string
	Color string
}

func (f Fruit) GetColor() string {
	return f.Color
}

func (f Fruit) GetName() string {
	return f.Name
}

func (f Fruit) SetColor(color string) {
	f.Color = color
}

func (f Fruit) SetName(name string) {
	f.Name = name
}
