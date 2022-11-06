package asosiasi

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	m1 := Mahasiswa{}
	m1.SetNim("117910001")
	m1.SetNama("Rahmat Rahimi")

	m2 := Mahasiswa{}
	m2.SetNim("117910002")
	m2.SetNama("Irsyad Sajali")

	d := Dosen{}

	d.SetNidn("ARI")
	d.SetNimMhs(m1.GetNim())
	d.SetNimMhs(m2.GetNim())

	fmt.Println("Kode Dosen : ", d.GetNidn())
	fmt.Println("Mengajar mahasiswa: ")

	for i := 0; i < d.GetJmlMhs(); i++ {
		fmt.Println("NIM : ", d.GetNimMhs(i))
	}

	require.Equal(t, "117910002", m2.GetNim())
	
}