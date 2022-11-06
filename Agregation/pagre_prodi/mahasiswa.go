package pagreprodi

type Mahasiswa struct {
	nim  string
	nama string
}

func (m *Mahasiswa) GetNim() string {
	return m.nim
}

func (m *Mahasiswa) GetNama() string {
	return m.nama
}

func (m *Mahasiswa) SetNim(nim string) {
	m.nim = nim
}

func (m *Mahasiswa) SetNama(nama string) {
	m.nama = nama
}
