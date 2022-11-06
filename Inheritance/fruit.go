package inheritance

type Fruit struct {
	Name  string
	Color string
}

type Strawberry struct {
	Fruit Fruit
}

func (s *Strawberry) Message() string {
	return "I am a strawberry"
}

func NewStrawberry(name, color string) *Strawberry {
	return &Strawberry{
		Fruit: Fruit{
			Name:  name,
			Color: color,
		},
	}
}
