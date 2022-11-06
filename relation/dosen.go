package asosiasi

type Dosen struct {
	nidn    string
	nimmMhs []string
	jmlMhs  int
}

func (d *Dosen) GetNidn() string {
	return d.nidn
}

func (d *Dosen) SetNidn(nidn string) {
	d.nidn = nidn
}

func (d *Dosen) SetNimMhs(nimMhs string) {
	d.nimmMhs = append(d.nimmMhs, nimMhs)
	d.jmlMhs++
}

func (d *Dosen) GetJmlMhs() int {
	return d.jmlMhs
}

func (d *Dosen) GetNimMhs(indeks int) string {
	return d.nimmMhs[indeks]
}
