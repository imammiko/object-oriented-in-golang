//case:
//www.w3schools.com/php/php_oop_classes_abstract.asp

package abstract

import "fmt"

type ICar interface {
	Intro() string
}

type Car struct {
	Name  string
}

type Audi struct {
	Car
}

func (n *Audi) Intro() string {
	return fmt.Sprintf("Choose German quality! I'm an %s", n.Name)
}

type Volvo struct {
	Car
}

func (n *Volvo) Intro() string {
	return fmt.Sprintf("Proud to be Swedish! I'm a %s", n.Name)
}

type Citron struct {
	Car
}

func (n *Citron) Intro() string {
	return fmt.Sprintf("French extravagance! I'm a %s", n.Name)
}

func NewAudi(name string) ICar {
	return &Audi{
		Car: Car{
			Name: name,
		},
	}
}

func  NewVolvo(name string) ICar {
	return &Volvo{
		Car: Car{
			Name: name,
		},
	}
}

func  NewCitron(name string) ICar {
	return &Citron{
		Car: Car{
			Name: name,
		},
	}
}