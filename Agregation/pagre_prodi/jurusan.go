package pagreprodi

type Jurusan struct {
	kode      string
	nama      string
	daftarMhs []Mahasiswa
	jmlMhs    int
}

func NewJurusan(kode string, nama string) Jurusan {
	return Jurusan{kode: kode, nama: nama}
}

func (j *Jurusan) GetKode() string {
	return j.kode
}

func (j *Jurusan) GetNama() string {
	return j.nama
}

func (j *Jurusan) GetJmlMhs() int {
	return j.jmlMhs
}

func (j *Jurusan) GetDaftarMhs() []Mahasiswa {
	return j.daftarMhs
}

func (j *Jurusan) AddMahasiswa(m Mahasiswa) {
	j.daftarMhs = append(j.daftarMhs, m)
	j.jmlMhs++
}
