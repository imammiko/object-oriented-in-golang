package asosiasi

type Mahasiswa struct {
	Nim   string
	Nama  string
	Prodi string
}

func (m *Mahasiswa) GetNim() string {
	return m.Nim
}

func (m *Mahasiswa) SetNim(nim string) {
	m.Nim = nim
}

func (m *Mahasiswa) GetNama() string {
	return m.Nama
}

func (m *Mahasiswa) SetNama(nama string) {
	m.Nama = nama
}

func (m *Mahasiswa) GetProdi() string {
	return m.Prodi
}
