package agredrive

type Engine struct {
	name       string
	horsePower int
}

func (e *Engine) GetHorsePower() int {
	return e.horsePower
}

func (e *Engine) SetHorsePower(horsePower int) {
	e.horsePower = horsePower
}

func (e *Engine) GetName() string {
	return e.name
}

func NewEngine(name string) Engine {
	return Engine{
		name: name,
	}
}
