package composition

type VGA struct {
	name      string
	kapasitas int
}

func (vga *VGA) GetName() string {
	return vga.name
}

func (vga *VGA) GetKapasitas() int {
	return vga.kapasitas
}

func NewVGA(nama string, kapasitas int) VGA {
	return VGA{name: nama, kapasitas: kapasitas}
}