package composition

import "testing"

func TestRun(t *testing.T) {


	AsusLaptop := NewLaptop("Asus", "x450c")

	AsusLaptop.AddKomponen(NewCPU("Intel Core i3"), NewVGA("Nvidia", 2))

	AsusLaptop.DisplayData()

	ToshibahLaptop := NewLaptop("Toshiba", "s700")

	ToshibahLaptop.AddKomponen(NewCPU("Intel Core i5"), NewVGA("AMD", 4))

	ToshibahLaptop.DisplayData()
}