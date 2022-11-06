package composition

type CPU struct {
	name string
}

func (cpu *CPU) GetName() string {
	return cpu.name
}

func NewCPU(name string) CPU {
	return CPU{name: name}
}
