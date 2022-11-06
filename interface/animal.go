package animal

type Animal interface {
	MakeSound() string
}

type Cat struct {
}

func (c *Cat) MakeSound() string {
	return "Meow"
}

type Dog struct{}

func (d *Dog) MakeSound() string {
	return "Woof"
}

type Mouse struct{}

func (m *Mouse) MakeSound() string {
	return "Squeak"
}

func NewCat() Animal {
	return &Cat{}
}

func NewDog() Animal {
	return &Dog{}
}

func NewMouse() Animal {
	return &Mouse{}
}