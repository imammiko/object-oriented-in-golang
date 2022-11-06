package fruit

type Fruit struct {
	Name  string //public
	color string //private
}

func (f *Fruit) GetColor() string { //public
	return f.color
}

func (f *Fruit) SetColor(color string) { //public
	f.color = color
}

func (f *Fruit) getName() string { //private
	return f.Name
}

func (f *Fruit) setName(name string) { //private
	f.Name = name
}
