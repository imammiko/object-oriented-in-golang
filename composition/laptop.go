package composition

import (
	"fmt"
	"strconv"
)

type Laptop struct {
	merk, tipe string
	vga        VGA
	cpu        CPU
}

func (Laptop *Laptop) AddKomponen(cpu CPU, vga VGA) {
	Laptop.cpu = cpu
	Laptop.vga = vga
}

func (Laptop *Laptop) DisplayData() {
	fmt.Println("Laptop  " + Laptop.merk + "  dengan tipe" + Laptop.tipe)
	fmt.Println("Mempunyai Spesifikasi : ")
	fmt.Println("-> " + Laptop.cpu.GetName())
	fmt.Println("-> " + Laptop.vga.GetName())
	fmt.Println("  ->" + Laptop.vga.GetName()+" dengan kapasitas " + strconv.Itoa( Laptop.vga.GetKapasitas()))
}

func NewLaptop(merk string, tipe string) Laptop {
	return Laptop{merk: merk, tipe: tipe}
}